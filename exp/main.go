package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "usegolang"
	password = "usegolang"
	dbname   = "lenslocked_dev"
)

// User is an existential concept and a GORM model
type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null; unique_index"`
	Orders []Order
}

// Order is a gorm model for things a user paid for
type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})

	var user User
	db.Preload("Orders").First(&user)
	if db.Error != nil {
		panic(db.Error)
	}

	fmt.Println("Email:", user.Email)
	fmt.Println("Number of orders:", len(user.Orders))
	fmt.Println("Orders:", user.Orders)
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("What is your email?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return name, email
}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	})
	if db.Error != nil {
		panic(db.Error)
	}
}
