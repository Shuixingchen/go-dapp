package utils

var FunctionHashes = map[string]string{
	"transfer(address,uint256)":             "a9059cbb",
	"transferFrom(address,address,uint256)": "23b872dd",
	"totalSupply()":                         "18160ddd",
	"balanceOf(address)":                    "70a08231",
	"approve(address,uint256)":              "095ea7b3",
	"allowance(address,address)":            "dd62ed3e",
	"getApproved(uint256)":                  "081812fc",
	"isApprovedForAll(address,address)":     "e985e9c5",
	"ownerOf(uint256)":                      "6352211e",
	"setApprovalForAll(address,bool)":       "a22cb465",
}

var ERC20FunctionSig = []string{
	FunctionHashes["transfer(address,uint256)"],
	FunctionHashes["transferFrom(address,address,uint256)"],
	FunctionHashes["totalSupply()"],
	FunctionHashes["balanceOf(address)"],
	FunctionHashes["approve(address,uint256)"],
	FunctionHashes["allowance(address,address)"],
}
