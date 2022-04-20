package plugins

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

type PluginOption interface {
	gtk.Widgetter

	OnActivate()
}

type GenericPlugin[T PluginOption] interface {
	SetHeader(current *gtk.ListBoxRow, before *gtk.ListBoxRow)
	onActivate(row *gtk.ListBoxRow)
	Name() string
	NewOptions() []T
}

type Plugin[T PluginOption] struct {
	*gtk.ListBox
	GenericPlugin[T]

	options []T
}

func NewPlugin[T PluginOption]() *Plugin[T] {
	this := Plugin[T]{}

	this.ListBox = gtk.NewListBox()
	this.SetHeaderFunc(this.setHeader)
	this.ConnectRowActivated(this.onActivate)

	this.options = this.NewOptions()
	for _, option := range this.options {
		this.Append(option)
	}

	return &this
}

func (this Plugin[T]) setHeader(current *gtk.ListBoxRow, before *gtk.ListBoxRow) {
	if current.Index() == 0 && current.Header() == nil {
		header := gtk.NewLabel(this.Name())
		current.SetHeader(header)
	} else if current.Header() != nil {
		current.SetHeader(nil)
	}
}

func (this Plugin[T]) onActivate(row *gtk.ListBoxRow) {
	gitRepository := row.Child().(T)
	gitRepository.OnActivate()
}
