package transformer

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type TypefaceTransformer struct {
	doc *goquery.Document
}

func (t *TypefaceTransformer) Transform() error {

	t.doc.Find("b").Each(func(_ int, s *goquery.Selection) {
		content, _ := s.Html()
		s.ReplaceWithHtml(fmt.Sprintf("**%s**", content))
	})

	t.doc.Find("i").Each(func(_ int, s *goquery.Selection) {
		content, _ := s.Html()
		s.ReplaceWithHtml(fmt.Sprintf("_%s_", content))
	})

	return nil
}
