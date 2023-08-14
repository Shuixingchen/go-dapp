package eth

import (
	"fmt"
	"strings"
)

var TransactionTypeStatusMap = map[string]string{
	"TYPE_0x1": "SUCCESS",
	"TYPE_0x0": "FAILED",
}

func GetTransactionStatus(t string) (status string) {
	consName := fmt.Sprintf("TYPE_%s", t)
	if v, exist := TransactionTypeStatusMap[consName]; exist {
		status = v
	} else {
		status = "UNKNOWN"
	}

	return
}

func GetAccountTransactionType(addr, sender, receiver string) (txType string) {
	if strings.EqualFold(strings.ToLower(addr), strings.ToLower(sender)) &&
		strings.EqualFold(strings.ToLower(addr), strings.ToLower(receiver)) {
		txType = "SELF"
		return
	}

	if strings.EqualFold(strings.ToLower(addr), strings.ToLower(sender)) {
		txType = "OUT"
	} else {
		txType = "IN"
	}

	return
}
