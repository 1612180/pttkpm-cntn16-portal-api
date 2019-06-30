package storage

import (
	"os"

	"github.com/jinzhu/gorm"
)

func MigrateAll(db *gorm.DB) {
	if os.Getenv("DATABASE_MODE") == "debug" {
		db.DropTableIfExists(&Account{})
		db.DropTableIfExists(&Student{})

		db.DropTableIfExists(&Program{})
		db.DropTableIfExists(&Faculty{})

		db.DropTableIfExists(&Subject{})
		db.DropTableIfExists(&TypeSub{})

		db.DropTableIfExists(&Enroll{})
		db.DropTableIfExists(&Score{})
	}

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Student{})

	db.AutoMigrate(&Program{})
	db.AutoMigrate(&Faculty{})

	db.AutoMigrate(&Subject{})
	db.AutoMigrate(&TypeSub{})

	db.AutoMigrate(&Enroll{})
	db.AutoMigrate(&Score{})
}
