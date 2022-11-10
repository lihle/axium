package handler

import (
	"net/http"
)

//Viewlogin : login page
func Viewlogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login/index.html")
}
