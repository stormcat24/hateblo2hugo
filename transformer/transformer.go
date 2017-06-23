package transformer

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

func NewTransformer(doc *goquery.Document) Transformer {

	return &ChainTransformer{
		transformers: []Transformer{
			&EmptyParagraphTransformer{doc},
			&HatenaKeywordTransformer{doc},
			&HatenaPhotolifeTransformer{doc},
			&TweetTransformer{doc},
			&SpeakerdeckTransformer{doc},
		},
	}
}
