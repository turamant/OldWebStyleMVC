package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "spartak"
	password = "spartak"
	dbname   = "mygolang1_exp"
)

type User struct{
	ID int `gorm:"autoIncrement;primaryKey"`
	Name string
	Email string `gorm:"not null;unique_index"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID uint
	Amount int
	Description string
}

func connectDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
				"password=%s dbname=%s sslmode=disable",
				host, port, user, password, dbname)


	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	
	return db
}

func createUser(db *gorm.DB){
	name, email := getInfo()
	u := &User{
		Name: name,
		Email: email,
	}
	if err := db.Create(u).Error; err != nil{
		panic(err)
	}
	fmt.Printf("%+v\n", u)
}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	db.Create(&Order{
		UserID: uint(user.ID),
	    Amount: amount,
	    Description: desc,
	})
	if db.Error != nil {
		panic(db.Error)
	}
}


func getUserID(db *gorm.DB) User {
	var u User
	ID := getUserIdInfo()
	db.First(&u, ID)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(u)
	return u
}

func getUserWhere(db *gorm.DB){
	var u User
	maxID := getUserIdInfo()
	db.Where("id <= ?", maxID).First(&u)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(u)
}

func preloadingUserOrders(db *gorm.DB, user User){
	db.Preload("Orders").First(&user)
	if db.Error != nil{
		panic(db.Error)
	}
	fmt.Println("Email:", user.Email)
	fmt.Println("Number of orders:", len(user.Orders))
	//fmt.Println("Orders:", user.Orders)
	for k,v := range user.Orders{
		fmt.Println(k,":",v.Amount, v.Description)
	}
}

func listUser(db *gorm.DB){
	var users []User
	db.Find(&users)
    if db.Error != nil{
		panic(db.Error)
	}
	fmt.Println("Retrieved", len(users), "users.")
	for k, v := range users{
		fmt.Println(k, ":", v.ID, v.Name)
	}
}

/*
func deleteUser(db *sql.DB,id int){
	_, err := db.Exec(`
		DELETE
		FROM users
		WHERE id=$1`,id)
	
	if err != nil {
		panic(err)
	}
	fmt.Println("Удален ID:", id)
}

func deleteOrder(db *sql.DB, id int){
	_, err := db.Exec(`
		DELETE
		FROM orders
		WHERE id=$1`,id)
	
	if err != nil {
		panic(err)
	}
	fmt.Println("Удален ID:", id)
}


func createOrders(db *sql.DB){
	var id int
	for i:=0; i<2;i++{
		userId := 1
		if i > 3 {
			userId = 2
		}
		amount := 1000*i
		description := fmt.Sprintf("USB-C Adapter x%d", i)
		
		err := db.QueryRow(`
		INSERT INTO orders (user_id, amount, description)
		VALUES ($1, $2, $3)
		RETURNING id`,
		userId, amount, description).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("created an order with the ID:", id)
	}
}

func listUserOrders(db *sql.DB){
	var user_name, user_email, order_description string
	var user_id, order_id, order_amount int

	rows, err := db.Query(`
		select users.id AS user_id, users.name AS user_name, users.email AS user_email, orders.id AS order_id,
		orders.amount AS order_amount,
		orders.description AS order_description
		from users INNER JOIN orders
		ON users.id = orders.user_id`,
	)

	if err != nil {
		panic(err)
	}
    
	for rows.Next(){
		rows.Scan(&user_id, &user_name, &user_email, &order_id, &order_amount, &order_description)
		fmt.Printf("usId:%d\n, usName:%s, us.Email:%s, or.Id:%d, orAmount:%d, orDescrip:%s\n",
		user_id, user_name, user_email, order_id, order_amount, order_description,)
	}

}
*/

func getInfo()(name, email string){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("What is your email?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return name, email
}

func getUserIdInfo() (ID string){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is user ID?")
	ID, _ = reader.ReadString('\n')
	ID = strings.TrimSpace(ID)
	return ID
}
func main(){
	db := connectDB()
	defer db.Close()
	//insertDB(db, "Petrov", "petrov@mail.ru")
	//getUser(db, 1)
	//listUser(db)
	//deleteUser(db,46)
	//deleteOrder(db, i)
	
	//createOrders(db)
	//listUserOrders(db)
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})
	//createUser(db)
	user := getUserID(db)
	//getUserWhere(db)
	//listUser(db)
	//createOrder(db, user, 337, "Volvo 60c")
	preloadingUserOrders(db, user)

	
}