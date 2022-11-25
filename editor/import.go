package editor

import (
	"fmt"
	//"unicode/utf8"

	//"io/ioutil"
	"os"
	"tbl-editor/utils"

	"github.com/labstack/gommon/log"
	"github.com/xuri/excelize/v2"

)

func Import(inFile, outFile string) {

	f, err := excelize.OpenFile(inFile)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := f.Rows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	buff := NewBuffer([]byte{})
	cols := []ColType{}

	if rows.Next() {
		headers, err := rows.Columns()
		if err != nil {
			return
		}
		
		buff.Write(fmt.Sprintf("%d", len(headers)), UINT32) // write column size

		for _, h := range headers {
			col := rTypeTitles[h]
			cols = append(cols, col)
			buff.Write(fmt.Sprintf("%d", col), UINT32) // write columns
		}
	}

	buff.Write("0", UINT32) // write rows size (placeholder)

	r := 0
	for rows.Next() {

		r++
		c := 0

		headers2, err := rows.Columns()
		if err != nil {
			return
		}

		for _, cell := range headers2 {

		
			col := cols[c]
			buff.Write(cell, col)
			c++
		}
	}

	index := (len(cols) + 1) * 4
	size := buff.GetOffset() - uint64(index+8)

	buff.Overwrite(utils.IntToBytes(uint64(r), 4, true), index) // write rows size
	buff.Overwrite(utils.IntToBytes(size, 4, true), index+4)    // write size

	os.WriteFile(outFile, buff.GetBytes(), 0644)
}
