package handler

import (
	"net/http"
)

//Viewhome : home page
func Viewhome(w http.ResponseWriter, r *http.Request) {

	view(w, newPage("Home page", ""))
}
