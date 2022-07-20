package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// one way
	bs := make([]byte, 32*1024)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	// another way
	io.Copy(os.Stdout, resp.Body)

	// lw implements the Writer interface by implementing Write().
	// therefore, it's able to be passed into io.Copy()
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
