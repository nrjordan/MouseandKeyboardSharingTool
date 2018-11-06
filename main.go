package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"mkShareReceiver/client"
	"mkShareReceiver/server"
)

func main() {
	gtk.Init(nil)

	win := setupWindow("Mouse and Keyboard Sharing")
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid: ", err)
	}
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	nb, err := gtk.NotebookNew()
	if err != nil {
		log.Fatal("unable to create notebook: ", err)
	}

	nb.SetHExpand(true)
	nb.SetVExpand(true)
	servTab, err := gtk.LabelNew("Server")
	if err != nil {
		log.Fatal("Unable to create server label: ", err)
	}
	clientTab, err := gtk.LabelNew("Client")
	if err != nil {
		log.Fatal("Unable to create client label: ", err)
	}

	logScroll, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		log.Fatal("unable to create scroll window: ", err)
	}

	logBox, err := gtk.TextViewNew()
	if err != nil {
		log.Fatal("Unable to create textview: ", err)
	}

	ipLabel, err := gtk.LabelNew("IP Address: ")
	if err != nil {
		log.Fatal("Unable to create IP label: ", err)
	}

	ipTextbox, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create IP label: ", err)
	}

	btnLabel := "Start"

	btn := setupBtn(btnLabel, func() {
		if nb.GetCurrentPage() == 0 {
			if btnLabel == "Start" {
				s, err := gdk.ScreenGetDefault()
				if err != nil {
					fmt.Println("Error getting screen size", err)
				}
				res := s.GetResolution()
				server.StartServer(res)
				btnLabel = "Stop"
			} else {
				server.StopServer()
				btnLabel = "Start"
			}
			fmt.Println("Server", btnLabel)
		} else if nb.GetCurrentPage() == 1 {
			ip, err := ipTextbox.GetText()
			if err != nil {
				fmt.Println("Error getting text...")
			}
			client.StartClient(ip)
			fmt.Println("Client")
		}
	})

	logBox.SetHExpand(true)
	logBox.SetVExpand(false)
	logBox.SetEditable(false)

	logScroll.SetMarginBottom(20)
	logScroll.SetMarginTop(27)
	logScroll.Add(logBox)

	grid.Attach(nb, 1, 1, 1, 1)
	grid.Attach(logScroll, 2, 1, 1, 2)
	grid.Attach(btn, 1, 2, 2, 2)


	/*cliTab, err := gtk.LabelNew("Client")
	if err != nil {
		log.Fatal("Unable to create client label: ", err)
	}*/

	servBox := setupBox(gtk.ORIENTATION_VERTICAL)
	clientBox := setupBox(gtk.ORIENTATION_VERTICAL)


	clientBox.Add(ipLabel)
	clientBox.Add(ipTextbox)

	nb.AppendPage(servBox, servTab)
	nb.AppendPage(clientBox, clientTab)
	win.Add(grid)


	// Set the default window size.
	//win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

func setupWindow(title string) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window: ", err)
	}
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(800, 600)
	win.SetPosition(gtk.WIN_POS_CENTER)
	return win
}

func setupBox(orient gtk.Orientation) *gtk.Box {
	box, err := gtk.BoxNew(orient, 0)
	if err != nil {
		log.Fatal("Unable to create box: ", err)
	}
	return box
}

func setupBtn(label string, onClick func()) *gtk.Button {
	btn, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		log.Fatal("Unable to create button: ", err)
	}
	btn.Connect("clicked", onClick)
	return btn
}