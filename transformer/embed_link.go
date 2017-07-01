package transformer

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

var (
	githubRepoURLRegex = regexp.MustCompile(`^https:\/\/github\.com\/(.+)\/(.+)$`)
)

type EmbedLinkTransformer struct {
	doc *goquery.Document
}

func (t *EmbedLinkTransformer) Transform() error {

	t.doc.Find("iframe.embed-card").Each(func(_ int, s *goquery.Selection) {
		title, _ := s.Attr("title")
		s.Next().Find("cite.hatena-citation > a").Each(func(_ int, ss *goquery.Selection) {
			href, _ := ss.Attr("href")

			tokens := githubRepoURLRegex.FindStringSubmatch(href)
			if len(tokens) == 3 {
				html := fmt.Sprintf(`
<div class="github-card" data-user="%s" data-repo="%s" data-width="400" data-height="" data-theme="default"></div>
<script src="https://cdn.jsdelivr.net/github-cards/latest/widget.js"></script>
`, tokens[1], tokens[2])
				s.ReplaceWithHtml(html)
				fmt.Println(html)
			} else {
				s.ReplaceWithHtml(fmt.Sprintf(`[%s](%s)`, title, href))
			}
			ss.Parent().Remove()
		})
	})

	return nil
}
