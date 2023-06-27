package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"askvart.com/goals/models"
	"askvart.com/goals/views"
	"github.com/julienschmidt/httprouter"
)

type Users struct{
	SignUpView 	*views.View
	ListView 	*views.View
	UserView    *views.View
	LoginView   *views.View
	us 			*models.UserService
}

type SignupForm struct{
	Name 	 	string `schema:"name"`
	Age     	uint   `schema:"age"`
	Email    	string `schema:"email"`
	Password 	string `schema:"password"`
}

type LoginForm struct {
	Email 		string `schema:"email"`
	Password 	string `schema:"password"`
}

func NewUsers(us *models.UserService) *Users{
	return &Users{
		SignUpView: views.NewView("bootstrap", "users/signup"),
		ListView: views.NewView("bootstrap", "users/list"),
		UserView:  views.NewView("bootstrap", "users/user"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us: us,
	}
}


func (u *Users) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	users, err := u.us.ListUsers()
	if err != nil {
		panic(err)
	}
	if err := u.ListView.Render(w, users); err != nil{
		panic(err)
	}
	
}
func (u *Users) UserID(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	s := ps.ByName("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	user, err := u.us.ByID(id)
	if err != nil {
		panic(err)
	}
	if err := u.UserView.Render(w, user); err != nil{
		panic(err)
	}
	
}

func (u *Users) SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	if err := u.SignUpView.Render(w, nil); err != nil{
		panic(err)
	}
}



func (u *Users) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	var form SignupForm
	if err := parseForm(r, &form); err != nil{
		panic(err)
	}
	user := models.User{
		Name: 	form.Name,
		Age:   	form.Age,
		Email: 	form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusFound)
}

func (u *Users) DropTable(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	err := u.us.DestructiveReset()
	if err != nil {
		http.Error(w, "error drop table", 500)
	}
	http.Redirect(w, r, "/users", http.StatusFound)
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
	panic(err)
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
	switch err {
	case models.ErrNotFound:
		//fmt.Fprintln(w, "Invalid email address")
		http.Redirect(w, r, "/login", http.StatusFound)
	case models.ErrInvalidPassword:
		http.Redirect(w, r, "/login", http.StatusFound)
		//fmt.Fprintln(w, "Invalid password provided")
	case nil:
		//fmt.Fprintln(w, user)
		http.Redirect(w, r, "/users", http.StatusFound)
	default:
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	return
	}
	cookie := http.Cookie{
		Name: "email",
		Value: user.Email,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, user)
}

func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie, err := r.Cookie("email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Email is: ", cookie.Value)
}