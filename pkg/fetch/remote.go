package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetContentOfRemoteFile(Url string) string {
	fmt.Printf("\033[34m[i] Requesting Url: %s\033[0m\n", Url)

	Response, Err1 := http.Get(Url)

	if Err1 != nil {
		fmt.Println(Err1)
		os.Exit(1)
	}

	defer Response.Body.Close()

	HttpBody, Err2 := io.ReadAll(Response.Body)
	if Err2 != nil {
		fmt.Println(Err2)
		os.Exit(1)
	}

	return string(HttpBody)
}
