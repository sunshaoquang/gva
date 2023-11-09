package example


type ExcelInfo struct {
	FileName string               `json:"fileName"` // 文件名
	InfoList []interface{} `json:"infoList"`
}