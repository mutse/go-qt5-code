package main

import (
    "github.com/salviati/go-qt5/qt5"
)

func main() {
    qt5.Main(func() {
        w := qt5.NewWidget()
        //w.SetWindowTitle("Qt5 Hello With Golang Binding")
        l := qt5.NewLabel()
        l.SetText("<center>Hello Go-qt5!</center>")
        l.SetSizev(200, 50)
        defer w.Close()
        l.Show()
        qt5.Run()
    })
}
