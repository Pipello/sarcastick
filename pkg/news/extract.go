package news

import (
	"slices"
	"strings"

	"golang.org/x/net/html"
)

func FindNewsLink(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && strings.Contains(a.Val, "czech-news-in-brief") {
				return a.Val
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if l := FindNewsLink(c); l != "" {
			return l
		}
	}
	return ""
}

func ExtractContentWithTitle(n *html.Node) []News {
	if n.Type == html.ElementNode && n.Data == "div" && slices.ContainsFunc(n.Attr, containContentClass) {
		ct := []News{}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.Data == "div" && slices.ContainsFunc(c.Attr, containTitleClass) {
				ct = append(ct, News{
					Title:   extractText(c),
					Content: extractContent(c.NextSibling),
				})
			}
		}
		if len(ct) > 0 {
			return ct
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if ct := ExtractContentWithTitle(c); len(ct) > 0 {
			return ct
		}
	}
	return nil
}

func containTextWrapperClass(a html.Attribute) bool {
	return a.Key == "class" && strings.Contains(a.Val, "widget text")
}

func containTitleClass(a html.Attribute) bool {
	return a.Key == "class" && strings.Contains(a.Val, "headinglevel2")
}

func containContentClass(a html.Attribute) bool {
	return a.Key == "class" && strings.Contains(a.Val, "content")
}

func extractContent(n *html.Node) string {
	r := ""
	for c := n; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "div" && slices.ContainsFunc(c.Attr, containTextWrapperClass) {
			r += extractText(c.FirstChild)
		}
		if c.Type == html.ElementNode && c.Data == "div" && slices.ContainsFunc(c.Attr, containTitleClass) {
			break
		}
	}
	return r
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	r := ""
	if n.Type == html.ElementNode {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			r += extractText(c)
		}
	}
	return r
}
