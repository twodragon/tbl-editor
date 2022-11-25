### Installation

```go
go build
```


### How to Use?
You can convert a .tbl file into excel table as following:
```
./tbl-editor export -i [tbl_file] -o [xlsx_file]
```

After tbl file is exported, you can edit the table on Excel

Finally, you can convert an excel table into .tbl file as following:
```
./tbl-editor import -i [xlsx_file] -o [tbl_file]
```

### Example
```
./tbl-editor export -i tb_cashshop.tbl -o cashshop.xlsx
./tbl-editor import -i cashshop.xlsx -o tb_cashshop.tbl
```
