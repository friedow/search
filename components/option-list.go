package components

import (
	"strings"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type OptionList struct {
	*gtk.ScrolledWindow

	pluginsWrapper struct {
		*gtk.Box

		plugins []*gtk.ListBox
	}
}

func NewOptionList() *OptionList {
	this := OptionList{}

	this.pluginsWrapper.plugins = []Plugin{
		// 		GitRepositoriesPlugin{},
		// 		OpenWindowsPlugin{},
		// 		ApplicationsPlugin{},
	}

	this.ListBox = gtk.NewListBox()

	this.Append(gtk.NewLabel("test"))
	this.Append(gtk.NewLabel("test"))
	this.Append(gtk.NewLabel("test"))
	this.Append(gtk.NewLabel("test"))

	this.selectFirstRow()

	this.optionList = components.NewOptionList()
	this.optionList.SetFilterFunction(this.searchBar)

	this.pluginsWrapper.Box = gtk.NewBox(gtk.OrientationHorizontal, 0)

	this.ScrolledWindow = gtk.NewScrolledWindow()
	this.ScrolledWindow.SetMinContentHeight(700)
	this.ScrolledWindow.SetChild(this.pluginsWrapper)

	return &this
}

func (this *OptionList) selectFirstRow() {
	firstRow := this.RowAtIndex(0)
	if firstRow != nil {
		this.SelectRow(firstRow)
	}
}

func (this *OptionList) selectPreviousRow() {
	currentRow := this.SelectedRow()
	if currentRow == nil {
		this.selectFirstRow()
		return
	}

	nextRow := this.RowAtIndex(currentRow.Index() - 1)
	if nextRow != nil {
		this.SelectRow(nextRow)
	}
}

func (this *OptionList) selectNextRow() {
	currentRow := this.SelectedRow()
	if currentRow == nil {
		this.selectFirstRow()
		return
	}

	nextRow := this.RowAtIndex(currentRow.Index() + 1)
	if nextRow != nil {
		this.SelectRow(nextRow)
	}
}

func (this *OptionList) OnKeyPress(keyVal uint) bool {
	if keyVal == gdk.KEY_Up {
		this.selectPreviousRow()
		return true
	}

	if keyVal == gdk.KEY_Down {
		this.selectNextRow()
		return true
	}

	if keyVal == gdk.KEY_Return {
		this.SelectedRow().Activate()
		return true
	}

	this.InvalidateFilter()
	return false
}

// func OptionListNew() *gtk.ListBox {
// 	optionList, _ := gtk.ListBoxNew()
// 	optionList.SetHeaderFunc(setHeader)

// 	pluginList := plugins.Plugins()
// 	for _, plugin := range pluginList {
// 		// bind the current value of plugin to the closure
// 		// https://go.dev/doc/faq#closures_and_goroutines
// 		plugin := plugin

// 		optionModels := plugin.GetOptionModels()
// 		for _, optionModel := range optionModels {
// 			optionModel := optionModel

// 			optionWidget := OptionWidgetNew(optionModel.Title, optionModel.ActionText)
// 			setOptionModel(optionWidget, optionModel)

// 			optionWidget.Connect("key_press_event", func() { plugin.OnActivate(optionModel) })

// 			optionList.Add(optionWidget)
// 		}
// 	}

// 	SelectFirstRow(optionList)
// 	return optionList
// }

// func SelectFirstRow(optionList *gtk.ListBox) {
// 	firstRow := optionList.GetRowAtIndex(0)
// 	if firstRow != nil {
// 		optionList.SelectRow(firstRow)
// 	}
// }

// func setHeader(currentRow *gtk.ListBoxRow, previousRow *gtk.ListBoxRow) {
// 	currentHeader, _ := currentRow.GetHeader()

// 	if previousRow != nil && getPluginName(currentRow) == getPluginName(previousRow) {
// 		if currentHeader == nil {
// 			return
// 		}
// 		currentRow.SetHeader(nil)

// 	} else {
// 		if currentHeader != nil {
// 			return
// 		}
// 		headerLabel, _ := gtk.LabelNew(getPluginName(currentRow))
// 		currentRow.SetHeader(headerLabel)
// 	}

// }

// func setOptionModel(optionWidget *gtk.Box, optionModel models.OptionModel) {
// 	optionModelEncoded, _ := json.Marshal(optionModel)
// 	optionWidget.SetName(string(optionModelEncoded))
// }

// func getOptionModel(optionWidget *gtk.Widget) models.OptionModel {
// 	optionModelString, _ := optionWidget.GetName()

// 	optionModel := models.OptionModel{}
// 	json.Unmarshal([]byte(optionModelString), &optionModel)
// 	return optionModel
// }

// func getOptionWidget(row *gtk.ListBoxRow) *gtk.Widget {
// 	currentOptionInterface, _ := row.GetChild()
// 	return currentOptionInterface.ToWidget()
// }

// func getPluginName(row *gtk.ListBoxRow) string {
// 	optionWidget := getOptionWidget(row)
// 	optionModel := getOptionModel(optionWidget)
// 	return optionModel.PluginName
// }

// func OnOptionListKeyPress(optionList *gtk.ListBox, event *gdk.Event) {
// 	key := gdk.EventKeyNewFromEvent(event)

// 	// Propagate key_press_event to option on activate
// 	if key.KeyVal() == gdk.KEY_Return {
// 		selectedListBoxRow := optionList.GetSelectedRow()
// 		optionInterface, _ := selectedListBoxRow.GetChild()
// 		option := optionInterface.ToWidget()
// 		option.Event(event)
// 		return
// 	}
// }

func (this *OptionList) SetFilterFunction(searchBar *SearchBar) {
	this.SetFilterFunc(func(row *gtk.ListBoxRow) bool {
		query := strings.ToLower(searchBar.Text())
		queryParts := strings.Split(query, " ")

		optionWidget := getOptionWidget(row)
		optionModel := getOptionModel(optionWidget)

		searchTerms := []string{
			strings.ToLower(optionModel.PluginName),
			strings.ToLower(optionModel.Title),
		}

		for _, searchTerm := range searchTerms {
			for _, queryPart := range queryParts {
				if strings.Contains(searchTerm, queryPart) {
					return true
				}
			}
		}

		return false
	})
}
