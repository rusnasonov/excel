# excel

Set of commands for manipulate with `.xlsx` files.

For more information use `excel --help` or `excel [command] --help`

## Example

Merge files into one — `excel transform merge *.xlsx`

Count rows — `excel row count myFile.xlsx`

Get rows slice — `excel row slice --from 2 --to 10 myFile.xlsx`

Combine with [xargs](https://en.wikipedia.org/wiki/Xargs) —   `excel transform merge *.xlsx | xargs excel row count`
