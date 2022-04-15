package components

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

type App struct {
	*gtk.Box

	searchBar *SearchBar
	optionList *OptionList
}

func NewApp() *App {
	this := App{}

	this.searchBar = NewSearchBar()

	searchBar := widgets.SearchBarNew()

	scrolledWindow, _ := gtk.ScrolledWindowNew(nil, nil)

	optionList := widgets.OptionListNew()
	scrolledWindow.Add(optionList)
	scrolledWindow.SetMinContentHeight(700)

	searchBar.Connect("key_press_event", func(_ *gtk.Entry, event *gdk.Event) bool { return onKeyPress(searchBar, optionList, event) })
	searchBar.Connect("changed", func() { onQueryChanged(optionList) })

	widgets.SetFilterFunction(optionList, searchBar)

	this.Box, _ = gtk.NewBox(gtk.OrientationVertical, 0)
	this.Append(this.searchBar)
	this.Append(this.optionList)

	return &this
}

func (this *App)

var _ Component = App{}
