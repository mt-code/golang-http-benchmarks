package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go StandardGetRequest(&wg)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.UnixNano() - start.UnixNano())

	start = time.Now()
	for i := 0; i < 1000; i++ {
		go SocketGetRequest()
	}
	end = time.Now()
	fmt.Println(end.UnixNano() - start.UnixNano())
}

func StandardGetRequest(wg *sync.WaitGroup) int {

	resp, err := http.Get("http://localhost/")

	if err != nil {
		log.Fatal(err)
	}

	wg.Done()

	return resp.StatusCode
}

func SocketGetRequest() string {
	conn, err :=
	 	net.Dial("tcp", "localhost:80")

	if err != nil {
		log.Fatal(err)
	}

	tmp := make([]byte, 12)

	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	conn.Read(tmp)
	conn.Close()

	return string(tmp[9:12])
}