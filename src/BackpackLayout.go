package main

import (
	"./bp"
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
	algorithmPicker *widgets.QComboBox
	calculateResult *widgets.QPushButton
	table           *widgets.QTableWidget
	result          *widgets.QLabel
)

func setTable(matrix bp.TwoDimMatrix, weight int) {
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

func setTableOneDim(matrix bp.Vector) {
	if matrix == nil {
		return
	}
	table.SetRowCount(1)
	table.SetColumnCount(len(matrix))
	for i, column := range matrix {
		table.SetItem(0, i, widgets.NewQTableWidgetItem2(strconv.Itoa(column.Value), 0))
	}
}

func setLayout() {
	mainWindow = widgets.NewQWidget(nil, 0)
	mainWindow.SetWindowTitle("Backpack Calculator")
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
	calculateResult = widgets.NewQPushButton(nil)
	calculateResult.SetText("Calculate matrix")
	calculateResult.ConnectClicked(func(checked bool) {
		weight, _ := strconv.Atoi(weightLine.Text())
		itemsCnt, _ := strconv.Atoi(itemsCount.Text())
		if algorithmPicker.CurrentText() == "Zero-One Backpack" {
			matrix, resultText := bp.SolveZeroOneBackpack(vectorLine.Text(), weight, itemsCnt)
			setTable(matrix, weight)
			result.SetText(resultText)
		} else if algorithmPicker.CurrentText() == "Unbounded backpack" {
			matrix, resultText := bp.SolveUnboundedBackpack(vectorLine.Text(), weight, itemsCnt)
			setTableOneDim(matrix)
			result.SetText(resultText)
		}
	})

	algorithmPicker = widgets.NewQComboBox(nil)
	comboBoxOptions := []string{"Zero-One Backpack", "Unbounded backpack"}
	algorithmPicker.AddItems(comboBoxOptions)

	mainLayout.AddWidget(vectorLine, 0, 0)
	mainLayout.AddWidget(weightLine, 0, 0)
	mainLayout.AddWidget(itemsCount, 0, 0)
	mainLayout.AddWidget(algorithmPicker, 0, 0)
	mainLayout.AddWidget(calculateResult, 0, 0)
	resultLayout = widgets.NewQHBoxLayout()
	resultLayout.AddStretch(1)
	resultLayout.AddWidget(result, 0, 0)
	resultLayout.AddStretch(1)

	mainLayout.AddLayout(resultLayout, 0)
	mainLayout.AddWidget(table, 0, 0)

	mainWindow.SetLayout(mainLayout)
	mainWindow.Show()
}
