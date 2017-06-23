package transformer

import "github.com/PuerkitoBio/goquery"

type HatenaKeywordTransformer struct {
	doc *goquery.Document
}

func (t *HatenaKeywordTransformer) Transform() error {

	t.doc.Find("a.keyword").Each(func(_ int, s *goquery.Selection) {
		s.ReplaceWithHtml(s.Text())
	})
	return nil
}
