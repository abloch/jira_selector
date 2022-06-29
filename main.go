// Demo code for the List primitive.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var table *tview.Table
var selected map[string]string

func handler(row, col int) {
	key := table.GetCell(row, 0).Text
	val := table.GetCell(row, 2).Text
	selected = make(map[string]string)
	selected[key] = val
	app.Stop()
}

func terminate(key tcell.Key) {
	app.Stop()
}

func readInput() map[string]string {
	stdin := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdin = stdin + scanner.Text()
	}

	var d map[string]string
	err := json.Unmarshal([]byte(stdin), &d)
	if err != nil {
		log.Fatal(fmt.Errorf("[Error]: input json was not an object of string keys and values:%v", err))
	}
	return d
}

func DrawTable(inputs map[string]string) map[string]string {
	app = tview.NewApplication().EnableMouse(true)
	table = tview.NewTable().SetSelectable(true, false) //.SetBorders(true)

	rowIndex := 0
	for k, v := range inputs {
		cell0 := tview.NewTableCell(k)
		cell0.SetTextColor(tcell.ColorAntiqueWhite).SetBackgroundColor(tcell.ColorDarkBlue)

		cell1 := tview.NewTableCell(v)
		cell1.SetTextColor(tcell.ColorGreen)

		table = table.InsertRow(rowIndex)
		table = table.SetCell(rowIndex, 0, cell0)
		table = table.SetCell(rowIndex, 2, cell1)

		mycell := tview.NewTableCell(" ")
		table.SetCell(rowIndex, 1, mycell)
		rowIndex += 1
	}

	table.SetSelectedFunc(handler)
	table.SetDoneFunc(terminate)
	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		log.Fatal(fmt.Errorf("[Error]: could not display application: %v", err))
	}

	return selected
}

func main() {
	inputs := readInput()
	selected := DrawTable(inputs)
	ret, err := json.Marshal(selected)
	if err != nil {
		log.Fatal(fmt.Errorf("[Error]: could not json-marshal result: %v", err))
	}
	fmt.Println(string(ret))
}
