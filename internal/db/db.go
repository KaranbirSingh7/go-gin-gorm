package db

import (
	"os"

	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB // only this package can access this field
	logger zerolog.Logger
}

// New - initiates a new DB connection
func New() (Database, error) {
	// create new db if not exists
	conn, err := gorm.Open(sqlite.Open("gormdummy.db"), &gorm.Config{})
	if err != nil {
		return Database{}, err
	}

	return Database{
		Client: conn,
		logger: zerolog.New(os.Stdout),
	}, nil
}

func (d *Database) MigrateSchemas(m ...interface{}) error {
	for _, s := range m {
		err := d.Client.AutoMigrate(&s)
		if err != nil {
			return err
		}
	}
	d.logger.Print("schemas migrated successfully.")
	return nil
}
