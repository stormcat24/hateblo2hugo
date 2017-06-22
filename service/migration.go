package service

import (
	"fmt"
	"path/filepath"

	"github.com/catatsuy/movabletype"
	"github.com/pkg/errors"
	"github.com/stormcat24/hateblo2hugo/helper"
	"github.com/stormcat24/hateblo2hugo/hugo"
)

type Migration interface {
	Execute() error
	OutputFilePath() string
}

type MigrationImpl struct {
	entry         *movabletype.Entry
	outputDirRoot string
}

func NewMigration(entry *movabletype.Entry, outputDirRoot string) Migration {
	return &MigrationImpl{
		entry:         entry,
		outputDirRoot: outputDirRoot,
	}
}

func (s *MigrationImpl) Execute() error {
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

func (s *MigrationImpl) OutputFilePath() string {
	return fmt.Sprintf("%s.md", s.entry.Basename)
}
