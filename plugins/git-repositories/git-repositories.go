package plugins

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type GitRepositoriesPlugin struct {
	*gtk.ListBox

	options []*GitRepository
}

// TODO: move this to generic plugin
func NewGitRepositoriesPlugin() *GitRepositoriesPlugin {
	this := GitRepositoriesPlugin{}

	this.ListBox = gtk.NewListBox()
	this.SetHeaderFunc(this.setHeader)
	this.ConnectRowActivated(this.onActivate)

	this.options = this.newOptions()
	for _, gitRepository := range this.options {
		this.Append(gitRepository)
	}

	return &this
}

func (this GitRepositoriesPlugin) setHeader(current *gtk.ListBoxRow, before *gtk.ListBoxRow) {
	if current.Index() == 0 && current.Header() == nil {
		header := gtk.NewLabel(this.name())
		current.SetHeader(header)
	} else if current.Header() != nil {
		current.SetHeader(nil)
	}
}

func (this GitRepositoriesPlugin) onActivate(row *gtk.ListBoxRow) {
	gitRepository := row.Child().(GitRepository)
	gitRepository.OnActivate()
}

// end TODO

func (this GitRepositoriesPlugin) name() string {
	return "Git Repositories"
}

func (this GitRepositoriesPlugin) newOptions() []*GitRepository {
	home := os.Getenv("HOME")
	gitRepositories := []*GitRepository{}

	filepath.WalkDir(home,
		func(path string, info fs.DirEntry, err error) error {
			// bubble errors
			if err != nil {
				return err
			}

			// Skip hidden directories
			pathParts := strings.Split(path, "/")
			if len(pathParts) >= 2 {
				parentDirIndex := len(pathParts) - 2
				parentDir := pathParts[parentDirIndex]

				if strings.HasPrefix(parentDir, ".") {
					return fs.SkipDir
				}
			}

			// Add git directories to list
			if strings.HasSuffix(path, ".git") {
				gitRepositoryPath := strings.TrimSuffix(path, "/.git")
				gitRepositoryTitle := strings.Replace(gitRepositoryPath, home, "~", 1)
				gitRepository := NewGitRepository(gitRepositoryTitle, gitRepositoryPath)
				gitRepositories = append(gitRepositories, gitRepository)
				return nil
			}

			return nil
		},
	)

	return gitRepositories
}
