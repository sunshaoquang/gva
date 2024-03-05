package example

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"strings"
	"reflect"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/xuri/excelize/v2"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

)

type ExcelService struct{}

type ExcelSheetSettings struct {
    SheetName      string
    SheetRow       string
    SheetHeadName  []string
    SheetHeadValue []string
    SheetModel     interface{} // 这里可能需要更具体的类型
}
var excelSheetList = map[string]ExcelSheetSettings{
    "客户列表": {
        SheetName:      "Sheet1",
        SheetRow:       "A1",
        SheetHeadName:  []string{"姓名", "电话"},
        SheetHeadValue: []string{"CustomerName", "CustomerPhoneData"},
        SheetModel:     example.ExaCustomer{},
    },
	"推广市场目标表": {
        SheetName:      "Sheet1",
        SheetRow:       "A1",
        SheetHeadName:  []string{"店铺名称", "店铺代码","任务目标","场景类型"},
        SheetHeadValue: []string{"ShopName", "ShopCode","Target","Type"},
        SheetModel:     manual.PromotionMarketTarget{},
    },
	"wmsKpiAll": {
        SheetName:      "Sheet1",
        SheetRow:       "A1",
        SheetHeadName:  []string{"项目", "2021年","2022年","2023年"},
        SheetHeadValue: []string{"Project", "Year2021","Year2022","Year2023"},
        SheetModel:     manual.WmsKpiSummary2Mi{},
    },
	"wmsKpi2023": {
        SheetName:      "Sheet1",
        SheetRow:       "A1",
        SheetHeadName:  []string{"指标分类","指标名称","目标值","1月","2月","3月","4月","5月","6月","7月","8月","9月","10月","11月","12月"},
        SheetHeadValue: []string{"IndicatorClassification","IndicatorName","Target","January","February","March","April","May","June","July","August","September","October","November","December"},
        SheetModel:     manual.WmsKpi2023SummaryMi{},
    },
	"物流成本明细表": {
        SheetName:      "Sheet1",
        SheetRow:       "A1",
        SheetHeadName:  []string{"产品简称","国家","类型","1月","2月","3月","4月","5月","6月","7月","8月","9月","10月","11月","12月"},
        SheetHeadValue: []string{"ProductAbbreviation","Country","Type","January","February","March","April","May","June","July","August","September","October","November","December"},
        SheetModel:     manual.WmsLogisticsPcs2023Mi{},
    },
	"主要产品成本降本汇总表": {
        SheetName:      "Sheet1",
        SheetRow:       "A1",
        SheetHeadName:  []string{"产品简称","国家","降本基数10月成本","承诺达成降本金额","1月VS10月降本","2月VS10月降本","3月VS10月降本","4月VS10月降本","5月VS10月降本","6月VS10月降本","7月VS10月降本","8月VS10月降本","9月VS10月降本","10月VS10月降本","11月VS10月降本","12月VS10月降本","1-当前月累计加权降本金额","1-当前月累计加权降本完成度"},
        SheetHeadValue: []string{"ProductAbbreviation","Country","ReduceBaseCostInOctober","CommitmentToAchieveCostReductionAmount","CostReductionInMonth1VsMonth10","CostReductionInMonth2VsMonth10","CostReductionInMonth3VsMonth10","CostReductionInMonth4VsMonth10","CostReductionInMonth5VsMonth10","CostReductionInMonth6VsMonth10","CostReductionInMonth7VsMonth10","CostReductionInMonth8VsMonth10","CostReductionInMonth9VsMonth10","CostReductionInMonth10VsMonth10","CostReductionInMonth11VsMonth10","CostReductionInMonth12VsMonth10","MonthlyCumulativeWeightedPrincipalReductionAmount","MonthlyCumulativeWeightedCostReductionCompletionDegree"},
        SheetModel:     manual.WmsLogisticsPcsSummaryMi{},
    },
	"主要产品成本明细表": {
        SheetName:      "Sheet1",
        SheetRow:       "A1",
        SheetHeadName:  []string{"产品简称","国家","费用类型CNY","降本基数(2022年10月)","2022年11月","2022年12月","2023年1月","2023年2月","2023年3月","2023年4月","2023年5月","2023年6月","2023年7月","2023年8月","2023年9月","2023年10月","2023年11月","2023年12月"},
        SheetHeadValue: []string{"ProductAbbreviation","Country","FeeTypeCNY","Year2022Month10","Year2022Month11","Year2022Month12","Year2023Month1","Year2023Month2","Year2023Month3","Year2023Month4","Year2023Month5","Year2023Month6","Year2023Month7","Year2023Month8","Year2023Month9","Year2023Month10","Year2023Month11","Year2023Month12"},
        SheetModel:     manual.WmsLogisticsPcsDetailMi{},
    },
	"物流占销比": {
        SheetName:      "Sheet1",
        SheetRow:       "A1",
        SheetHeadName:  []string{"区域","月份(YYYY-MM-DD)","指标","指标值"},
        SheetHeadValue: []string{"Area","$Mmonth","Name","Value"},
        SheetModel:     manual.WmsLogisticsRateMonths{},
    },
	"费用by运输方式": {
        SheetName:      "Sheet1",
        SheetRow:       "A1",
        SheetHeadName:  []string{"区域","运输方式","月份(YYYY-MM-DD)","费用"},
        SheetHeadValue: []string{"Area","TransMethod","$Mmonth","Cost"},
        SheetModel:     manual.WmsLogisticsModeTransportMi{},
    },
	"费用by供应商":{
		SheetName:      "Sheet1",
        SheetRow:       "A1",
		SheetHeadName:  []string{"类型","区域","目的国","供应商","1月","2月","3月","4月","5月","6月","7月","8月","9月","10月","11月","12月"},
        SheetHeadValue: []string{"Type","Area","Country","Supplier","Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sept","Oct","Nov","Dec"},
		SheetModel:     manual.WmsLogisticsProviderMi{},
	},
	"收入by国家":{
		SheetName:      "Sheet1",
        SheetRow:       "A1",
		SheetHeadName:  []string{"类型","区域","目的国","1月","2月","3月","4月","5月","6月","7月","8月","9月","10月","11月","12月"},
        SheetHeadValue: []string{"Type","Area","Country","Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sept","Oct","Nov","Dec"},
		SheetModel:     manual.WmsLogisticsSalebycountryMi{},
	},
	"仓储费":{
		SheetName:      "Sheet1",
        SheetRow:       "A1",
		SheetHeadName:  []string{"区域","费用类型","月份(YYYY-MM-DD)","不含税费用"},
		SheetHeadValue: []string{"Area","CostType","$Mmonth","ExcludingTaxCost"},
		SheetModel:     manual.WmsLogisticsStoragefeeMi{},
	},
	"按天销售目标录入表":{
		SheetName:      "Sheet1",
        SheetRow:       "A1",
		SheetHeadName:  []string{"类型","月份","渠道","版本","日期(YYYY-MM-DD)","销售目标"},
		SheetHeadValue: []string{"Department","Month","Channel","Version","$Date","AmountTarget"},
		SheetModel:     manual.OmsCnSalesTargetDaily{},
	},
	"京东自营销售录入表":{
		SheetName:      "Sheet1",
        SheetRow:       "A1",
		SheetHeadName:  []string{"月份(YYYY-MM-DD)","渠道","类型","负责人","回款收入(未税)"},
		SheetHeadValue: []string{"$Date","Channel","Department","Principal","Amount"},
		SheetModel:     manual.OmsCnChannelSales{},
	},
}


