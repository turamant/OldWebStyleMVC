package main

import (
	"net/http"

	//"github.com/gorilla/mux"
	"askvart.com/goals/controllers"
	"askvart.com/goals/views"
	"github.com/julienschmidt/httprouter"
)

var (
	homeView    *views.View
	contactView *views.View
	faqView     *views.View
	aboutView   *views.View

)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}
func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	must(aboutView.Render(w, nil))
}



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

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bulma", "views/faq.gohtml")
	aboutView = views.NewView("tailwind", "views/about.gohtml")
    usersC := controllers.NewUsers()

	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("/static/"))
	router.GET("/", home)
	router.GET("/contact", contact)
	router.GET("/faq", faq)
	router.GET("/about", about)
	router.GET("/signup", usersC.New) //
	http.ListenAndServe(":3000", router)
}
