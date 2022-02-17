package my_http

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var MyFakeWrite string

type WithMyWriter struct{}

func (WithMyWriter) Write(p []byte) (int, error) {
	MyFakeWrite = "I was Wrote"
	// Some logic to apply to my writer!
	return len(p), nil
}

func CallGoogleAndUsingReaderInterfaceUsingMyWriter() string {
	resp := callGoogle()
	myWriter := WithMyWriter{}
	io.Copy(
		myWriter,
		resp.Body,
	)
	return MyFakeWrite
}
func CallGoogleAndUsingReaderInterface() []byte {
	resp := callGoogle()
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	return bs
}

func callGoogle() *http.Response {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	return resp
}
