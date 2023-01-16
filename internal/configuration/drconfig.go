package configuration

import (
	"context"

	"gorm.io/gorm"
)

// RecoveryScenario - contains the DR config registered
type RecoveryScenario struct {
	gorm.Model
	AppCode         string `json:"appcode"`
	ApplicationName string `json:"application_name"`
	Environment     string `json:"environment"`
}

// Store - defines the interface to interact with RecoveryScenarios
type RecoveryScenarioStore interface {
	GetRecoveryScenarios(ctx context.Context) ([]RecoveryScenario, error)
	// GetRecoveryScenarioByID(ctx context.Context, id string) (RecoveryScenario, error)
	InsertRecoveryScenario(ctx context.Context, rs RecoveryScenario) (RecoveryScenario, error)
	// DeleteRecoveryScenario(id string) error
	UpdateRecoveryScenario(ctx context.Context, rs RecoveryScenario) (RecoveryScenario, error)
}

// Service - responsible for interacting with DB operations + wraps our interface
type Service struct {
	Store RecoveryScenarioStore
}

// New - returns new instance of our recoveryScenario service
// Principle - "accept interface return structs"
func New(store RecoveryScenarioStore) *Service {
	return &Service{
		Store: store,
	}
}

// // GetRecoveryScenarios - list all scenarios present in database
// func (s *Service) GetRecoveryScenarios(ctx context.Context) ([]RecoveryScenario, error) {
// 	scenarios, err := s.Store.GetRecoveryScenarios(ctx)
// 	if err != nil {
// 		return []RecoveryScenario{}, err
// 	}
// 	return scenarios, nil
// }

// // InsertRecoveryScenario - insert a single recovery scenario in database
// func (s *Service) InsertRecoveryScenario(ctx context.Context, rs RecoveryScenario) (RecoveryScenario, error) {
// 	scenario, err := s.Store.InsertRecoveryScenario(ctx, rs)
// 	if err != nil {
// 		return RecoveryScenario{}, err
// 	}
// 	return scenario, nil
// }