func (exa *ExcelService) ParseInfoList2Excel(infoList []interface{}, filePath string) error {

	obj, ok := excelSheetList[filePath]
	if !ok {
		// 处理类型断言失败的情况
		fmt.Println("无法访问%s的设置",filePath)
	}
	excel := excelize.NewFile()
	
	excel.SetSheetRow(obj.SheetName, obj.SheetRow, &obj.SheetHeadName)
	data := make([][]interface{}, len(infoList))

	for i, info := range infoList {
		infoMap, ok := info.(map[string]interface{})
		if !ok {
			fmt.Println("info 不是 map 类型")
		} else {
			// 初始化一个新的切片来存储属性值
			data[i] = []interface{}{}
	
			for _, propertyValue := range obj.SheetHeadValue {
				var valueItem string
				// 使用反射来获取属性值
				if strings.HasPrefix(propertyValue, "$") {
					// 去掉$
					valueItem = strings.Trim(propertyValue, "$")
				}else{
					valueItem = propertyValue
				}
				value := infoMap[valueItem]
				
				// 检查值是否有效
				if value != nil {
					// 转换值为接口类型并存储在 data 中
					data[i] = append(data[i], value)
				} else {
					// 如果属性不存在，可以选择执行适当的处理
					data[i] = append(data[i], nil)
				}
			}
		}
	}
	for i, row := range data {
		index,err := strconv.Atoi(obj.SheetRow[0:])
		if err != nil {
			index = 1
		}
		axis := fmt.Sprintf("A%d", i+1+ index)
		
		excel.SetSheetRow(obj.SheetName, axis, row)
	}
	err := excel.SaveAs(filePath)
	return err
}

