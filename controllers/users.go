package controllers

import (
	"net/http"

	"askvart.com/goals/views"
	"github.com/julienschmidt/httprouter"
)

type Users struct{
	NewView *views.View
}

func NewUsers() *Users{
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	if err := u.NewView.Render(w, nil); err != nil{
		panic(err)
	}
}