package project

import (
	"os"
	"path"
	"strings"
	"fmt"
	"errors"
	"github.com/gregfolker/auto-project-builder/pkg/files"
	"github.com/gregfolker/auto-project-builder/pkg/languages"
)

type Project struct {
	Name string
	Author string
	Language string
	Path string
}

func NewProject() *Project {
	p := &Project{}

	return p
}

func (prog *Project) CreateNewProjectDir() error {
	var dir string
	userHome, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	// Avoid creating directories with spaces
	if strings.Contains(prog.Name, " ") {
		dir = strings.ReplaceAll(prog.Name, " ", "-")
	} else {
		dir = strings.ToLower(prog.Name)
	}

	prog.Path = path.Join(userHome, strings.ToLower(prog.Language), "src", dir)

	if err := os.MkdirAll(prog.Path, os.FileMode(0755)); err != nil {
		return err
	}

	fmt.Printf("Created new project directory %s\n", prog.Path)

	return nil
}

func (prog *Project) CreateProjectFile(filename string) error {
	if prog.Path == "" {
		return errors.New("Project directory does not exist\n")
	}

	f := path.Join(prog.Path, filename)

	switch filename {
	case "README":
		return files.GenerateReadMe(f + languages.MARKDOWN_EXT, prog.Name, prog.Author)
	case "main":
		return files.GenerateMain(f, prog.Name, prog.Author, prog.Language)
	default:
		return errors.New("Unknown file " + filename + ", unable to create\n")
	}

	return nil
}
