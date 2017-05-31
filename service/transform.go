package service

import (
	"fmt"
	"path/filepath"

	"github.com/catatsuy/movabletype"
	"github.com/pkg/errors"
	"github.com/stormcat24/hateblo2hugo/helper"
	"github.com/stormcat24/hateblo2hugo/hugo"
)

type Transform interface {
	Execute() error
	OutputFilePath() string
}

type TransformImpl struct {
	entry         *movabletype.Entry
	outputDirRoot string
}

func NewTransform(entry *movabletype.Entry, outputDirRoot string) Transform {
	return &TransformImpl{
		entry:         entry,
		outputDirRoot: outputDirRoot,
	}
}

func (s *TransformImpl) Execute() error {
	outpath := filepath.Join(s.outputDirRoot, s.OutputFilePath())

	page := hugo.CreateHugoPage(s.entry)
	content, err := page.Render()
	if err != nil {
		return errors.Wrapf(err, "render hugo markdown is failed. [%s]", s.entry.Basename)
	}

	if err := helper.WriteFileWithDirectory(outpath, content, 0644); err != nil {
		return errors.Wrapf(err, "failed to write data file. path=%s", outpath)
	}

	return nil
}

func (s *TransformImpl) OutputFilePath() string {
	return fmt.Sprintf("%s.md", s.entry.Basename)
}
