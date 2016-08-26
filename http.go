package main

import (
    "io"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "abcdkksdkasda"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func to_roman(n int)  string {
    if n == 2 {
        return "II"
    } else if n == 3 {
   	return "III"
    } else if n == 4 {
	return "IV"
    } else if n == 5 {
	return "V"
    } else if n == 6 {
	return "VI"
    } else if n == 7 {
	return "VII"
    } else if n == 10 {
	reiturn "X"
    }
    return "I"
}

type romanGenerator int
func (n romanGenerator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ascii_num := r.URL.Path[7:]
    i, err := strconv.Atoi(ascii_num)
    if err != nil {
        log.Print(err)
    }
    fmt.Fprintf(w, "Here's your number: %s\n", to_roman(i))
}



func main() {
    h := http.NewServeMux()

    h.Handle("/roman/", romanGenerator(1))
    h.HandleFunc("/", hello)

    err := http.ListenAndServe(":8000", h)
    log.Fatal(err)
}
