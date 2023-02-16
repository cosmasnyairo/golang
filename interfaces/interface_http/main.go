package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	respose, err := http.Get("https://en.wikipedia.org/wiki/One_Piece")
	if err != nil {
		log.Fatal("Rrror is:", err)
	}
	logwriter := logWriter{}
	io.Copy(logwriter, respose.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
