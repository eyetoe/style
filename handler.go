// Package of functions for the xonk.org site
package style

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// serve: code.css
func Technotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Handler - technotes css")

	css, err := ioutil.ReadFile("/git/GO/src/github.com/eyetoe/webserver/assets/technotes.css")
	check(err)
	fmt.Fprintln(w, string(css))
}

// serve: jabberwebby.css
func Jabberwebby(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Handler - jabberwebby css")

	css, err := ioutil.ReadFile("/git/GO/src/github.com/eyetoe/webserver/assets/jabberwebby.css")
	check(err)
	fmt.Fprintln(w, string(css))
}

// serve: landing.css
func Landing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Handler - landing css")

	css, err := ioutil.ReadFile("/git/GO/src/github.com/eyetoe/webserver/assets/landing.css")
	check(err)
	fmt.Fprintln(w, string(css))
}

// serve: style.css
func Default(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Handler - style css")

	css, err := ioutil.ReadFile("/git/GO/src/github.com/eyetoe/webserver/assets/style.css")
	check(err)
	fmt.Fprintln(w, string(css))
}
