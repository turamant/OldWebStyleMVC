package main

import (
	"fmt"
	"net/http"

	"askvart.com/goals/controllers"
	"askvart.com/goals/models"
	"github.com/julienschmidt/httprouter"
	//"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "spartak"
	password = "spartak"
	dbname   = "mygolang1"
)



func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()
	
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)
	//usersD := controllers.DestrDB(us)

	router := httprouter.New()
	router.GET("/", staticC.Home.ServeHTTP)
	router.GET("/contact", staticC.Contact.ServeHTTP)
	router.GET("/faq", staticC.Faq.ServeHTTP)
	router.GET("/about", staticC.About.ServeHTTP)
    router.GET("/users", usersC.List)
	router.GET("/users/:id", usersC.UserID)
	
	router.GET("/signup", usersC.SignUp)
	router.POST("/signup", usersC.Create)
	
	router.GET("/login", usersC.LoginView.ServeHTTP)
	router.POST("/login", usersC.Login)

	router.GET("/cookietest", usersC.CookieTest)

	router.GET("/droptable", usersC.DropTable)
	

	http.ListenAndServe(":3000", router)
}
