package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(980, 800)
	window.SetWindowTitle("Pythia")

	widget := widgets.NewQWidget(nil, 0)
	window.SetCentralWidget(widget)

	pa_label := widgets.NewQLabel2("Public Address", widget, 0)
	pa_input := widgets.NewQLineEdit(widget)
	pa_input.SetPlaceholderText("Enter your public address")

	priv_label := widgets.NewQLabel2("Private key", widget, 0)
	priv_input := widgets.NewQLineEdit(widget)
	priv_input.SetPlaceholderText("Enter your private key")

	sh_label := widgets.NewQLabel2("Server Host", widget, 0)
	sh_input := widgets.NewQLineEdit(widget)
	sh_input.SetPlaceholderText("Enter the server host address for your datasever")

	port_label := widgets.NewQLabel2("Server Port", widget, 0)
	port_input := widgets.NewQLineEdit(widget)
	port_input.SetPlaceholderText("Enter the server port for your dataserver")

	form := widgets.NewQFormLayout(widget)
	form.AddRow(pa_label, pa_input)
	form.AddRow(priv_label, priv_input)
	form.AddRow(sh_label, sh_input)
	form.AddRow(port_label, port_input)
	widget.SetLayout(form)

	button := widgets.NewQPushButton2("submit", nil)
	button.ConnectClicked(func(checked bool) {
		widgets.QMessageBox_Information(
			nil, "Address Submitted",
			"Your address has been submitted to ./config.json", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	widget.Layout().AddWidget(button)

	window.Show()

	app.Exec()
}
