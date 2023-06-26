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
	//r := mux.NewRouter()
	//r.NotFoundHandler = http.HandlerFunc(page404)
	//r.HandleFunc("/", home)
	//r.HandleFunc("/contact", contact)
	//r.HandleFunc("/faq", faq)
	//http.ListenAndServe(":3000", r)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()
	  
	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()

	router := httprouter.New()
		
	router.GET("/", staticC.Home.ServeHTTP)
	router.GET("/contact", staticC.Contact.ServeHTTP)
	router.GET("/faq", staticC.Faq.ServeHTTP)
	router.GET("/about", staticC.About.ServeHTTP)

	router.GET("/signup", usersC.New) 
	router.POST("/signup", usersC.Create)

	http.ListenAndServe(":3000", router)
}
