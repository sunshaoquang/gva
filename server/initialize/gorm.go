package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manual"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	ht_smartdata := global.GetGlobalDBByDBName("ht_smartdata")
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysChatGptOption{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{}, manual.PromotionMarketTarget{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
	ht_smartdata.AutoMigrate(manual.WmsKpi2023SummaryMi{}, manual.WmsKpiSummary2Mi{}, manual.WmsLogisticsPcsDetailMi{}, manual.WmsLogisticsPcsSummaryMi{}, manual.WmsLogisticsPcs2023Mi{}, manual.WmsLogisticsRateMonths{}, manual.WmsLogisticsProviderMi{}, manual.WmsLogisticsModeTransportMi{}, manual.WmsLogisticsSalebycountryMi{}, manual.WmsLogisticsStoragefeeMi{}, manual.OmsCnSalesTargetDaily{}, manual.OmsCnChannelSales{})
}
