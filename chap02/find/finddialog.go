package main

import (
    "fmt"
    "github.com/salviati/go-qt5/qt5"
)

type FindDialog struct {
    qt5.Dialog
    label *qt5.Label
    ledit *qt5.LineEdit
    checkbox *qt5.CheckBox
    backward *qt5.CheckBox
    findbtn *qt5.Button
    closebtn *qt5.Button
}

func (d *FindDialog) Init() *FindDialog {
    if d.Dialog.Init() == nil {
        return nil
    }

    d.label = qt5.NewLabel()
    d.label.SetText("Find &what:")
    d.ledit = qt5.NewLineEdit()
    d.label.SetBuddy(d.ledit)

    d.checkbox = qt5.NewCheckBox()
    d.checkbox.SetText("Match &case")
    d.backward = qt5.NewCheckBox()
    d.backward.SetText("Search &backward")

    d.findbtn = qt5.NewButton()
    d.findbtn.SetText("&Find")
    d.findbtn.SetDefault(true)
    d.findbtn.SetEnabled(false)
    d.findbtn.OnClicked(func () {
        d.Findclicked()
    })
    d.closebtn = qt5.NewButton()
    d.closebtn.SetText("Close")
    d.closebtn.OnClicked(func () {
        d.Close()
    })

    d.ledit.OnTextChanged(func (s string) {
        d.EnableFindButton(s)
    })

    topleftlayout := qt5.NewHBoxLayout()
    topleftlayout.AddWidget(d.label)
    topleftlayout.AddWidget(d.ledit)

    leftlayout := qt5.NewVBoxLayout()
    leftlayout.AddLayout(topleftlayout)
    leftlayout.AddWidget(d.checkbox)
    leftlayout.AddWidget(d.backward)

    rightlayout := qt5.NewVBoxLayout()
    rightlayout.AddWidget(d.findbtn)
    rightlayout.AddWidget(d.closebtn)
    rightlayout.AddStretch(2)

    mainlayout := qt5.NewHBoxLayout()
    mainlayout.AddLayout(leftlayout)
    mainlayout.AddLayout(rightlayout)
    d.SetLayout(mainlayout)

    d.SetWindowTitle("Find")

    return d
}

func main() {
    qt5.Main(ui_main)
}

func ui_main() {
    exit := make(chan bool)
    go func() {
		qt5.OnInsertObject(func(v interface{}) {
			//fmt.Println("add item", v)
		})
		qt5.OnRemoveObject(func(v interface{}) {
			//fmt.Println("remove item", v)
		})
        w := new(FindDialog).Init()
        defer w.Close()

        w.SetSizev(300, 100)
        w.Show()
        <- exit
    }()
    qt5.Run()
    exit <- true
}

func (d *FindDialog) Findclicked() {
    text := d.ledit.Text()
    fmt.Println(text)
}

func (d *FindDialog) EnableFindButton(text string) {
    d.findbtn.SetEnabled(text != "")
}
