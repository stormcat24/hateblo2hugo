package transformer

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

var (
	tweetURLRegex = regexp.MustCompile(`^https:\/\/twitter\.com\/[\w-]+\/status\/(\d+)$`)
)

type TweetTransformer struct {
	doc *goquery.Document
}

func (t *TweetTransformer) Transform() error {
	t.doc.Find("blockquote.twitter-tweet").Each(func(_ int, s *goquery.Selection) {
		s.Find("a").Each(func(_ int, ss *goquery.Selection) {
			href, ok := ss.Attr("href")
			if ok {
				tokens := tweetURLRegex.FindStringSubmatch(href)
				if len(tokens) == 2 {
					tweetID := tokens[1]
					tn := html.Node{Type: html.TextNode, Data: fmt.Sprintf(`{{< tweet %s >}}`, tweetID)}
					s.Before("").AfterNodes(&tn)
					s.Remove()
				}
			}
		})
	})

	t.doc.Find("script[src='//platform.twitter.com/widgets.js']").Remove()
	return nil
}
