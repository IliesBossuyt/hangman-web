package main

import (
	"fmt"
	"net/http"

)

func main() {
	http.Handle("/", http.FileServer(http.Dir("template")))
	fmt.Println("Serveur web démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
