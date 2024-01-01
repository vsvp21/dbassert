package dbassert

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

type TestModel struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func TestDBHas(t *testing.T) {
	dsn := "host=localhost user=db_user password=secretsecret dbname=factory port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&TestModel{})
	if err != nil {
		log.Fatal(err)
	}

	m := &TestModel{Name: "testName", Age: 50}
	db.Model(&TestModel{}).Create(&m)

	DBHas(t, db, &m)

	err = db.Migrator().DropTable(&TestModel{})
	if err != nil {
		log.Fatal(err)
	}
}

func TestDBHasNot(t *testing.T) {
	dsn := "host=localhost user=db_user password=secretsecret dbname=factory port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&TestModel{})
	if err != nil {
		log.Fatal(err)
	}

	DBHasNot(t, db, &TestModel{ID: 1000000})

	err = db.Migrator().DropTable(&TestModel{})
	if err != nil {
		log.Fatal(err)
	}
}
