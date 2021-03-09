package main

import (
	"fmt"
	"github.com/gregfolker/auto-project-builder/pkg/project"
	"github.com/gregfolker/auto-project-builder/pkg/colors"
	"github.com/gregfolker/auto-project-builder/pkg/git"
	"github.com/gregfolker/auto-project-builder/internal/ui"
)

const (
	appName = "Auto Project Builder"
)

func main() {
	fmt.Printf("%s\n\n", appName)

	prog := project.NewProject()

	if err := ui.GetUserInput(prog); err != nil {
		fmt.Printf(colors.ColorText("Error: ", colors.ANSI_RED) + "%v\n\n", err)
		return
	}

	fmt.Printf("\nGenerating project...\n\n")

	if err := prog.CreateNewProjectDir(); err != nil {
		fmt.Printf(colors.ColorText("Error: ", colors.ANSI_RED) + "%v\n\n", err)
		return
	}

	if err := prog.CreateProjectFile("README"); err != nil {
		fmt.Printf(colors.ColorText("Error: ", colors.ANSI_RED) + "%v\n\n", err)
		return
	}

	if err := prog.CreateProjectFile("TODO"); err != nil {
		fmt.Printf(colors.ColorText("Error: ", colors.ANSI_RED) + "%v\n\n", err)
		return
	}

	if err := prog.CreateProjectFile("main"); err != nil {
		fmt.Printf(colors.ColorText("Error: ", colors.ANSI_RED) + "%v\n\n", err)
		return
	}

	if err := git.InitializeNewRepo(prog.Path); err != nil {
		fmt.Printf(colors.ColorText("Error: ", colors.ANSI_RED) + "%v\n\n", err)
		return
	}

	prog.PrintProjectInfo()
	fmt.Printf("\nDone!\n")
}
