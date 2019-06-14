package transform

import (
	"github.com/tealeg/xlsx"
)

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