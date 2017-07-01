package transformer

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type SyntaxTransformer struct {
	doc *goquery.Document
}

func (t *SyntaxTransformer) Transform() error {

	t.doc.Find("pre.code > span").Each(func(_ int, s *goquery.Selection) {
		s.ReplaceWithHtml(s.Text())
	})

	t.doc.Find("pre.code").Each(func(_ int, s *goquery.Selection) {
		lang, _ := s.Attr("data-lang")
		content, _ := s.Html()
		s.ReplaceWithHtml(fmt.Sprintf("```%s\n%s\n```", lang, content))
	})
	return nil
}
