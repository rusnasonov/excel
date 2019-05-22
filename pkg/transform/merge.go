package transform

import (
	"sync"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
)

// Merge2 ...
func Merge2(sources []string, destination string) error {
	destXls := excelize.NewFile()
	destXls.NewSheet("Sheet1")
	needHeader := false
	totalRows := 2
	var wg sync.WaitGroup
	wg.Add(len(sources))

	if needHeader {
		var sheet string
		sourceXls, err := excelize.OpenFile(sources[0])
		if err != nil {
			return err
		}
		for _, name := range sourceXls.GetSheetMap() {
			sheet = name
			break
		}
		rows := sourceXls.GetRows(sheet)
		destXls.SetSheetRow("Sheet1", "A1", &rows[0])
	}
	for _, srcPath := range sources {
		var sheet string
		sourceXls, err := excelize.OpenFile(srcPath)
		if err != nil {
			return err
		}
		for _, name := range sourceXls.GetSheetMap() {
			sheet = name
			break
		}
		rows := sourceXls.GetRows(sheet)
	
		for i:=1; i < len(rows); i++ {
			destXls.SetSheetRow("Sheet1", fmt.Sprintf("A%d", totalRows),  &rows[i])
			totalRows++
		}
	}
	fmt.Printf("total rows: %d", totalRows)
	err := destXls.SaveAs(destination)
	if err != nil {
		return err
	}
	return nil
}

// Merge ...
func Merge(sources []string, destination string) error {
	dstFile :=xlsx.NewFile()
	needHeader := false
	sheet, err := dstFile.AddSheet("Sheet1")
    if err != nil {
        return err
	}
	if needHeader {
		srcFile, err := xlsx.OpenFile(sources[0])
		if err != nil {
			return err
		}
		srcSheet := srcFile.Sheets[0]
		row := srcSheet.Rows[0]
		dstRow := sheet.AddRow()
		for _, cell := range row.Cells {
			c := dstRow.AddCell()
			c.Value = cell.String()
		}
	}
	for _, srcPath := range sources {
		srcFile, err := xlsx.OpenFile(srcPath)
		if err != nil {
			return err
		}
		srcSheet := srcFile.Sheets[0]
		for _, row := range srcSheet.Rows {
			dstRow := sheet.AddRow()
			for _, cell := range row.Cells {
				c := dstRow.AddCell()
				c.Value = cell.String()
			}
		}
	}
	err = dstFile.Save(destination)
    if err != nil {
        return err
	}
	return nil
}