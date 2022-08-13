package scaffold

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.edwing123.com/ruth/pkg/resources"
)

// NewPage creates and returns an HTML page with the given title.
func NewPage(title string) (string, error) {
	html, _ := executeLayout(map[string]any{
		"Title": title,
	})

	return html, nil
}

type Project struct {
	Name        string
	Path        string
	HasPrettier bool
	HasVSCode   bool
}

func NewProject(project Project) error {
	_ = project
	projectDir := project.Path
	os.Mkdir(projectDir, 0o744)

	// Copy files from "assets" to the new project.
	err := copy(
		projectDir,
		resources.AssetsFS,
		"resources/assets", "resources/",
	)
	if err != nil {
		return err
	}

	// Copy files from "configs/editorconfig" to the new project.
	err = copy(
		projectDir,
		resources.ConfigsFS,
		"resources/configs/editorconfig",
		"resources/configs/editorconfig",
	)
	if err != nil {
		return err
	}

	if project.HasPrettier {
		// Copy files from "configs/prettier" to the new project.
		err := copy(
			projectDir,
			resources.ConfigsFS,
			"resources/configs/prettier",
			"resources/configs/prettier",
		)
		if err != nil {
			return err
		}
	}

	if project.HasVSCode {
		// Copy files from "configs/.vscode" to the new project.
		err := copy(
			projectDir,
			resources.ConfigsFS,
			"resources/configs/.vscode",
			"resources/configs",
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func copy(outDir string, src fs.ReadFileFS, dir string, pathToTrim string) error {
	err := fs.WalkDir(src, dir, func(path string, d fs.DirEntry, err error) error {
		name := strings.TrimPrefix(path, pathToTrim)
		absName := filepath.Join(outDir, name)

		isDir := d.IsDir()

		if isDir {
			os.MkdirAll(absName, 0o744)
			return nil
		}

		destFd, err := os.Create(absName)
		if err != nil {
			return err
		}
		defer destFd.Close()

		content, _ := src.ReadFile(path)
		destFd.Write(content)

		return nil
	})

	return err
}
