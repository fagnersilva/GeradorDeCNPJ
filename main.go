package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getCPNJ(w http.ResponseWriter, r *http.Request) {

	cnpj := arrayCpnj()
	newResponse := map[string][3]string{"data": cnpj}
	json.NewEncoder(w).Encode(newResponse)

}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/cnpj", getCPNJ)
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
