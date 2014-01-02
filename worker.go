package main

import (
	"net/http"
	"flag"
	"fmt"
	"bufio"
	"time"
)

var (
	n = flag.String("n", "0", "Worker Identifier")
	p = flag.Int("p", 0, "Sleep between reads")
)

func main() {
	flag.Parse()
	for {
		fmt.Println("fn=stream at=start")
		stream()
		fmt.Println("fn=stream at=finish")
		time.Sleep(500 * time.Millisecond)
	}
}

func stream() {
	r, err := http.Get("http://localhost:5000?n="+*n)
	if err != nil {
		fmt.Printf("fn=Get error=%q", err)
		return
	}

	s := bufio.NewReader(r.Body)
	d := time.Duration(*p)

	for {
		b, err := s.ReadBytes('\r')
		if err != nil {
			fmt.Printf("fn=ReadBytes error=%q", err)
			return
		}

		fmt.Printf("n=%s p=%d counter=%s\n", *n, *p, b[:len(b)-1])
		time.Sleep(d * time.Second)
	}
}
