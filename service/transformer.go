package service

import (
	"reflect"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
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

func NewChainTransformer(doc *goquery.Document) Transformer {

	return &ChainTransformer{
		transformers: []Transformer{
			&HatenaKeywordRemoveTransformer{
				doc: doc,
			},
			&HatenaPhotolifeTransformer{
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
