package start

import "github.com/models"

func SyncDB() {
	DB.AutoMigrate(&models.Customer{}, &models.Book{}, &models.Order{})
}
