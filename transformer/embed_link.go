package transformer

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type EmbedLinkTransformer struct {
	doc *goquery.Document
}

func (t *EmbedLinkTransformer) Transform() error {

	t.doc.Find("iframe.embed-card").Each(func(_ int, s *goquery.Selection) {
		title, _ := s.Attr("title")
		s.Next().Find("cite.hatena-citation > a").Each(func(_ int, ss *goquery.Selection) {
			href, _ := ss.Attr("href")
			s.ReplaceWithHtml(fmt.Sprintf(`[%s](%s)`, title, href))
			ss.Parent().Remove()
		})
	})

	return nil
}
