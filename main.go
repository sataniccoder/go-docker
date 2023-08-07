package main

import (
	"fmt"
	"net/http"
	"os"

	gets "webz/mod/get"
)

func banner() {
	dock_stat := "False"
	if _, err := os.Stat("/.dockerenv"); err == nil {
		dock_stat = "True"
	}

	fmt.Println(`
	█████████████████████████████
	█▄ █▀▀▀█ ▄█▄ ▄▄ █▄ ▄ ▀█ ▄▄ ▄█
	██ █ █ █ ███ ▄█▀██ ▄ ▀██▀▄█▀█
	██▄▄▄█▄▄▄██▄▄▄▄▄█▄▄▄▄██▄▄▄▄▄█

	URL    -> http://127.0.0.1
	DOCKER -> [` + dock_stat + `]
	
	`)
}

func main() {

	banner()

	http.HandleFunc("/", gets.Get_handler)

	http.ListenAndServe(":8080", nil)
}
