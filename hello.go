package main

import (
	"fmt"
	"net/http"

	"github.com/huntermentees2018/gae-hello-world/src/scheduler"
	"google.golang.org/appengine"
)

func main() {
	scheduler.Schedule()
	// fmt.Println(src.RegexCommand("hello"))
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, hunter mentees!")
}
