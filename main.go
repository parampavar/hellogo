package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/parampavar/helloGo/Godeps/_workspace/src/github.com/cihub/seelog"
)

// const (
//     DB_USER     = "postgres"
//     DB_PASSWORD = "postgres"
//     DB_LOCATION = "localhost"
//     DB_NAME     = "postgres"
//     DB_SSLMODE  = "disable"
// )

const (
	DB_USER     = "u311d07be533d42da8c704a4c29f0d573"
	DB_PASSWORD = "c9e75db43e744176a5970138c3b7f080"
	DB_LOCATION = "10.72.6.110:5432"
	DB_NAME     = "d311d07be533d42da8c704a4c29f0d573"
	DB_SSLMODE  = "disable" //verify-full"
)

const (
	DEFAULT_PORT = "9000"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("HomeHandler Starting")
	fmt.Fprintln(w, "Hello, GO World!n")
	log.Info("HomeHandler Ending")
}

func main() {
	// defer log.Flush()
	log.Info("App Started")
	fmt.Println("App started")

	// router := mux.NewRouter()
	// router.HandleFunc("/", HomeHandler)
	// router.HandleFunc("/db", DBHandler)
	// // Bind to a port and pass our router in
	// http.ListenAndServe(":8000", nil)

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Info("Warning, PORT not set. Defaulting to %+vn", DEFAULT_PORT)
		port = DEFAULT_PORT
	}
	//	port = DEFAULT_PORT

	http.HandleFunc("/", HomeHandler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Info("ListenAndServe: ", err)
	}

}
