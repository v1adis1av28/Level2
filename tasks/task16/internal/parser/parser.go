package parser

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func ExtractLinksAndAssets(htmlContent, baseURL string) ([]string, []string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, nil, err
	}

	var links, assets []string
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, nil, err
	}

	var extract func(*html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "a":
				if href := getAttr(n, "href"); href != "" {
					if absURL, err := resolveURL(base, href); err == nil && isHTTP(absURL) {
						links = append(links, absURL)
					}
				}
			case "link":
				if rel := getAttr(n, "rel"); rel == "stylesheet" || rel == "icon" {
					if href := getAttr(n, "href"); href != "" {
						if absURL, err := resolveURL(base, href); err == nil {
							assets = append(assets, absURL)
						}
					}
				}
			case "script":
				if src := getAttr(n, "src"); src != "" {
					if absURL, err := resolveURL(base, src); err == nil {
						assets = append(assets, absURL)
					}
				}
			case "img":
				if src := getAttr(n, "src"); src != "" {
					if absURL, err := resolveURL(base, src); err == nil {
						assets = append(assets, absURL)
					}
				}
				if srcset := getAttr(n, "srcset"); srcset != "" {
					urls := parseSrcSet(srcset)
					for _, u := range urls {
						if absURL, err := resolveURL(base, u); err == nil {
							assets = append(assets, absURL)
						}
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}

	extract(doc)
	return removeDuplicates(links), removeDuplicates(assets), nil
}

func getAttr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

func resolveURL(base *url.URL, ref string) (string, error) {
	u, err := base.Parse(ref)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func isHTTP(u string) bool {
	return strings.HasPrefix(u, "http://") || strings.HasPrefix(u, "https://")
}

func parseSrcSet(srcset string) []string {
	var urls []string
	parts := strings.Split(srcset, ",")
	for _, part := range parts {
		fields := strings.Fields(part)
		if len(fields) > 0 {
			urlPart := strings.TrimSpace(fields[0])
			if urlPart != "" {
				urls = append(urls, urlPart)
			}
		}
	}
	return urls
}

func removeDuplicates(items []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, item := range items {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
