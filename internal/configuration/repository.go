package configuration

// import (
// 	"context"

// 	"gorm.io/gorm"
// )

// type Repository interface {
// 	Migrate(ctx context.Context) error
// 	All(ctx context.Context) ([]RecoveryScenario, error)
// 	GetByID(ctx context.Context, id uint) (RecoveryScenario, error)
// }

// type GORMRepository struct {
// 	db *gorm.DB
// }

// // NewGORMRepository - create a new client to interact with DB operations
// func NewGORMRepository(db *gorm.DB) *GORMRepository {
// 	return &GORMRepository{
// 		db: db,
// 	}
// }

// // Migrate - responsible for migrating schema
// func (r *GORMRepository) Migrate(ctx context.Context) error {
// 	m := &RecoveryScenario{}
// 	return r.db.WithContext(ctx).AutoMigrate(&m)
// }

// // GetByID - get only single scenario by ID
// func (r *GORMRepository) GetByID(ctx context.Context, id uint) (RecoveryScenario, error) {
// 	var rs RecoveryScenario

// 	// set ID on lookup object
// 	rs.ID = id

// 	err := r.db.WithContext(ctx).Model(&RecoveryScenario{}).Find(&rs).Error
// 	if err != nil {
// 		return rs, err
// 	}

// 	return rs, err
// }

// // All - get all scenarios from database
// func (r *GORMRepository) All(ctx context.Context) ([]RecoveryScenario, error) {
// 	var recoveryScenarios []RecoveryScenario

// 	// get all from DB
// 	err := r.db.WithContext(ctx).Model(&RecoveryScenario{}).Find(&recoveryScenarios).Error
// 	if err != nil {
// 		return recoveryScenarios, err
// 	}

// 	return recoveryScenarios, nil
// }
