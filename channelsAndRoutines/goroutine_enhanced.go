package channelsAndRoutines

import (
	"fmt"
	"net/http"
	"sync"
)

func checkTheLinksEnhanced(links []string) int {
	result := 0
	var wg sync.WaitGroup
	for _, link := range links {
		wg.Add(1)
		go checkLinkEnhanced(link, &result, &wg)
	}
	wg.Wait()
	return result
}

func checkLinkEnhanced(link string, addition *int, wg *sync.WaitGroup) {
	defer (*wg).Done()
	_, error := http.Get(link)
	if error == nil {
		*addition = *addition + 1
	} else {
		fmt.Println(link, "might be down!", "Error:", error)
	}
}
