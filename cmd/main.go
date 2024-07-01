package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ElnurKoke/crypto"
)

func main() {
	a := app.New()
	w := a.NewWindow("cipher")
	w.Resize(fyne.NewSize(400, 400))

	var encType string

	word := widget.NewLabel("Heyya")
	entry := widget.NewMultiLineEntry()
	entry.SetPlaceHolder("Message")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	radios := widget.NewRadioGroup([]string{"sha-1", "aes128"}, func(s string) {
		encType = s
	})

	btn := widget.NewButton("Run", func() {
		switch encType {
		case "sha-1":
			word.SetText(crypto.Sha1(entry.Text))
		case "aes128":
			if len(key.Text) != 16 {
				word.SetText("Aes128 need 16 byte key!")
			} else {
				word.SetText(crypto.AES128enc(entry.Text, key.Text))
			}
		default:
			word.SetText("Please select type of cryptoalghoritm!")
		}
	})

	w.SetContent(container.NewVBox(
		word,
		radios,
		entry,
		key,
		btn,
	))
	w.ShowAndRun()
}
