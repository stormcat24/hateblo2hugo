package transformer

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/catatsuy/movabletype"
	"github.com/pkg/errors"
)

type HatenaPhotolifeTransformer struct {
	doc             *goquery.Document
	entry           *movabletype.Entry
	outputImageRoot string
	updateImage     bool
}

var (
	regexImgStyle = regexp.MustCompile(`width:([0-9]+)px`)
)

func (t *HatenaPhotolifeTransformer) Transform() (e error) {
	t.doc.Find("span[itemtype='http://schema.org/Photograph'] > img").Each(func(_ int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		style, _ := s.Attr("style")

		if t.updateImage {
			if src != "" {
				if err := t.saveImage(src); err != nil {
					e = err
					return
				}
				log.Printf("dowloaded %s is success", src)
			}
		}

		extAttr := ""
		if style != "" {
			tokens := regexImgStyle.FindStringSubmatch(style)
			if len(tokens) > 1 {
				extAttr = fmt.Sprintf(`width="%spx"`, tokens[1])
			}
		}

		imgPath := filepath.Join("/images", t.entry.Basename, filepath.Base(src))
		s.Parent().ReplaceWithHtml(fmt.Sprintf(`{{< figure src="%s" %s >}}`, imgPath, extAttr))

		s.Remove()
	})
	return nil
}

func (t *HatenaPhotolifeTransformer) saveImage(src string) error {

	outputImageDir := fmt.Sprintf("%s/%s", t.outputImageRoot, t.entry.Basename)
	if err := os.MkdirAll(outputImageDir, 0777); err != nil {
		return errors.Wrapf(err, "create directory is failed. [%s]", outputImageDir)
	}

	res, err := http.Get(src)
	if err != nil {
		return errors.Wrapf(err, "download %s is failed", src)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Status Code %d: src=%s", res.StatusCode, src)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrapf(err, "read file %s is failed", src)
	}

	filename := filepath.Base(src)
	outputImagePath := fmt.Sprintf("%s/%s", outputImageDir, filename)
	file, err := os.OpenFile(outputImagePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.Wrapf(err, "create file %s is failed", outputImagePath)
	}

	defer func() {
		file.Close()
	}()

	if _, err := file.Write(body); err != nil {
		return errors.Wrapf(err, "write file %s is failed", outputImagePath)
	}
	return nil
}
