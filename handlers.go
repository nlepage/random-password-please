package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type indexParams struct {
	Password, Counter, Host string
}

var index *template.Template

func registerHandlers() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/password.txt", apiHandler)
	http.HandleFunc("/counter", counterHandler)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	params := indexParams{
		Password: getPassword()[:minPasswordLength],
		Counter:  fmt.Sprint(counter),
		Host:     req.Host,
	}
	w.Header().Set("Cache-Control", "no-cache")
	index.Execute(w, params)
}

func apiHandler(w http.ResponseWriter, req *http.Request) {
	n := minPasswordLength
	n, err := strconv.Atoi(req.FormValue("len"))
	if err != nil {
		n = minPasswordLength
	} else if n < minPasswordLength {
		n = minPasswordLength
	} else if n > maxPasswordLength {
		n = maxPasswordLength
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Length", strconv.Itoa(n))
	fmt.Fprint(w, getPassword()[:n])
}

func counterHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache")
	s := strconv.FormatUint(counter, 10)
	w.Header().Set("Content-Length", strconv.Itoa(len(s)))
	fmt.Fprint(w, s)
}
