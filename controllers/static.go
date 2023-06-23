package controllers

import "askvart.com/goals/views"



type Static struct {
	Home 	*views.View
	Contact *views.View
	Faq 	* views.View
	About 	*views.View
}

func NewStatic() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
		Faq: views.NewView("bulma", "static/faq"),
		About: views.NewView("tailwind", "static/about"),
		}		
}
