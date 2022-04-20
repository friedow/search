package plugins

import (
	"friedow/tucan-search/components"
	"os/exec"
)

type GitRepository struct {
	*components.TextOption

	title string
	path  string
}

func NewGitRepository(title string, path string) *GitRepository {
	this := GitRepository{}

	this.TextOption = components.NewTextOption(title, "Enter to open VSCode")

	this.title = title
	this.path = path

	return &this
}

func (this *GitRepository) OnActivate() {
	exec.Command("code", this.path).Output()
}
