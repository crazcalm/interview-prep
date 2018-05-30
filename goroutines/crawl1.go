/*
go run crawl1.go https://en.wikipedia.org/wiki/Binary_search_algorithm


The initial error message is a surprising report of a DNS lookup failure for a reliable domain. The
subsequent error message (socket: too many open files) reveals the cause: the program created so many
network connections at once that it exceeded the per-process limit on the number of open files, causing
operations such as DNS lookup and calls to net.Dial to start failing.

Summary: The program is too parallel.
*/
package main

import (
	"fmt"
	"log"
	"os"

	//ch5/links
	"golang.org/x/net/html"
	"net/http"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := extract(url)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	//Start with the command-line arguments
	go func() { worklist <- os.Args[1:] }()

	//Start with the command-line arguments.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

//gopl.io/ch5/links
func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue //ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil) //Where is this function defined?!
	return links, nil
}

//gopl.io/ch5/outline2
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
