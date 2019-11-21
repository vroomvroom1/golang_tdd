package concurrency

import "time"

// WebsiteChecker type is a function that takes a string and returns a bool
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites accepts a websitechecker (function) and a slice of urls to check
// then returns a map or urls to the result of true or false
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// go routine to run process with anon func
		// the trailing () executes the function, also all variables are available
		go func(u string) {
			// u is a copy of the url passed in
			// <- send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// <- receive expresssion to assign value to var
		result := <-resultChannel
		results[result.string] = result.bool
	}

	time.Sleep(2 * time.Second)

	return results
}
