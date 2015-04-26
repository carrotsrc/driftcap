package main
//import "time"
import "fmt"
import "strings"

func branch(resource *SiteResource, domain string, c chan int) {
	var url string

	if resource.ResourceLocation == Local {
		if strings.HasPrefix(resource.Ref, "/") {
			url = domain+resource.Ref
		} else {
			url = resource.Ref
		}
	} else {
		if !strings.HasPrefix(resource.Ref, "http")  {
			url = "http://"+resource.Ref
		} else {
			url = resource.Ref
		}
	}

	response, delta := performRequest(url)
	resource.Delta = delta
	if response == nil {
		fmt.Println(url,"is unreachable")
		c <- 0
		return
	}
	defer response.Body.Close()
	c <- 0
}

func runner(url string, c chan int) {
	response, delta := performRequest(url)
	
	defer response.Body.Close()
	resources := parseLinks(response.Body, url)

	total := len(resources)
	assets := 0

	fb := make(chan int)

	for index := 0; index < total; index++ {
		if resources[index].ResourceType == Asset {
			assets++
			go branch(resources[index], url, fb)
		}
	}

	branches := 0
	for branches < assets {
		<-fb
		branches++
	}

	for index := 0; index < total; index++ {
		if resources[index].ResourceType == Asset {
			delta += resources[index].Delta
		}
	}
	fmt.Println("Page and ",assets,"assets -- delta:", delta)
	c <- 0
}

func main() {
	c := make(chan int)
	url := "http://localhost"

	fmt.Println("Performing run...")
	go runner(url, c)

	users := 1
	for users > 0 {
		<-c
		users--
	}
}
