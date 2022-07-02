package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RR2609/go-microservices.git/details"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking Application Health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is UP and running")

}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching details")
	hostname, err := details.GetHostName()
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIP()
	fmt.Println(hostname, IP)
	response := map[string]string{
		"Hostname": hostname,
		"IP":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	log.Println("Server has Started")
	log.Fatal(http.ListenAndServe(":80", r))

}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func roothandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s \n", r.URL.Path, r.URL.Query().Get("token"))

// }

// func main() {

// 	http.HandleFunc("/", roothandler)

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println(("Web Server has started!!"))
// 	http.ListenAndServe(":80", nil)
// }

// package main

// import (
// 	"fmt"

// 	"rsc.io/quote"
// )

// func rect(length, width float64) (area, perimeter float64) {
// 	area = length * width
// 	perimeter = 2 * (length + width)
// 	return
// }

// func main() {
// 	fmt.Println("Hello, World!")
// 	fmt.Println(quote.Go())

// 	a, p := rect(1, 2)
// 	fmt.Printf(" Area is %f and perimeter is %f", a, p)
// }
