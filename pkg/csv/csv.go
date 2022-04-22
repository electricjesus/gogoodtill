package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func CSV() {
	f, err := os.Open("tests/testdata/products.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}
