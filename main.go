package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	lang := p.ByName("lang")
	fmt.Fprintf(w, lang)
}

func main() {
	router := httprouter.New()
	router.GET("/Hello/:lang", Hello)

	err := http.ListenAndServe(":8080", router)
	if err != nil{
		log.Fatal(err)
	}
}
