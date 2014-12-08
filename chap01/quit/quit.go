package main

import (
    "github.com/salviati/go-qt5/qt5"
)

func main() {
    qt5.Main(func() {
        app := qt5.NewWidget()
        btn := qt5.NewButton()
        btn.SetText("Quit")
        btn.SetSizev(200, 50)
        btn.OnClicked(func() {
            btn.Close()
        })
        defer app.Close()
        btn.Show()
        qt5.Run()
    })
}

