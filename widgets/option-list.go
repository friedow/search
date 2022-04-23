package widgets

// import (
// 	"C"
// 	"friedow/tucan-search/plugins"

// 	"github.com/gotk3/gotk3/gdk"
// 	"github.com/gotk3/gotk3/gtk"
// )
// import (
// 	"encoding/json"
// 	"friedow/tucan-search/models"
// 	"strings"
// )

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

// func SetFilterFunction(optionList *gtk.ListBox, searchBar *gtk.Entry) {
// 	optionList.SetFilterFunc(func(row *gtk.ListBoxRow) bool {
// 		query, _ := searchBar.GetText()
// 		query = strings.ToLower(query)
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
