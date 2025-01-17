package login

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	apiSystem "github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/levigross/grequests"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LoginApi struct {
}

type BaseResponse struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

type WxLoginResponse struct {
	BaseResponse
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid,omitempty"`
}
type WxLoginResponseToken struct {
	BaseResponse
	Token  string `json:"access_token"`
	OpenID string `json:"openid"`
}

type Watermark struct {
	AppID     string `json:"appid"`
	Timestamp int    `json:"timestamp"`
}

type PhoneNumberData struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       Watermark `json:"watermark"`
}

type UserInfo struct {
	OpenID    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	UnionID   string `json:"unionId,omitempty"`
}

type WxLoginRequest struct {
	Code          string `json:"code"`
	EncryptedData string `json:"encryptedData"`
	IV            string `json:"iv"`
}

type WxUserBindPasswordRequest struct {
	WxLoginRequest
	Username string `json:"username"`
	Password string `json:"password"`
}

func (loginApi *LoginApi) TempTest(c *gin.Context) {
	response.OkWithData("hello", c)
}

func (loginApi *LoginApi) WxUserBindPassword(c *gin.Context) {
	var request WxUserBindPasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	appid := global.GVA_CONFIG.Wx.AppID
	secret := global.GVA_CONFIG.Wx.AppSecret
	wxURL := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, request.Code)

	resp, err := grequests.Get(wxURL, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get response from Weixin"})
		return
	}

	var wxResp WxLoginResponse
	if err := json.Unmarshal(resp.Bytes(), &wxResp); err != nil {
		c.JSON(500, gin.H{"error": "Failed to parse response from Weixin"})
		return
	}

	if wxResp.ErrCode != 0 {
		c.JSON(500, gin.H{"error": wxResp.ErrMsg})
		return
	}

	userInfo, err := decryptUserInfo(request.EncryptedData, request.IV, wxResp.SessionKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to decrypt user info"})
		return
	}

	var user system.SysUser
	err = global.GVA_DB.Where("openid = ?", wxResp.OpenID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//FEATURE 新增一个接口用于单独查询帐号是否重复
		// 检查用户名是否重复
		var existingUser system.SysUser
		err = global.GVA_DB.Where("username = ?", request.Username).First(&existingUser).Error
		if err == nil {
			c.JSON(400, gin.H{"error": "Username already exists"})
			return
		} else if err != gorm.ErrRecordNotFound {
			c.JSON(500, gin.H{"error": "Failed to check username"})
			return
		}

		// 用户不存在，自动注册
		user = system.SysUser{
			UUID:     uuid.Must(uuid.NewV4()),
			OpenID:   wxResp.OpenID,
			Username: request.Username,
			Password: utils.BcryptHash(request.Password), // 存储加密后的密码
			NickName: userInfo.NickName,
			Avatar:   userInfo.AvatarURL,
		}
		if err := global.GVA_DB.Create(&user).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user"})
			return
		}
	} else if err != nil {
		c.JSON(500, gin.H{"error": "Failed to query user"})
		return
	} else {
		// 用户存在，检查用户名是否为空或需要更新
		if user.Username != request.Username {
			// 检查新的用户名是否重复
			var existingUser system.SysUser
			err = global.GVA_DB.Where("username = ?", request.Username).First(&existingUser).Error
			if err == nil {
				c.JSON(400, gin.H{"error": "帐号已经被其他用户绑定,请更换一个"})
				return
			} else if err != gorm.ErrRecordNotFound {
				c.JSON(500, gin.H{"error": "程序出错,请联系开发人员"})
				global.GVA_LOG.Error("Failed to check username", zap.Error(err))
				return
			}
			user.Username = request.Username
		}

		// 更新用户信息
		user.Password = utils.BcryptHash(request.Password)
		if err := global.GVA_DB.Save(&user).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update user"})
			return
		}
	}

	c.JSON(200, gin.H{"message": "密码绑定成功"})
}

