package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