func (exa *ExcelService) OutputExcelInfoListData(infos []interface{},fileName string) error {

	// 将 infos 转换为 []map[string]interface{}
	obj, ok := excelSheetList[fileName]
	if !ok {
		// 处理类型断言失败的情况
		fmt.Println("无法访问%s的设置",fileName)
	}
	infoList := make([]map[string]interface{}, len(infos))
	for i, info := range infos {
		infoMap, ok := info.(map[string]interface{})
		if !ok {
			// 处理类型断言失败的情况，如果 info 不是 map 类型
			fmt.Println("info 不是 map 类型")
		} else {
			// 进行后续操作
			infoList[i] = infoMap
		}
		
	}
	
	var err error = nil
	if global.GVA_DB.Migrator().HasTable(&obj.SheetModel) {
		global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			if err = tx.Model(&obj.SheetModel).Create(&infoList).Error; err != nil {
				  return err
			}
			return nil
		})
	}else{
		global.MustGetGlobalDBByDBName("ht_smartdata").Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&obj.SheetModel).Create(&infoList).Error; err != nil {
				  return err
			}
			stmt := &gorm.Statement{DB: tx}
			stmt.Parse(&obj.SheetModel)
			backupTableName := fmt.Sprintf("%s_%s",stmt.Schema.Table,time.Now().Format("20060102150405"))
			// 使用tx.Model(&obj.SheetModel)来获取表名
			// 使用create table as
			if err := tx.Exec(fmt.Sprintf("CREATE TABLE %s AS SELECT * FROM %s",backupTableName , stmt.Schema.Table)).Error; err!= nil {
				return err
			}
		
			return nil
		})
	}
	return err
}

func (exa *ExcelService) ParseExcel2InfoList(c *gin.Context,fileName, sheetName,createdName string) ([]interface{}, error) {

	skipHeader := true
	obj, ok := excelSheetList[fileName]
	if !ok {
		// 处理类型断言失败的情况
		fmt.Println("无法访问%s的设置",fileName)
	}
	
	fixedHeader := obj.SheetHeadName
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir+fmt.Sprintf("%s__ExcelImport.xlsx",fileName),excelize.Options{RawCellValue: true})
	if err != nil {
		return nil, err
	}
	infos := make([]interface{}, 0)
	rows, err := file.GetRows(obj.SheetName)
	if err != nil {
		return nil, err
	}
	for j,row:= range rows {
		if err != nil {
			return nil, err
		}

		if skipHeader {
			if exa.compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("Excel格式错误在第"+strconv.Itoa(j+1)+"行【"+strings.Join(row, ",")+"】")
			}
		}
		if len(row) != len(fixedHeader) {
			//如果row没有这么长不要忽略 给他值空
			// 但是如果row一个都没有就忽略
			if len(row) == 0 {
				continue
			}
			// 计算需要填充的空值数量
			numEmptyValues := len(fixedHeader) - len(row)

			// 创建一个切片，其中包含需要填充的空值（使用空字符串）
			emptyValues := make([]string, numEmptyValues)
			for i := range emptyValues {
				emptyValues[i] = ""
			}
		
			// 将空值切片添加到 row 切片末尾
			row = append(row, emptyValues...)
		}
		

		info := make(map[string]interface{}, 0)
		
		for i,_ := range row { 
			cellName,err:= excelize.CoordinatesToCellName(i+1, j+1);
			var valueItem string
			if err != nil {
				fmt.Println(err)
				continue
			}
			// 获取单元格的原始值
			var originalValue string
			
			if strings.HasPrefix(obj.SheetHeadValue[i], "$") {
				// 去掉$
				valueItem = strings.Trim(obj.SheetHeadValue[i], "$")
				// 去掉空格
				valueItem = strings.TrimSpace(valueItem)
				originalValue, err = file.GetCellValue(obj.SheetName, cellName,excelize.Options{RawCellValue: false})
			}else{
				valueItem = strings.TrimSpace(obj.SheetHeadValue[i])
				originalValue, err = file.GetCellValue(obj.SheetName, cellName,excelize.Options{RawCellValue: true})
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			info[valueItem] = originalValue
		}
		// 获取结构体类型
		structType := reflect.TypeOf(obj.SheetModel)

		// 要检查的字段名
	
		// 检查结构体中是否包含指定字段
		if hasField(structType, "SysUserID") {
			info["SysUserID"] = utils.GetUserID(c)
		}
		if hasField(structType, "SysUserAuthorityID") {
			info["SysUserAuthorityID"] = utils.GetUserAuthorityId(c)
		}
		if hasField(structType, "CreatedName") {
			if createdName != "" {
				info["CreatedName"] = createdName
			} else {
				info["CreatedName"] = utils.GetUserName(c)
			}
		}
		if hasField(structType, "CreatedBy") {
			info["CreatedBy"] = 1
		}
		if hasField(structType, "UpdatedBy") {
			info["UpdatedBy"] = 0
		}
		if hasField(structType, "DeletedBy") {
			info["DeletedBy"] = 0
		}
		if hasField(structType, "SheetName") {
			info["SheetName"] = sheetName
		}
		info["CreatedAt"] = time.Now()
		info["UpdatedAt"] = time.Now()
		infos = append(infos, info)
	}
	return infos, nil
}

func (exa *ExcelService) compareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}

func hasField(structType reflect.Type, fieldName string) bool {
    for i := 0; i < structType.NumField(); i++ {
        field := structType.Field(i)
        if field.Name == fieldName {
            return true
        }
    }
    return false
}