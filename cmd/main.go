package main

import (
	"fmt"
	"net/http"

	"github.com/opusdv/myapi/internal"
)

func main() {
	fmt.Println("Start server.........")
	resp := internal.DriveHendler{
		Message: "test",
	}

	http.Handle("/", resp)

	http.ListenAndServe(":8080", nil)
}
