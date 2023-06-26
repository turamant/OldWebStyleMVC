package controllers

import (
	"fmt"
	"net/http"
	"askvart.com/goals/views"
	"github.com/julienschmidt/httprouter"
)

type Users struct{
	NewView *views.View
}

type SignupForm struct{
	Name 	 string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers() *Users{
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
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
	fmt.Fprintln(w, "Name is", form.Name)
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprint(w, "Password is", form.Password)
}