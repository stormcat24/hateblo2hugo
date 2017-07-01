package transformer

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type BlockquoteTransformer struct {
	doc *goquery.Document
}

func (t *BlockquoteTransformer) Transform() error {
	t.doc.Find("blockquote").Each(func(_ int, s *goquery.Selection) {
		content := s.Text()
		s.ReplaceWithHtml(fmt.Sprintf(">%s", content))
	})

	return nil
}
