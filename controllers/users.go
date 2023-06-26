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
	NewView *views.View
	us 		*models.UserService
}

type SignupForm struct{
	Name 	 string `schema:"name"`
	Age      uint   `schema:"age"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers(us *models.UserService) *Users{
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us: us,
	}
}

func ListU(us *models.UserService) *Users{
	return &Users{
		NewView: views.NewView("bootstrap", "users/list"),
		us: us,
	}
}

func UsId(us *models.UserService) *Users{
	return &Users{
		NewView: views.NewView("bootstrap", "users/user"),
		us: us,
	}
}




func (u *Users) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	users, err := u.us.ListUsers()
	if err != nil {
		panic(err)
	}
	if err := u.NewView.Render(w, users); err != nil{
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
	if err := u.NewView.Render(w, user); err != nil{
		panic(err)
	}
	
}



func (u *Users) New(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	if err := u.NewView.Render(w, nil); err != nil{
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
	}
	if err := u.us.Create(&user); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "User is", user)
}