package channelsAndRoutines

import (
	"fmt"
	"net/http"
)

func checkTheLinks(links []string) int {
	result := 0
	for _, link := range links {
		checkLink(link, &result)
	}

	return result
}

func checkLink(link string, addition *int) {
	_, error := http.Get(link)
	if error == nil {
		*addition = *addition + 1
	} else {
		fmt.Println(link, "might be down!", "Error:", error)
	}
}
