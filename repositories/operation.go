package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/Mague/forex-bridge/models"
	"github.com/Mague/forex-bridge/utils"
)

type OperationRepository struct{}

func (r *OperationRepository) Create(operations []models.Operation) (tx *gorm.DB) {
	var result *gorm.DB
	utils.Query(func(db *gorm.DB) {
		result = db.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "order_number"}},
			/* DoUpdates: clause.AssignmentColumns([]string{"symbol", "price", "account_id"}), */
			UpdateAll: true,
		}).Create(&operations)
	})
	return result
}

func (r *OperationRepository) All() []models.Operation {
	var result []models.Operation
	utils.Query(func(db *gorm.DB) {
		db.Find(&result)
	})
	return result
}
