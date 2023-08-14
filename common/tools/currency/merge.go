package currency

var CoinMergeHeight = make(map[string]int64)
var RealMergeTime = make(map[string]int64)

// GetRealMergeTime 获取合并完成的时间
func GetRealMergeTime(coin string) int64 {
	time, ok := RealMergeTime[coin]
	if !ok {
		time = RealMergeTime["eth"]
	}
	return time
}

// GetCoinMergeHeight 获取币种合并高度
func GetCoinMergeHeight(coin string) int64 {
	height, ok := CoinMergeHeight[coin]
	if !ok {
		height = CoinMergeHeight["eth"]
	}
	return height
}
