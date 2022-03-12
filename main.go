package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/* func handlerFunc(w http.ResponseWriter, req *http.Request) {
	} else {
		w.WriteHeader(http.StatusNotFound)
		// if WriteHeader is not called before Write(), the first call to Write
		// will trigger implicit WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<h1>We couldn't find the page you are looking for :(</h1>" +
			"<p>Please email us if you keep being sent to an invalid page</p>")
	}
} */

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to my site!</h1>")
}

func contact(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "To get in touch, please send an email to " +
		"<a href=\"mailto:support@lenslocked.com\">" +
		"support@lenslocked.com</a>.")
}

func main() {
	r := mux.NewRouter()  // like we do with http.NewServeMux()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq_helper)
	http.ListenAndServe(":3000", r)
}


/** Things to checkout:
 *  Dynamic reloading: https://github.com/pilu/fresh - Box3.2
 */
