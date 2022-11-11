package handler

import (
	"context"
	"fmt"
)

func PenddingTx() {
	txCh := make(chan string)
	sub, err := evm.SubscribependdingTx(context.Background(), txCh)
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case txHash := <-txCh:
			fmt.Println(txHash)
		case err = <-sub.Err():
			fmt.Println(err)
		}
	}
}
