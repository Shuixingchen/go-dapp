package models

// RPCTrace represents transaction subtraces.
type RPCTrace struct {
	Action    Action `json:"action"`
	Result    Result `json:"result"`
	Error     string `json:"error"`
	Subtraces int    `json:"subtraces"`

	BlockNumber uint64 `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`

	TxHash string `json:"transactionHash"`
	TxPos  int    `json:"transactionPosition"`

	TraceAddress []int  `json:"traceAddress"`
	Type         string `json:"type"`
}

// Action represents action object in trace.
type Action struct {
	CallType string  `json:"callType"`
	From     string  `json:"from"`
	GasLimit string  `json:"gas"`
	Input    string  `json:"input"`
	To       string  `json:"to"`
	Value    *HexBig `json:"value"`

	Init string `json:"init,omitempty"`
	// for suicide
	Addr       string  `json:"address"`
	Balance    *HexBig `json:"balance"`
	RefundAddr string  `json:"refundAddress"`

	// for block reward
	Author     string `json:"author"`
	RewardType string `json:"rewardType"`
}

// Result represents result object in trace.
type Result struct {
	GasUsed string `json:"gasUsed"`
	Output  string `json:"output"`

	// contract creation address
	Addr string `json:"address"`
	Code string `json:"code"`
}
