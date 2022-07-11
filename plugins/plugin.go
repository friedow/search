package plugins

import (
	"sync"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type PluginOption interface {
	gtk.Widgetter

	OnActivate()
	PluginName() string
	IsVisible(queryPart string) bool
}

func PluginOptions() []PluginOption {
	newPluginOptionsFunctions := []func() []PluginOption{
		newClockPluginOptions,
		newSystemMonitorPluginOptions,
		newOpenWindowsPluginOptions,
		newApplicationsPluginOptions,
		newGitRepositoriesPluginOptions,
	}

	var pluginOptionsPromises sync.WaitGroup
	pluginOptionsChannel := make(chan []PluginOption, len(newPluginOptionsFunctions))

	for _, newPluginOptionsFunction := range newPluginOptionsFunctions {
		pluginOptionsPromises.Add(1)

		go func(function func() []PluginOption, result chan []PluginOption) {
			defer pluginOptionsPromises.Done()
			result <- function()
		}(newPluginOptionsFunction, pluginOptionsChannel)
	}
	pluginOptionsPromises.Wait()
	close(pluginOptionsChannel)

	pluginOptions := []PluginOption{}
	for resultingPluginOptions := range pluginOptionsChannel {
		pluginOptions = append(pluginOptions, resultingPluginOptions...)
	}

	return pluginOptions
}
