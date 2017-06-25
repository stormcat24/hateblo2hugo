package transformer

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

var (
	speakerdeckURLRegex = regexp.MustCompile(`^\/\/speakerdeck.com\/player\/(\w+)$`)
)

type SpeakerdeckTransformer struct {
	doc *goquery.Document
}

func (t *SpeakerdeckTransformer) Transform() error {
	t.doc.Find("iframe").Each(func(_ int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if ok {
			tokens := speakerdeckURLRegex.FindStringSubmatch(src)
			if len(tokens) == 2 {
				slideID := tokens[1]
				tn := html.Node{Type: html.TextNode, Data: fmt.Sprintf(`{{< speakerdeck %s >}}`, slideID)}
				s.Before("").AfterNodes(&tn)
				s.Remove()
			}
		}
	})

	return nil
}
