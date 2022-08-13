package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.edwing123.com/ruth/pkg/scaffold"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	titleCaser = cases.Title(language.Tag{})
)

var addPageCmd = &cobra.Command{
	Use:   "add [page-name]",
	Short: "Add new pages to the project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()

		for _, pageName := range args {
			pagePath := filepath.Join(cwd, pageName)
			err := os.Mkdir(pagePath, 0o744)
			if err != nil {
				log.Fatalln(err)
			}

			title := titleCaser.String(strings.Replace(pageName, "-", " ", -1))
			pageHTML, err := scaffold.NewPage(title)
			if err != nil {
				log.Fatalln(err)
			}

			pageFile := filepath.Join(pagePath, "index.html")
			err = ioutil.WriteFile(pageFile, []byte(pageHTML), 0o744)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addPageCmd)
}
