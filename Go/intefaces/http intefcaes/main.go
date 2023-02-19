package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// fmt.Println(resp)

	// bs := make([]byte, 9999)
	// // fmt.Println(bs)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// fmt.Println(*os.Stdout)
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("BS Length:", len(bs))
	return len(bs), nil
}
