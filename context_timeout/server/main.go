package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// server端,随机出现慢反应
func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10)
		_, _ = fmt.Fprintf(w, "slow response")
		return
	}
	_, _ = fmt.Fprintf(w, "quick response")

}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
