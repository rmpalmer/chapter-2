package search

import (
	"log"
	"sync"
)

var matchers = make(map[string],Matcher)

func Run(searchTerm string) {

	// retrieve the list of feeds to search through
	feeds, err := RetrieveFeeds()
	if (err != nil) {
		log.Fatal(err)
	}

	// create an unbuffered channel to receive match results
	results := make(chan *Result)

	// set up a wait group so we can process all the feeds
	var waitGroup sync.WaitGroup

	// set the number of goroutines we need to wait for
	waitGroup.Add(len(feeds))

	for _, feed := range(feeds) {
		//retrieve a marchter for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// launch the goroutine to perform the search
		go func (matcher Matcher, feed *Feed) {
			Match (matcher, feet, searchTerm, results)
			waitGroup.Done()
		} (matcher, feed)
	}

	// launch a goroutine to monitor when all the work is done
	go func () {
		//wait for everything to be processed
		waitGroup.Wait()

		//close the channel to signal to the display function that 
		// can exit the program
		close(results)
	}()

	// start displaying results as they are available and return
	// after the final result is displayed
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _,exists := matchers[feedType], exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}