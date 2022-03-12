package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if req.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Welcome to my site!</h1>")
	} else if req.URL.Path == "/contact" {
		fmt.Fprintf(w, "To get in touch, please send an email to " +
			"<a href=\"mailto:support@lenslocked.com\">" +
			"support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		// if WriteHeader is not called before Write(), the first call to Write
		// will trigger implicit WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<h1>We couldn't find the page you are looking for :(</h1>" +
			"<p>Please email us if you keep being sent to an invalid page</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}


/** Things to checkout:
 *  Dynamic reloading: https://github.com/pilu/fresh - Box3.2
 */
