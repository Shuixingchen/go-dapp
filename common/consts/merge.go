package consts

var CoinMergeTotalDifficulty = map[string]string{
	CoinTypeEthereum:        "58750000000000000000000",
	CoinTypeEthereumClassic: "",
}

// 合并状态
var (
	NotMERGE     = "not_merge"
	MERGING      = "merging"
	MERGED       = "merged"
	MERGEFailure = "merge_failure"
)
