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
		db.DropTableIfExists(&RequireSub{})

		db.DropTableIfExists(&Enroll{})
		db.DropTableIfExists(&Score{})
		db.DropTableIfExists(&TryEnroll{})
	}

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Student{})

	db.AutoMigrate(&Program{})
	db.AutoMigrate(&Faculty{})

	db.AutoMigrate(&Subject{})
	db.AutoMigrate(&TypeSub{})
	db.AutoMigrate(&RequireSub{})

	db.AutoMigrate(&Enroll{})
	db.AutoMigrate(&Score{})
	db.AutoMigrate(&TryEnroll{})

	db.Create(&Program{Short: "cntn", Long: "Cử nhân tài năng"})
	db.Create(&Program{Short: "cq", Long: "Chính quy"})
	db.Create(&Faculty{Short: "cntt", Long: "Công nghệ thông tin"})
	db.Create(&Faculty{Short: "sh", Long: "Sinh học"})
	db.Create(&TypeSub{Short: "bb", Long: "Bắt buộc"})
	db.Create(&TypeSub{Short: "tc", Long: "Tự chọn"})
}
