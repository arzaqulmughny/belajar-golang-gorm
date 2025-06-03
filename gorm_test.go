package belajargolanggorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/belajar_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db;
} 

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("INSERT INTO samples(name) VALUES (?)", "Arza").Error;
	assert.Nil(t, err)
}

type Sample struct {
	Id int
	Name string
}

func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("SELECT id, name FROM samples WHERE name = ?", "Arza").Scan(&sample).Error

	assert.Nil(t, err)
	assert.Equal(t, "Arza", sample.Name)
}