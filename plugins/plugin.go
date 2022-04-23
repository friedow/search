package plugins

type PluginOption interface {
	OnActivate()
	PluginName() string
}
