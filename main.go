package main

import (
	"fmt"
	"net/http"

	//"github.com/gorilla/mux"
	"askvart.com/goals/controllers"
	"github.com/julienschmidt/httprouter"
)







func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//r := mux.NewRouter()
	//r.NotFoundHandler = http.HandlerFunc(page404)
	//r.HandleFunc("/", home)
	//r.HandleFunc("/contact", contact)
	//r.HandleFunc("/faq", faq)
	//http.ListenAndServe(":3000", r)

	
	  
	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()
	fmt.Printf("%T...%v",usersC, usersC)

	router := httprouter.New()
		
	router.GET("/", staticC.Home.ServeHTTP)
	router.GET("/contact", staticC.Contact.ServeHTTP)
	router.GET("/faq", staticC.Faq.ServeHTTP)
	router.GET("/about", staticC.About.ServeHTTP)

	router.GET("/signup", usersC.New) 
	router.POST("/signup", usersC.Create)

	http.ListenAndServe(":3000", router)
}
