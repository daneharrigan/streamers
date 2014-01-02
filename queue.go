package main

import (
	"net/http"
	"fmt"
	"time"
)

var c int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "text/plain")
		f, _ := w.(http.Flusher)
		n := r.FormValue("n")

		for {
			c++
			_, err := fmt.Fprintf(w, "%d\r", c)
			if err != nil {
				fmt.Printf("fn=Fprintf n=%s counter=%d error=%q\n", n, c, err)
				c--
				return
			}

			f.Flush()
			fmt.Printf("fn=Flush n=%s counter=%d\n", n, c)
			time.Sleep(time.Second)
		}
	})
	http.ListenAndServe(":5000", nil)
}
