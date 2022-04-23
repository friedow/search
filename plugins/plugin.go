package plugins

type PluginOption interface {
	OnActivate()
	PluginName() string
	IsVisible(queryPart string) bool
}
