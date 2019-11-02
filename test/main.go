package main

import (
	"encoding/json"
	"log"
	"net/http"

	//"bitbucket.org/emicklei/dollar"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/recognize", doRecognize)
	log.Println("dollar test available on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func doRecognize(w http.ResponseWriter, req *http.Request) {
	// read points from the request body
	var list [][]float64
	json.NewDecoder(req.Body).Decode(&list)
	defer req.Body.Close()
	log.Printf("%#v", list)
}
