package transformer

import (
	"fmt"

	"bytes"

	"github.com/PuerkitoBio/goquery"
)

type HeadingTransformer struct {
	doc *goquery.Document
}

func (t *HeadingTransformer) Transform() error {

	for i := 1; i <= 5; i++ {
		t.doc.Find(fmt.Sprintf("h%d", i)).Each(func(_ int, s *goquery.Selection) {
			var buf = bytes.NewBuffer(make([]byte, 0, 100))
			for ii := 1; ii <= i; ii++ {
				buf.WriteString("#")
			}
			s.ReplaceWithHtml(fmt.Sprintf("%s %s", buf.String(), s.Text()))
		})
	}
	return nil
}
