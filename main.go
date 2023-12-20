package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/webcrawler/link"
	"github.com/webcrawler/utils"
	"golang.org/x/net/html"
)

func extractChildLinks(n *html.Node, uris *link.Link) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && len(a.Val) > 1 {
				t := a.Val
				uris.ChildLink = append(uris.ChildLink, t)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractChildLinks(c, uris)
	}
	return uris.ChildLink
}

func extractParent(root string, link *link.Link) []string {
	var childs []string
	var response *http.Response
	var err error
	if strings.Contains(link.Root, "https") {
		response, err = http.Get(link.Root)

	} else {
		response, err = http.Get(root + link.Root)
	}

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		doc, err := html.Parse(strings.NewReader(string(data)))
		if err != nil {
			fmt.Println("error while parsing response::", err)
		}
		childs = extractChildLinks(doc, link)
	}
	//fmt.Println("before return ************ childs-------", childs)
	return childs
}

func main() {
	//root := "https://www.monzo.com"
	root := "https://monzo.com/help/"
	var wg sync.WaitGroup
	l := link.NewLink(root)
	ll := extractParent("", l)
	for _, v := range ll {
		go func(v string) {
			wg.Add(1)
			if !utils.ToIgnore(v) {
				li := link.NewLink(v)
				tempL := extractParent(root, li)
				fmt.Println(" ************ links-------", tempL)
			} else {
				fmt.Println("  ************ bypass links-------", v)
			}
			wg.Done()
		}(v)
		wg.Wait()

	}
}
