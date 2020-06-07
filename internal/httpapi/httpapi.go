package httpapi

import (
	"fmt"
	"log"
	"net/http"
)

// TODO:
// move call to pred.py here

func newbird(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("get?")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseMultipartForm(5 * 1024 * 1024); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostForm = %v\n", r.Form)
		filename := r.FormValue("filename")
		info := r.FormValue("info")
		fmt.Fprintf(w, "fn = %s\n", filename)
		fmt.Fprintf(w, "info = %s\n", info)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Start() {
	http.HandleFunc("/newbird", newbird)
	fmt.Println("Starting http server")
	if err := http.ListenAndServe(":8880", nil); err != nil {
		log.Fatal(err)
	}
}
