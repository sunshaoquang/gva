// 自动生成模板WmsLogisticsPcsSummaryMi
package manual

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 主要产品成本降本汇总表 结构体  WmsLogisticsPcsSummaryMi
type WmsLogisticsPcsSummaryMi struct {
      global.GVA_MODEL
      ProductAbbreviation  string `json:"productAbbreviation" form:"productAbbreviation" gorm:"column:product_abbreviation;comment:产品简称;size:255;"`  //产品简称 
      Country  string `json:"country" form:"country" gorm:"column:country;comment:国家;size:255;"`  //国家 
      ReduceBaseCostInOctober  string `json:"reduceBaseCostInOctober" form:"reduceBaseCostInOctober" gorm:"column:reduce_base_cost_in_october;comment:降本基数10月成本;size:255;"`  //降本基数10月成本 
      CommitmentToAchieveCostReductionAmount  string `json:"commitmentToAchieveCostReductionAmount" form:"commitmentToAchieveCostReductionAmount" gorm:"column:commitment_to_achieve_cost_reduction_amount;comment:承诺达成降本金额;size:255;"`  //承诺达成降本金额 
      CostReductionInMonth1VsMonth10  string `json:"costReductionInMonth1VsMonth10" form:"costReductionInMonth1VsMonth10" gorm:"column:cost_reduction_in_month1_vs_month10;comment:1月VS10月降本;size:255;"`  //1月VS10月降本 
      CostReductionInMonth2VsMonth10  string `json:"costReductionInMonth2VsMonth10" form:"costReductionInMonth2VsMonth10" gorm:"column:cost_reduction_in_month2_vs_month10;comment:2月VS10月降本;size:255;"`  //2月VS10月降本 
      CostReductionInMonth3VsMonth10  string `json:"costReductionInMonth3VsMonth10" form:"costReductionInMonth3VsMonth10" gorm:"column:cost_reduction_in_month3_vs_month10;comment:3月VS10月降本;size:255;"`  //3月VS10月降本 
      CostReductionInMonth4VsMonth10  string `json:"costReductionInMonth4VsMonth10" form:"costReductionInMonth4VsMonth10" gorm:"column:cost_reduction_in_month4_vs_month10;comment:4月VS10月降本;size:255;"`  //4月VS10月降本 
      CostReductionInMonth5VsMonth10  string `json:"costReductionInMonth5VsMonth10" form:"costReductionInMonth5VsMonth10" gorm:"column:cost_reduction_in_month5_vs_month10;comment:5月VS10月降本;size:255;"`  //5月VS10月降本 
      CostReductionInMonth6VsMonth10  string `json:"costReductionInMonth6VsMonth10" form:"costReductionInMonth6VsMonth10" gorm:"column:cost_reduction_in_month6_vs_month10;comment:6月VS10月降本;size:255;"`  //6月VS10月降本 
      CostReductionInMonth7VsMonth10  string `json:"costReductionInMonth7VsMonth10" form:"costReductionInMonth7VsMonth10" gorm:"column:cost_reduction_in_month7_vs_month10;comment:7月VS10月降本;size:255;"`  //7月VS10月降本 
      CostReductionInMonth8VsMonth10  string `json:"costReductionInMonth8VsMonth10" form:"costReductionInMonth8VsMonth10" gorm:"column:cost_reduction_in_month8_vs_month10;comment:8月VS10月降本;size:255;"`  //8月VS10月降本 
      CostReductionInMonth9VsMonth10  string `json:"costReductionInMonth9VsMonth10" form:"costReductionInMonth9VsMonth10" gorm:"column:cost_reduction_in_month9_vs_month10;comment:9月VS10月降本;size:255;"`  //9月VS10月降本 
      CostReductionInMonth10VsMonth10  string `json:"costReductionInMonth10VsMonth10" form:"costReductionInMonth10VsMonth10" gorm:"column:cost_reduction_in_month10_vs_month10;comment:10月VS10月降本;size:255;"`  //10月VS10月降本 
      CostReductionInMonth11VsMonth10  string `json:"costReductionInMonth11VsMonth10" form:"costReductionInMonth11VsMonth10" gorm:"column:cost_reduction_in_month11_vs_month10;comment:11月VS10月降本;size:255;"`  //11月VS10月降本 
      CostReductionInMonth12VsMonth10  string `json:"costReductionInMonth12VsMonth10" form:"costReductionInMonth12VsMonth10" gorm:"column:cost_reduction_in_month12_vs_month10;comment:12月VS10月降本;size:255;"`  //12月VS10月降本 
      MonthlyCumulativeWeightedPrincipalReductionAmount  string `json:"monthlyCumulativeWeightedPrincipalReductionAmount" form:"monthlyCumulativeWeightedPrincipalReductionAmount" gorm:"column:monthly_cumulative_weighted_principal_reduction_amount;comment:1-当前月累计加权降本金额;size:255;"`  //1-当前月累计加权降本金额 
      MonthlyCumulativeWeightedCostReductionCompletionDegree  string `json:"monthlyCumulativeWeightedCostReductionCompletionDegree" form:"monthlyCumulativeWeightedCostReductionCompletionDegree" gorm:"column:monthly_cumulative_weighted_cost_reduction_completion_degree;comment:1-当前月累计加权降本完成度;size:255;"`  //1-当前月累计加权降本完成度 
      CreatedName  string `json:"createdName" form:"createdName" gorm:"column:created_name;comment:创建人;size:255;"`  //创建人 
      SheetName  string `json:"sheetName" form:"sheetName" gorm:"column:sheet_name;comment:手工表名称;size:255;"`  //手工表名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 主要产品成本降本汇总表 WmsLogisticsPcsSummaryMi自定义表名 wms_logistics_pcs_summary_mi
func (WmsLogisticsPcsSummaryMi) TableName() string {
  return "wms_logistics_pcs_summary_mi"
}

