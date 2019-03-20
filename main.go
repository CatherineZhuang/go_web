// main: first thing to run
package main

// pulling in libraries to reuse code
// fmt: print
// http: create web server, handling web requests and create web requests
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)



// can be renamed, takes in two args
// w: what is sent back to users
// r: what we want to send to server
// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h1>Welcome to my new site</h1>")
// 	} else if r.URL.Path == "/contact" {
// 		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h1>We could not find the page you are looking for</h1><p>Please send us a email</p>")
// 	}
// }

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my new site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//     fmt.Fprint(w, "Welcome!\n")
// }
// // ps: specific to http router
// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//     fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

// main function
// tell http package how to handle requests
// /: match prefix
func main() {
	// gorilla toolkit
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	//julien schmit
	// router := httprouter.New()
    // router.GET("/", Index)
    // router.GET("/hello/:name", Hello)

	// self defined
	// mux := &http.ServeMux{}
	// // always match the longer one, order does not matter
	// // / match to every single path
	// mux.HandleFunc("/", handlerFunc)
	// http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", r)
}