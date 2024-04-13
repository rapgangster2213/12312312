package handlers

import (
	"./calculator-main/pkg/keygen"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"pkg/mathparse"
	"strconv"
)

type Calculation struct {
	Expression string
	Result     float64
	Session    string
}

var calculations []Calculation

func MainHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get("session", r)
	username := session.Values["username"]
	if username == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/calc", http.StatusSeeOther)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["username"] = keygen.RandStr()
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates.Execute(w, "calc.html")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		session, err := store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rpn := mathparse.RpnGo{}
		untypedUsername, ok := session.Values["username"]
		if !ok {
			return
		}
		username, ok := untypedUsername.(string)
		if !ok {
			return
		}
		session.Values["expression"] = r.FormValue("expression")
		untypedExpression, ok := session.Values["expression"]
		if !ok {
			return
		}
		expr, _ := untypedExpression.(string)
		expression, ok := untypedExpression.(string)
		if !ok {
			return
		}
		if expression[0] == '-' {
			expression = "0-" + expression[1:]
		}
		rpn.CalculateExpression(expression)
		result := rpn.GetResult()

		session.Values["result"] = result

		calculations = append(calculations, Calculation{
			Expression: expr,
			Result:     result,
			Session:    username,
		})

		fmt.Fprintf(w, "<p>Result: %s<p>", strconv.FormatFloat(result, 'f', -1, 64)) // result output
		templates.Execute(w, calculations)
		for _, calc := range calculations {
			if username == calc.Session {
				fmt.Fprintf(w, "<p>%s = %s", calc.Expression, strconv.FormatFloat(calc.Result, 'f', -1, 64))              // calc history output
				fmt.Printf("%s = %s, key: %s\n", expression, strconv.FormatFloat(calc.Result, 'f', -1, 64), calc.Session) // bebra
			}
		}
	}
}
