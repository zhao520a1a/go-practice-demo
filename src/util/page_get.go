package util

import (
	"context"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func ListByPage(ctx context.Context) {
	uids := make([]int64, 100)
	pageSize := 50
	for i := 0; i < len(uids); i += pageSize {
		uidList := uids[i:MinInt(i+pageSize, len(uids))]
		//rpc call with uidList
		fmt.Printf(uidList)
	}
	return
}
