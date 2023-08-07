package get

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func file_real(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}

	return false
}

func Get_handler(w http.ResponseWriter, r *http.Request) {
	// get the url
	path := r.URL.Path[1:]

	// index url
	if path == "" {
		fmt.Println("[GET] index.html -> client [200]")

		http.ServeFile(w, r, "templates/html/index.html")
	}

	// all other reqs
	fmt.Print("[GET] " + path)
	if strings.Contains(path, ".css") {
		fmt.Print(" -> client ")

		// verify the file exists
		if file_real("templates/css/" + path) {
			fmt.Println(" [200]")
			http.ServeFile(w, r, "templates/css/"+path)
		} else {
			fmt.Println(" [404]")
			w.WriteHeader(http.StatusNotFound)
		}
	} else if strings.Contains(path, "html") {
		fmt.Print(" -> client ")

		// verify the file exists
		if file_real("templates/html/" + path) {
			fmt.Println(" [200]")
			http.ServeFile(w, r, "templates/html/"+path)
		} else {
			fmt.Println(" [404]")
			http.ServeFile(w, r, "templates/html/404.html")
		}
	} else {
		fmt.Println(" -> client [404]")
		http.ServeFile(w, r, "templates/html/404.html")
	}
}
