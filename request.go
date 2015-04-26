package main
import "fmt"
import "time"
import "net/http"

func performRequest(url string) (response *http.Response, delta time.Duration) {

	startTime := time.Now()
	resp, err := http.Get(url)
	endTime := time.Now()
	delta = endTime.Sub(startTime)

	if err != nil {
		fmt.Printf("Error on request\n")
		return nil, delta
	}

	return resp, delta
}
