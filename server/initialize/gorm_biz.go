package initialize

import (
	"gorm.io/gorm"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tally"
)

func bizModel(db *gorm.DB) error {
	return db.AutoMigrate(tally.TallyBill{})
}