func (loginApi *LoginApi) WxPhoneLogin(c *gin.Context) {
	var request WxLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, wxResp, err := handleWxLogin(request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	phoneNumberData, err := decryptPhoneNumber(request.EncryptedData, request.IV, wxResp.SessionKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to decrypt phone number"})
		return
	}

	c.JSON(200, gin.H{
		"openid":      wxResp.OpenID,
		"userInfo":    user,
		"phoneNumber": phoneNumberData.PhoneNumber,
	})
}
func (loginApi *LoginApi) WxLogin(c *gin.Context) {
	var request WxLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	global.GVA_LOG.Debug("json绑定成功开始进行数据验证")
	//user, wxResp, err := handleWxLogin(request)
	user, _, err := handleWxLogin(request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		//c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	systemResLoginStruct, err := apiSystem.TokenNext(c, user)
	if err != nil {
		response.FailWithDetailed(err, "登录失败", c)
	}
	global.GVA_LOG.Debug("登录成功")
	response.OkWithDetailed(systemResLoginStruct, "登录成功", c)
}

// 手机号和不用手机号登录公共部分
func handleWxLogin(request WxLoginRequest) (system.SysUser, WxLoginResponse, error) {
	appid := global.GVA_CONFIG.Wx.AppID
	secret := global.GVA_CONFIG.Wx.AppSecret
	wxURL := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, request.Code)
	global.GVA_LOG.Debug("开始请求微信...")
	resp, err := grequests.Get(wxURL, nil)
	global.GVA_LOG.Debug("得到结果,开始验证是否成功获取数据")
	if err != nil {
		global.GVA_LOG.Error("微信一键登录向微信获取数据失败", zap.Error(err))
		return system.SysUser{}, WxLoginResponse{}, errors.New(fmt.Sprintf("Failed to get response from Weixin Error:%s", err))
	}
	global.GVA_LOG.Debug("请求获取数据成功...")

	var wxResp WxLoginResponse
	if err := json.Unmarshal(resp.Bytes(), &wxResp); err != nil {
		return system.SysUser{}, WxLoginResponse{}, errors.New(fmt.Sprintf("Failed to parse response from Weixin Error:%s", err))
	}
	global.GVA_LOG.Debug("解析数据")

	if wxResp.ErrCode != 0 {
		return system.SysUser{}, WxLoginResponse{}, errors.New(wxResp.ErrMsg)
	}
	global.GVA_LOG.Debug("解析数据成功,开始解密转化为明文")

	userInfo, err := decryptUserInfo(request.EncryptedData, request.IV, wxResp.SessionKey)

	if err != nil {
		return system.SysUser{}, WxLoginResponse{}, errors.New(fmt.Sprintf("Failed to decrypt user info Error:%s", err))
	}
	global.GVA_LOG.Debug("转化为明文成功,微信自动登录完成")

	var user system.SysUser
	err = global.GVA_DB.Where("openid = ?", wxResp.OpenID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 用户不存在，自动注册
		user = system.SysUser{
			UUID:        uuid.Must(uuid.NewV4()),
			OpenID:      wxResp.OpenID,
			NickName:    userInfo.NickName,
			Username:    fmt.Sprintf("xy_%s", uuid.Must(uuid.NewV4())),
			Avatar:      userInfo.AvatarURL,
			AuthorityId: 8881,
			Authorities: []system.SysAuthority{
				{
					AuthorityId: 8881, // 自动注册默认角色
				},
			},
		}
		if err := global.GVA_DB.Create(&user).Error; err != nil {
			return system.SysUser{}, WxLoginResponse{}, errors.New(fmt.Sprintf("Failed to create user Error:%s", err))
		}
	} else if err != nil {
		return system.SysUser{}, WxLoginResponse{}, errors.New(fmt.Sprintf("Failed to query user Error:%s", err))
	} else {
		// 用户存在，但 `openid` 不存在
		if user.OpenID == "" {
			user.OpenID = wxResp.OpenID
			if err := global.GVA_DB.Save(&user).Error; err != nil {
				return system.SysUser{}, WxLoginResponse{}, errors.New(fmt.Sprintf("Failed to update user with new OpenID Error:%s", err))
			}
		}
	}

	return user, wxResp, nil
}

// TODO 工具函数规范 以及上分的路由函数和逻辑函数分开
func decryptUserInfo(encryptedData, iv, sessionKey string) (*UserInfo, error) {
	decodedEncryptedData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	decodedSessionKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	decodedIv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(decodedSessionKey)
	if err != nil {
		return nil, err
	}
	if len(decodedEncryptedData) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	cbc := cipher.NewCBCDecrypter(block, decodedIv)
	decrypted := make([]byte, len(decodedEncryptedData))
	cbc.CryptBlocks(decrypted, decodedEncryptedData)

	decrypted = pkcs7Unpad(decrypted)

	var userInfo UserInfo
	if err := json.Unmarshal(decrypted, &userInfo); err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func decryptPhoneNumber(encryptedData, iv, sessionKey string) (*PhoneNumberData, error) {
	decodedEncryptedData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	decodedSessionKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	decodedIv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(decodedSessionKey)
	if err != nil {
		return nil, err
	}
	if len(decodedEncryptedData) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	cbc := cipher.NewCBCDecrypter(block, decodedIv)
	decrypted := make([]byte, len(decodedEncryptedData))
	cbc.CryptBlocks(decrypted, decodedEncryptedData)

	decrypted = pkcs7Unpad(decrypted)

	var phoneNumberData PhoneNumberData
	if err := json.Unmarshal(decrypted, &phoneNumberData); err != nil {
		return nil, err
	}
	return &phoneNumberData, nil
}

func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil
	}
	return data[:(length - unpadding)]
}
