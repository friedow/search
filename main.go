// package main

// import (
// 	"log"

// 	"github.com/gotk3/gotk3/gtk"
// )

// func main() {
// 	gtk.Init(nil)

// 	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
// 	if err != nil {
// 		log.Fatal("Unable to create window:", err)
// 	}
// 	win.SetTitle("Tucan Search")
// 	win.SetModal(true)
// 	win.Connect("destroy", func() {
// 		gtk.MainQuit()
// 	})

// 	app := App()

// 	win.Add(app)

// 	win.SetDefaultSize(800, 600)

// 	win.ShowAll()

// 	gtk.Main()
// }

package main

import (
	"friedow/tucan-search/views"
	"log"
	"os"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.github.friedow.tucan-search", 0)
	app.ConnectActivate(func() { activate(app) })

	glib.TimeoutAdd(2000, func() bool {
		log.Print("toast")
		return true
	})

	code := app.Run(os.Args)
	if code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	window := gtk.NewApplicationWindow(app)
	window.SetTitle("Tucan Search")
	window.SetDefaultSize(800, 600)
	window.SetModal(true)

	searchView := views.NewSearchView()
	window.SetChild(searchView.Box)

	window.Show()
}
