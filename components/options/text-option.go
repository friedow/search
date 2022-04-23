package options

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

type TextOption struct {
	*gtk.Box

	title  *gtk.Label
	action *gtk.Label
}

func NewTextOption(title string, action string) *TextOption {
	this := TextOption{}

	this.title = gtk.NewLabel(title)
	this.action = gtk.NewLabel(action)

	this.Box = gtk.NewBox(gtk.OrientationHorizontal, 0)
	this.Append(this.title)
	this.Append(this.action)

	return &this
}
