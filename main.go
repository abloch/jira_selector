// Demo code for the List primitive.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var list *tview.List
var key, val string

func handler() {
	ind := list.GetCurrentItem()
	key, val = list.GetItemText(ind)
	time.Sleep(1 * time.Second)
	app.Stop()
}
func main() {

	stdin := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = stdin + scanner.Text()
	}

	var d map[string]string
	err := json.Unmarshal([]byte(stdin), &d)
	if err != nil {
		panic(err)
	}

	underlineStyle := tcell.Style{}.Underline(true)
	app = tview.NewApplication()
	list = tview.NewList().SetMainTextStyle(underlineStyle)

	for k, v := range d {
		list.AddItem(k, v, ' ', handler)
	}
	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})
	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	ret, err := json.Marshal(map[string]string{key: val})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ret))
}
