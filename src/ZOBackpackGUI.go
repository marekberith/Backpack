package main

import (
	"./zo"
	"github.com/therecipe/qt/widgets"
	"strconv"
)

var (
	mainWindow      *widgets.QWidget
	mainLayout      *widgets.QVBoxLayout
	resultLayout    *widgets.QHBoxLayout
	vectorLine      *widgets.QLineEdit
	weightLine      *widgets.QLineEdit
	itemsCount      *widgets.QLineEdit
	calculateMatrix *widgets.QPushButton
	table           *widgets.QTableWidget
	result          *widgets.QLabel
)

func setTable(matrix zo.Matrix, weight int) {
	if matrix == nil {
		return
	}
	table.SetRowCount(len(matrix) - 1)
	table.SetColumnCount(weight)
	for i, row := range matrix {
		if i == 0 {
			continue
		}
		for j, column := range row {
			if j == 0 {
				continue
			}
			table.SetItem(i-1, j-1, widgets.NewQTableWidgetItem2(strconv.Itoa(column.Value), 0))
		}
	}
}

func setLayout() {
	mainWindow = widgets.NewQWidget(nil, 0)
	mainWindow.SetWindowTitle("Zero-One Backpack Calculator")
	mainWindow.SetMinimumWidth(500)
	mainWindow.SetMinimumHeight(400)

	mainLayout = widgets.NewQVBoxLayout()
	vectorLine = widgets.NewQLineEdit(nil)
	vectorLine.SetPlaceholderText("Insert tuples (price, weight), separated by a semicolon, that represent elements added into backpack.")
	weightLine = widgets.NewQLineEdit(nil)
	weightLine.SetPlaceholderText("Insert maximum weight of backpack.")
	itemsCount = widgets.NewQLineEdit(nil)
	itemsCount.SetPlaceholderText("Insert number of tuples.")

	result = widgets.NewQLabel(nil, 0)

	table = widgets.NewQTableWidget(nil)
	table.HorizontalHeader().SetSectionResizeMode(widgets.QHeaderView__Stretch)
	table.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	calculateMatrix = widgets.NewQPushButton(nil)
	calculateMatrix.SetText("Calculate matrix")
	calculateMatrix.ConnectClicked(func(checked bool) {
		weight, _ := strconv.Atoi(weightLine.Text())
		itemsCnt, _ := strconv.Atoi(itemsCount.Text())
		matrix, resultText := zo.GetZeroOneMatrix(vectorLine.Text(), weight, itemsCnt)
		result.SetText(resultText)
		setTable(matrix, weight)
	})

	mainLayout.AddWidget(vectorLine, 0, 0)
	mainLayout.AddWidget(weightLine, 0, 0)
	mainLayout.AddWidget(itemsCount, 0, 0)
	mainLayout.AddWidget(calculateMatrix, 0, 0)
	resultLayout = widgets.NewQHBoxLayout()
	resultLayout.AddStretch(1)
	resultLayout.AddWidget(result, 0, 0)
	resultLayout.AddStretch(1)

	mainLayout.AddLayout(resultLayout, 0)
	mainLayout.AddWidget(table, 0, 0)

	mainWindow.SetLayout(mainLayout)
	mainWindow.Show()
}
