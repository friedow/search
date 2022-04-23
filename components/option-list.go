package components

import (
	"friedow/tucan-search/plugins"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type OptionList struct {
	*gtk.ScrolledWindow

	optionList struct {
		*gtk.ListBox

		options []*plugins.PluginOption
	}
}

func NewOptionList() *OptionList {
	this := OptionList{}

	this.optionList.ListBox = gtk.NewListBox()
	this.optionList.SetHeaderFunc(this.setHeader)
	this.optionList.ConnectRowActivated(this.onActivate)

	this.optionList.options = []*plugins.PluginOption{}

	// TODO: add list of plugin options to this component to acess them in onActivate handlers
	gitRepositoryPluginOptions := plugins.NewGitRepositoriesPluginOptions()
	for _, gitRepositoryPluginOption := range gitRepositoryPluginOptions {
		this.optionList.Append(gitRepositoryPluginOption)
		var pluginOption plugins.PluginOption = gitRepositoryPluginOption
		this.optionList.options = append(this.optionList.options, &pluginOption)
	}

	this.ScrolledWindow = gtk.NewScrolledWindow()
	this.ScrolledWindow.SetMinContentHeight(700)
	this.ScrolledWindow.SetChild(this.optionList)

	return &this
}

func (this *OptionList) setHeader(currentRow *gtk.ListBoxRow, previousRow *gtk.ListBoxRow) {
	currentHeader := currentRow.Header()

	if previousRow != nil && pluginName(currentRow) == pluginName(previousRow) {
		if currentHeader == nil {
			return
		}
		currentRow.SetHeader(nil)

	} else {
		if currentHeader != nil {
			return
		}
		newHeader := gtk.NewLabel(pluginName(currentRow))
		currentRow.SetHeader(newHeader)
	}
}

func (this *OptionList) onActivate(row *gtk.ListBoxRow) {
	pluginOption(row).OnActivate()
}

func (this *OptionList) selectFirstRow() {
	firstRow := this.optionList.RowAtIndex(0)
	if firstRow != nil {
		this.optionList.SelectRow(firstRow)
	}
}

func (this *OptionList) selectPreviousRow() {
	currentRow := this.optionList.SelectedRow()
	if currentRow == nil {
		this.selectFirstRow()
		return
	}

	nextRow := this.optionList.RowAtIndex(currentRow.Index() - 1)
	if nextRow != nil {
		this.optionList.SelectRow(nextRow)
	}
}

func (this *OptionList) selectNextRow() {
	currentRow := this.optionList.SelectedRow()
	if currentRow == nil {
		this.selectFirstRow()
		return
	}

	nextRow := this.optionList.RowAtIndex(currentRow.Index() + 1)
	if nextRow != nil {
		this.optionList.SelectRow(nextRow)
	}
}

func pluginOption(row *gtk.ListBoxRow) plugins.PluginOption {
	row.setn
	return row.Child().(plugins.PluginOption)
}

func pluginName(row *gtk.ListBoxRow) string {
	return pluginOption(row).PluginName()
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
		this.optionList.SelectedRow().Activate()
		return true
	}

	this.optionList.InvalidateFilter()
	return false
}

// xxxxxxxxxxxxxxxxxxxxxxxxxx

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

// func (this *OptionList) SetFilterFunction(searchBar *SearchBar) {
// 	this.SetFilterFunc(func(row *gtk.ListBoxRow) bool {
// 		query := strings.ToLower(searchBar.Text())
// 		queryParts := strings.Split(query, " ")

// 		optionWidget := getOptionWidget(row)
// 		optionModel := getOptionModel(optionWidget)

// 		searchTerms := []string{
// 			strings.ToLower(optionModel.PluginName),
// 			strings.ToLower(optionModel.Title),
// 		}

// 		for _, searchTerm := range searchTerms {
// 			for _, queryPart := range queryParts {
// 				if strings.Contains(searchTerm, queryPart) {
// 					return true
// 				}
// 			}
// 		}

// 		return false
// 	})
// }
