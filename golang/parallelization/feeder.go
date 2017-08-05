package main

import (
	"fmt"
	"sync"

	"github.com/mmcdole/gofeed"
)

var urls = []string{"http://www.yjc.ir/fa/rss/allnews", "http://www.irna.ir/fa/rss.aspx?kind=-1", "http://www.tabnak.ir/fa/rss/allnews"}

func main() {
	fmt.Println("starting ...")

	output := readTitles()

	count := 0
	for _ = range output {
		count++
	}

	fmt.Println("count is", count)
}

func readTitles() <-chan string {
	c := make(chan string)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(urls))

		for _, url := range urls {
			go func() {
				fp := gofeed.NewParser()
				feed, _ := fp.ParseURL(url)
				for _, item := range feed.Items {
					c <- item.Title
				}

				wg.Done()
			}()
		}

		wg.Wait()
		close(c)
	}()
	return c
}
