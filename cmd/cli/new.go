package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.edwing123.com/ruth/pkg/scaffold"
)

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new web project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		projectDir, _ := filepath.Abs(name)
		hasPrettier, _ := cmd.Flags().GetBool("prettier")
		hasVsCode, _ := cmd.Flags().GetBool("vscode")

		// Create the basic files.
		err := scaffold.NewProject(scaffold.Project{
			Name:        name,
			Path:        projectDir,
			HasPrettier: hasPrettier,
			HasVSCode:   hasVsCode,
		})
		if err != nil {
			log.Fatalln(err)
		}

		// Add the default `index.html` page.
		indexHTML, err := scaffold.NewPage("Home")
		err = ioutil.WriteFile(filepath.Join(projectDir, "index.html"), []byte(indexHTML), 0o744)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().Bool("prettier", false, "add Prettier config file")
	newCmd.Flags().Bool("vscode", false, "add VSCode config file")
}
