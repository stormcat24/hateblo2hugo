package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stormcat24/hateblo2hugo/service"
	"github.com/stormcat24/hateblo2hugo/transformer"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate hatenablog to hugo",
	RunE: func(cmd *cobra.Command, args []string) error {

		inputPath, err := cmd.Flags().GetString("input-path")
		if err != nil {
			return err
		}

		outputPath, err := cmd.Flags().GetString("output-path")
		if err != nil {
			return err
		}

		inputTarget, err := resolvePath(inputPath)
		if err != nil {
			return err
		}

		statIn, err := os.Stat(inputTarget)
		if err != nil {
			return err
		}

		if statIn.IsDir() {
			return errors.Wrapf(err, "%s is not file", inputTarget)
		}

		outputTarget, err := resolvePath(outputPath)
		if err != nil {
			return err
		}

		statOut, err := os.Stat(outputTarget)
		if err != nil {
			return err
		}

		if !statOut.IsDir() {
			return errors.Wrapf(err, "%s is not directory", outputTarget)
		}

		mts := service.NewMovableType()
		entries, err := mts.Parse(inputTarget)
		if err != nil {
			return err
		}

		for _, entry := range entries {

			sr := strings.NewReader(entry.Body)
			doc, err := goquery.NewDocumentFromReader(sr)
			if err != nil {
				return err
			}

			tf := transformer.NewTransformer(doc)
			if err := tf.Transform(); err != nil {
				return err
			}

			newHtml, err := doc.Find("body").Html()
			if err != nil {
				return err
			}
			entry.Body = newHtml

			ts := service.NewMigration(entry, outputTarget)
			if err := ts.Execute(); err != nil {
				return err
			}
		}

		return nil
	},
}

func resolvePath(path string) (string, error) {
	var target string
	if filepath.IsAbs(path) {
		target = path
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		target = filepath.Join(pwd, path)
	}
	return target, nil
}

func initMigrateCmd() {
	migrateCmd.PersistentFlags().StringP("input-path", "i", "", "input movable type data file")
	migrateCmd.PersistentFlags().StringP("output-path", "o", "", "output hugo data directory")
}
