package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getCPNJ(w http.ResponseWriter, r *http.Request) {

	cnpj := arrayCpnj()
	newResponse := map[string][3]string{"data": cnpj}
	json.NewEncoder(w).Encode(newResponse)

}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/cnpj", getCPNJ)
	log.Println("Executando ...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
