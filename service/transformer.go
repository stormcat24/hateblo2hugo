package service

import (
	"fmt"
	"reflect"

	"golang.org/x/net/html"

	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

var (
	tweetURLRegex = regexp.MustCompile(`^https:\/\/twitter\.com\/[\w-]+\/status\/(\d+)$`)
)

type Transformer interface {
	Transform() error
}

type ChainTransformer struct {
	transformers []Transformer
	doc          *goquery.Document
}

func (t *ChainTransformer) Transform() error {

	for _, tf := range t.transformers {
		if err := tf.Transform(); err != nil {
			return errors.Wrapf(err, "transform entry is failed. at %s", reflect.TypeOf(tf))
		}
	}
	return nil
}

func NewTransformer(doc *goquery.Document) Transformer {

	return &ChainTransformer{
		transformers: []Transformer{
			&EmptyParagraphTransformer{
				doc: doc,
			},
			&HatenaKeywordRemoveTransformer{
				doc: doc,
			},
			&HatenaPhotolifeTransformer{
				doc: doc,
			},
			&TweetTransformer{
				doc: doc,
			},
		},
	}
}

type HatenaKeywordRemoveTransformer struct {
	doc *goquery.Document
}

func (t *HatenaKeywordRemoveTransformer) Transform() error {

	t.doc.Find("a.keyword").Each(func(_ int, s *goquery.Selection) {
		s.ReplaceWithHtml(s.Text())
	})
	return nil
}

type HatenaPhotolifeTransformer struct {
	doc *goquery.Document
}

func (t *HatenaPhotolifeTransformer) Transform() error {
	return nil
}

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

type TweetTransformer struct {
	doc *goquery.Document
}

func (t *TweetTransformer) Transform() error {
	t.doc.Find("blockquote.twitter-tweet").Each(func(_ int, s *goquery.Selection) {
		s.Find("a").Each(func(_ int, ss *goquery.Selection) {
			href, ok := ss.Attr("href")
			if ok {
				tokens := tweetURLRegex.FindStringSubmatch(href)
				if len(tokens) == 2 {
					tweetID := tokens[1]
					tn := html.Node{Type: html.TextNode, Data: fmt.Sprintf(`{{< tweet %s >}}`, tweetID)}
					s.Before("").AfterNodes(&tn)
					s.Remove()
				}
			}
		})
	})

	t.doc.Find("script[src='//platform.twitter.com/widgets.js']").Remove()
	return nil
}
