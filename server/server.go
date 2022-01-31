package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"testProgect/cash"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./UI/template/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

func orderPage(cash *cash.Cash) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}
		obj, resul := cash.GetItemCashById(id)
		if resul == true {
			ts, err := template.ParseFiles("./UI/template/order.html")
			if err != nil {
				log.Println(err.Error())
				http.Error(w, "Internal Server Error", 500)
				return
			}

			err = ts.Execute(w, obj)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, "Internal Server Error", 500)
			}

		}
	}
}
func RunServer(cash *cash.Cash) {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", orderPage(cash))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
