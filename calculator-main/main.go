package main

import (
	"calculator-main/pkg/handlers"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
	"html/template"
	"./handlers" // Import the handlers package
)

var store = sessions.NewCookieStore([]byte("pass"))
var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/", handlers.MainHandler) // Use handlers package functions
	http.HandleFunc("/calc", handlers.CalcHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
