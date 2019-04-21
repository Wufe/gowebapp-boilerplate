package infrastructure

import (
	"fmt"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Sqlite
	"github.com/wufe/boilerplateprj/models"
)

// Database gives access to main database functions
type DatabaseAccessor interface {
	Connect()
	Close()
	Automigrate()
	Seed()
	DB() *gorm.DB
}

type databaseAccessorImpl struct {
	db *gorm.DB
}

func NewDatabase() *databaseAccessorImpl {
	return &databaseAccessorImpl{}
}

func (database *databaseAccessorImpl) DB() *gorm.DB {
	return database.db
}

// Connect - Initializes connection to the database
func (database *databaseAccessorImpl) Connect() {
	dbPath, _ := filepath.Abs("database/db.db")
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("Failed to connect database, reason: " + err.Error() + " " + dbPath)
	}
	database.db = db
}

// Close - Closes the connection to the database
func (database *databaseAccessorImpl) Close() {
	database.db.Close()
}

// Automigrate - Automatically creates database schemas
func (database *databaseAccessorImpl) Automigrate() {
	database.db.AutoMigrate([]interface{}{
		&models.Role{},
		&models.User{},
	}...)
}

// Seed - Creates required records if not present
func (database *databaseAccessorImpl) Seed() {

	users := []*models.User{
		&models.User{
			Email:    "root@example.dev",
			Password: "toor",
			Roles: []*models.Role{
				&models.Role{
					Description: "Admin",
					Name:        "ADMIN",
				},
			},
		},
	}

	for _, user := range users {
		notFound := database.db.First(&models.User{}, "email = ?", user.Email).RecordNotFound()
		if notFound {
			fmt.Printf("User %s not found: creating..\n", user.Email)
			database.db.Create(user)
		}
	}
}
