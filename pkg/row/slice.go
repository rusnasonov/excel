package row

import (
	"github.com/tealeg/xlsx"
)


// Slice ...
func Slice(source string, from int, to int) ([]*xlsx.Row, error) {
	srcFile, err := xlsx.OpenFile(source)
	if err != nil {
		return nil, err
	}
	srcSheet := srcFile.Sheets[0]
	if(from < 0){
		from = 0
	}
	if(to < 0){
		to = len(srcSheet.Rows)
	}
	rows := srcSheet.Rows[from:to]
	return rows, nil
}