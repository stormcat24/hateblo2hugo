package transformer

import "github.com/PuerkitoBio/goquery"

type EmptyParagraphTransformer struct {
	doc *goquery.Document
}

func (t *EmptyParagraphTransformer) Transform() error {
	t.doc.Find("p").Each(func(_ int, s *goquery.Selection) {
		content, _ := s.Html()
		if content == "" {
			s.Remove()
		}
	})
	return nil
}
