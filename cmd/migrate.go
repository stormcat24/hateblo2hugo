package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"path/filepath"
	"os"
	"github.com/pkg/errors"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate hatenablog to hugo",
	RunE: func(cmd *cobra.Command, args []string) error {

		inputPath, err := cmd.Flags().GetString("input-path")
		if err != nil {
			return err
		}

		var target string
		if filepath.IsAbs(inputPath) {
			target = inputPath
		} else {
			pwd, err := os.Getwd()
			if err != nil {
				return err
			}
			target = filepath.Join(pwd, inputPath)
		}

		stat, err := os.Stat(target)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			return errors.Wrapf(err, "%s is not file", target)
		}
		fmt.Println(target)

		return nil
	},
}

func initMigrateCmd() {
	migrateCmd.PersistentFlags().StringP("input-path", "i", "", "input movable type data file")
}