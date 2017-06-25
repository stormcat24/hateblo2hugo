package transformer

import (
	"github.com/PuerkitoBio/goquery"
)

type ParagraphTransformer struct {
	doc *goquery.Document
}

func (t *ParagraphTransformer) Transform() error {
	t.doc.Find("p").Each(func(_ int, s *goquery.Selection) {
		content, _ := s.Html()
		s.ReplaceWithHtml(content)
	})
	return nil
}
