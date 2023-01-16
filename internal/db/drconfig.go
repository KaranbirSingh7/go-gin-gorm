package db

import (
	"context"

	"github.com/karanbirsingh7/go-gin-gorm/internal/configuration"
)

// GetRecoveryScenarios - list all scenarios present in database
func (d *Database) GetRecoveryScenarios(ctx context.Context) ([]configuration.RecoveryScenario, error) {
	var rs []configuration.RecoveryScenario
	// get all from DB
	err := d.Client.WithContext(ctx).Find(&rs).Error
	if err != nil {
		return rs, err
	}

	return rs, nil
}

// InsertRecoveryScenario- creates a new recoveryScenario record in database
func (d *Database) InsertRecoveryScenario(ctx context.Context, rs configuration.RecoveryScenario) (configuration.RecoveryScenario, error) {

	err := d.Client.WithContext(ctx).Model(&configuration.RecoveryScenario{}).Create(&rs).Error
	if err != nil {
		return configuration.RecoveryScenario{}, err
	}

	return rs, nil
}

// UpdateRecoveryScenario - updates an existing recovery scenario
func (d *Database) UpdateRecoveryScenario(ctx context.Context, rs configuration.RecoveryScenario) (configuration.RecoveryScenario, error) {
	err := d.Client.Model(&configuration.RecoveryScenario{}).Updates(&rs).Error
	if err != nil {
		return rs, err
	}
	return rs, err
}
