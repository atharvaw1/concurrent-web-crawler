package crawler

import (
	"golang.org/x/net/html"
	"net/url"
)

func ExtractLinks(n *html.Node) []string {
	var links []string

	var extract func(*html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link, err := url.Parse(attr.Val)
					if err == nil {
						links = append(links, link.String())
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}

	extract(n)

	return links
}

