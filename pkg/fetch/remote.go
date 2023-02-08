package fetch

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func GetContentOfRemoteFile(Url string) string {
	fmt.Printf("\033[34m[i] Requesting Url: %s\033[0m\n", Url)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	Request, Err1 := http.NewRequest("GET", Url, nil)

	if Err1 != nil {
		fmt.Println(Err1)
		os.Exit(1)
	}

	Request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36")
	Request.Header.Set("Referer", Url)
	Request.Header.Set("Origin", Url)
	Request.Header.Set("Accept-Language", "en-US,en;q=0.9")
	Request.Header.Set("Accept", "*/*")

	Response, Err2 := client.Do(Request)

	if Err2 != nil {
		fmt.Println(Err2)
		os.Exit(1)
	}

	defer Response.Body.Close()

	HttpBody, Err3 := io.ReadAll(Response.Body)

	if Err3 != nil {
		fmt.Println(Err2)
		os.Exit(1)
	}

	fmt.Printf("\033[34m[i] Finished requesting successfully with status %d\033[0m\n", Response.StatusCode)

	return string(HttpBody)
}
