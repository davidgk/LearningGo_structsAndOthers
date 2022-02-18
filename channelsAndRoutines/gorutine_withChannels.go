package channelsAndRoutines

import (
	"fmt"
	"net/http"
)

func checkTheLinksUsingChannels(links []string) int {
	result := 0
	c := make(chan int)
	value := 0
	for _, link := range links {
		go checkLinkSteven(link, &result, c)
		value = value + <-c
	}
	return result
}

func checkLinkSteven(link string, addition *int, c chan int) {
	_, error := http.Get(link)
	if error == nil {
		*addition = *addition + 1
		c <- 1
	} else {
		fmt.Println(link, "might be down!", "Error:", error)
		c <- 0
	}
}
