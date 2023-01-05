package app

import (
	"net/http"

	"github.com/carballoandrea/gingonicAPI/controllers"
)

func StartApp(){
	http.HandleFunc("/users", controllers.GetUser)
}