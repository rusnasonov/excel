package transform

import "testing"

func TestMerge(t *testing.T) {
	sources := []string{
		"/Users/perseus/Documents/inbox_excel/0001.xlsx",
		"/Users/perseus/Documents/inbox_excel/0002.xlsx",
		"/Users/perseus/Documents/inbox_excel/0003.xlsx",
		"/Users/perseus/Documents/inbox_excel/0004.xlsx",
		"/Users/perseus/Documents/inbox_excel/0005.xlsx",
		"/Users/perseus/Documents/inbox_excel/0006.xlsx",
		"/Users/perseus/Documents/inbox_excel/0007.xlsx",
	}
	destination := "/Users/perseus/Documents/inbox_excel/merged.xlsx"
	
	err := Merge(sources, destination)

	if err != nil {
		t.Error(err)
	}
}
