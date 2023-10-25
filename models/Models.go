package models

type Order struct {
	Order_Number uint `gorm:"primaryKey;autoIncrement"`
	Customer_ID  uint
	Customer     Customer `gorm:"foreignKey:Customer_ID"`
	Total_Price  uint
}

type Book struct {
	Book_ID     uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"not null;unique"`
	Author      string
	Price       uint `gorm:"not null"`
	Quantity    uint `gorm:"not null"`
	Customer_ID uint
	Customer    Customer `gorm:"foreignKey:Customer_ID"`
}

type Customer struct {
	Customer_ID uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"primaryKey;not null"`
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Phone       uint   `gorm:"not null; unique"`
	Address     string `gorm:"not null"`
}
