package fedach_test

import (
	"fmt"

	"github.com/thylong/fedach"
)

func ExampleMarshal() {
	routingDirectoryFile := []fedach.RoutingDirectoryRecord{
		{
			"011000015",
			"O",
			"011000015",
			"0",
			"122415",
			"000000000",
			"FEDERAL RESERVE BANK                ",
			"1000 PEACHTREE ST N.E.              ",
			"ATLANTA             ",
			"GA",
			"30309",
			"4470",
			"877",
			"372",
			"2457",
			"1",
			"1",
			"     ",
		},
	}
	encoded, err := fedach.Marshal(routingDirectoryFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", encoded)
}
