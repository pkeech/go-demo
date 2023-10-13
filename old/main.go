// NAME PACKAGE
package main

// IMPORT REQUIRED PACKAGES
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

// DEFINE VARIABLES
var tmpl *template.Template

// STRUCT: HOME PAGE DATA
type HomePage struct {
	Message     string `json:"message"`
	Environment string `json:"environment"`
	API         string `json:"api"`
}

// STRUCTURE: BOOKS
type Book struct {
	ISDN  int    `json:"isdn"`
	Title string `json:"title"`
	Pages int    `json:"pages"`
}

// ROUTE: INDEX PAGE
func Index(w http.ResponseWriter, r *http.Request) {
	// SET HEADER
	w.Header().Set("Content-Type", "text/html")

	// GENERATE RESPONSE
	w.Write([]byte("<h1>Hello World</h1>"))
}

// ROUTE: HOME PAGE
func Home(w http.ResponseWriter, r *http.Request) {
	// SET HEADER
	w.Header().Set("Content-Type", "text/html")

	// GENERATE PAGE DATA
	homeData := HomePage{
		Message:     "Hello World !!",
		Environment: os.Getenv("ENV"),
		API:         os.Getenv("API_KEY"),
	}

	// RENDER TEMPLATE
	tmpl.Execute(w, homeData)
}

// ROUTE: API PAGE
func Api(w http.ResponseWriter, r *http.Request) {
	// SET HEADER
	w.Header().Set("Content-Type", "application/json")

	// GENERATE FAKE BOOK DATA
	myBook := Book{
		ISDN:  3546543354354,
		Title: "Hitchhikers Guide to the Galaxy",
		Pages: 231,
	}

	// ENCODE JSON OBJECT AND GENERATE RESPONSE
	json.NewEncoder(w).Encode(myBook)
}

// FUNCTION: MAIN APPLICATION FUNCTION
func main() {
	// LOAD ENVIRONMENT FILE
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("[ERROR] Unable to find/load environment file")
		os.Exit(1)
	}

	// INSTANTIATE TEMPLATE
	tmpl = template.Must(template.ParseFiles("templates/index.tmpl"))

	// CREATE STATIC CONTENT FILE SERVER
	fs := http.FileServer(http.Dir("./static"))

	// INSTANTIATE MUX SERVER
	mux := http.NewServeMux()

	// ADD ROUTE(S) TO WEB SERVER
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/home", Home)
	mux.HandleFunc("/api", Api)

	// PUBLISH STATIC FILES
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// START WEB SERVER
	log.Print("Server is Listening on :8080 ...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
