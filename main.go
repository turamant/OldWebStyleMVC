package main

import (
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

	//homeView = views.NewView("bootstrap", "views/home.gohtml")
	//contactView = views.NewView("bootstrap", "views/contact.gohtml")
	  
	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()

	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("/static/"))
	router.GET("/", staticC.Home.ServeHTTP)
	router.GET("/contact", staticC.Contact.ServeHTTP)
	router.GET("/faq", staticC.Faq.ServeHTTP)
	router.GET("/about", staticC.About.ServeHTTP)
	router.GET("/signup", usersC.New) 
	router.POST("/signup", usersC.Create)
	http.ListenAndServe(":3000", router)
}
