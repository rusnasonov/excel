package row

import (
	"github.com/tealeg/xlsx"
)


// Count ...
func Count(source string) (int, error) {
	srcFile, err := xlsx.OpenFile(source)
		if err != nil {
			return -1, err
		}
	srcSheet := srcFile.Sheets[0]
	return len(srcSheet.Rows), nil
}