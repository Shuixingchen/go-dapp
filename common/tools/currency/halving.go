// Package currency 减半相关操作
package currency

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Shuixingchen/go-dapp/common/consts"
)

// CoinHalvingHeight zec第一次减半高度参见 https://github.com/zcash/zips/blob/master/zip-0208.rst#preblossom-protocol
// zip208
// FoundersRewardLastBlockHeight := max({ height ⦂ N | Halving(height) < 1 })
// Halving(h) is defined as floor(f(h)) where f is a strictly increasing rational
// function, so it's sufficient to solve for f(height) = 1 in the rationals and
// then take ceiling(height - 1).
// H := blossom activation height; SS := SubsidySlowStartShift(); R := BLOSSOM_POW_TARGET_SPACING_RATIO
// preBlossom:
// 1 = (height - SS) / preInterval
// height = preInterval + SS
// postBlossom:
// 1 = (H - SS) / preInterval + (height - H) / postInterval
// height = H + postInterval - (H - SS) * (postInterval / preInterval)
// height = H + postInterval - (H - SS) * R
// Note: This depends on R being an integer
// https://github.com/zcash/zcash/blob/e93586a0c43d0d77d641f692c22798d9b00bb698/src/consensus/params.cpp#L39
var CoinHalvingHeight = make(map[string]int64)
var RealHalvingTime = make(map[string]int64)

// GetRealHalvingTime 获取已经完成减半的时间
func GetRealHalvingTime(coin string) int64 {
	time, ok := RealHalvingTime[coin]
	if !ok {
		time = RealHalvingTime["btc"]
	}
	return time
}

// GetCoinHalvingHeightByConfig 获取币种减半高度
func GetCoinHalvingHeightByConfig(coin string) int64 {
	height, ok := CoinHalvingHeight[coin]
	if !ok {
		height = CoinHalvingHeight["btc"]
	}
	return height
}

// GetCoinNextHalvingHeight 获取币种下一次减半高度
func GetCoinNextHalvingHeight(coin string, height int64) (halvingHeight int64) {
	switch coin {
	case consts.CoinTypeEthereumClassic:
		halvingHeight = (height/5000000 + 1) * 5000000
	case consts.CoinTypeDash:
		halvingHeight = (height/210240 + 1) * 210240
	case consts.CoinTypeZcash:
		// Todo:
	case consts.CoinTypeBitcoin:
		halvingHeight = (height/210000 + 1) * 210000
	case consts.CoinTypeBitcoinCash:
		halvingHeight = (height/210000 + 1) * 210000
	case consts.CoinTypeBitcoinSV:
		halvingHeight = (height/210000 + 1) * 210000
	case consts.CoinTypeLitecoin:
		halvingHeight = (height/840000 + 1) * 840000
	default:
		halvingHeight = GetCoinHalvingHeightByConfig(coin)
	}

	if halvingHeight == 0 {
		halvingHeight = GetCoinHalvingHeightByConfig(coin)
	}

	return halvingHeight
}

// GetCoinMiningReward 获取币种减半前出块奖励
func GetCoinMiningReward(coin string, height int64) float64 {
	var reward float64

	switch coin {
	case consts.CoinTypeEthereumClassic:
		reward = GetEtcMiningReward(height)
	case consts.CoinTypeZcash:
		reward = GetZecMiningReward(height)
	case consts.CoinTypeLitecoin:
		reward = GetLtcMiningReward(height)
	default:
		reward = GetBtcMiningReward(height)
	}
	return reward
}

// GetCoinHalvingMiningReward 获取币种减半后出块奖励
func GetCoinHalvingMiningReward(coin string, miningReward float64) float64 {
	var reward float64

	switch coin {
	case consts.CoinTypeEthereumClassic:
		reward = miningReward * 8 / 10
	case consts.CoinTypeDash:
		reward = miningReward * 9286 / 10000
		reward, _ = strconv.ParseFloat(fmt.Sprintf("%.8f", reward), 64)
	default:
		reward = miningReward * 5 / 10
	}
	return reward
}

// GetHalvingIncomecoinOrIncomeOptimizeCoin 获取减半后pps/fpps
func GetHalvingIncomecoinOrIncomeOptimizeCoin(coin, income string) string {
	var pps float64
	if income == "" {
		return "0"
	}
	pps, err := strconv.ParseFloat(income, 64)
	if err != nil {
	}

	switch coin {
	case consts.CoinTypeEthereumClassic:
		pps = pps * 8 / 10
	case consts.CoinTypeDash:
		pps = pps * 9286 / 10000
	default:
		pps = pps * 5 / 10
	}
	return strconv.FormatFloat(pps, 'f', -1, 64)
}

// RecalculateIncome 计算当前24h每T/G/M/K收益，由于获取的Income数据10分钟更新一次，在减半完成时需要实时计算一次，以及考虑遇到redis没有数据的情况
//
//	/* ** Dash ** */
//	https://docs.dash.org/en/latest/introduction/features.html#total-coin-emission
//	Dash每210240块（约383.25天）将块奖励减少十四分之一（约7.14％）, 10% 不固定的发行金额不计入
//	最终矿工矿池分块奖励 45%
//	计算：
//	$btc_number_in_one_block = pow(13/14, floor($latest_block['height'] / 210240)) * (5 * 0.9);
//	$income = pow(10, 9) * 86400 / (pow(2, 32) * $latest_block['diff']) * $btc_number_in_one_block * 0.5;
//	/* ** Zec ** */
//	前4年产生的区块奖励的20%归zcash公司（总量的10%），80%归矿工。4年后的区块奖励全部归矿工。
func RecalculateIncome(coin string, reward, diff float64, height int64) string {
	var Income float64

	switch coin {
	case consts.CoinTypeEthereumClassic:
		Income = 1e+6 * reward * 86400 / diff
	case consts.CoinTypeDash:
		Income = 1e+9 * reward * 0.45 * 86400 / diff / math.Pow(2, 32)
	case consts.CoinTypeZcash:
		if height < CoinHalvingHeight["zec"] {
			Income = 1e+3 * reward * 0.8 * 86400 / diff / 8192
		} else {
			Income = 1e+3 * reward * 86400 / diff / 8192
		}
	default:
		Income = 1e+12 * reward * 86400 / diff / math.Pow(2, 32)
	}

	return fmt.Sprintf("%.21f", Income)
}

// GetEtcMiningReward 获取 etc 挖矿奖励
func GetEtcMiningReward(height int64) float64 {
	blockEra := (int(height)-1)/5000000 + 1
	// The blockEra is 2 in 2018.
	// Avoid calculations by giving the result directly.
	if blockEra == 2 {
		return 4
	}
	var reward float64 = 5
	for i := 1; i < blockEra; i++ {
		// ECIP-1017: all rewards will be reduced at a constant rate of 20% upon
		// entering a new Era. reward *= 0.8 (avoid converts to float)
		reward = reward * 8 / 10
	}

	return reward
}

// GetBtcMiningReward 获取 btc 挖矿奖励
func GetBtcMiningReward(height int64) float64 {
	blockEra := int(height)/210000 + 1
	var reward float64 = 50
	for i := 1; i < blockEra; i++ {
		reward /= 2
	}
	return reward
}

func GetLtcMiningReward(height int64) float64 {
	blockEra := int(height)/840000 + 1
	var reward float64 = 50
	for i := 1; i < blockEra; i++ {
		reward /= 2
	}
	return reward
}

// GetZecMiningReward 获取 zec 挖矿奖励
func GetZecMiningReward(height int64) float64 {
	// 前10000个块，
	// 块1的奖励是0.000625币，随着高度线性递增，增加幅度是0.000625币，块9999的奖励是6.249375币。
	// 块10000～19999，
	// 块10000的奖励是6.250625，随着高度线性递增，增加幅度是0.000625币，块19999的奖励是12.5币。
	// 正常块奖励从块20000开始，奖励是12.5个币，奖励按照每4年减半。
	var basicBlockReward = 0.000625
	var reward float64
	if height < 10000 {
		reward = float64(height) * basicBlockReward
	} else if height == 10000 {
		reward = 6.250625
	} else if height > 10000 && height < 20000 {
		reward = 6.250625 + float64(height-10000)*basicBlockReward
	} else if height >= 20000 && height < 653600 { // 北京时间：12月12日，区块高度：653,600  升级内容：减少全网验证时间，从原先的150秒缩短到75秒，单个区块奖励已由之前的12.5个降为当前的6.25个。
		reward = 12.5
	} else if height >= 653600 && height < 1046400 {
		reward = 6.25
	} else {
		reward = 3.125
	}

	return reward
}

// CalculateIncome 计算24h的每T/M收益
func CalculateIncome(coin string, reward float64, diff []float64, height int64) string {
	var Income float64

	switch coin {
	case consts.CoinTypeEthereumClassic:
		for _, v := range diff {
			Income += 1e+6 * reward / v
		}
	case consts.CoinTypeDash:
		for _, v := range diff {
			Income += 1e+9 * reward * 0.45 / v / math.Pow(2, 32)
		}
	case consts.CoinTypeZcash:
		for _, v := range diff {
			Income += 1e+3 * reward / v / 8192
		}
		if height < CoinHalvingHeight["zec"] {
			Income *= 0.8
		}
	default:
		for _, v := range diff {
			Income += 1e+12 * reward / v / math.Pow(2, 32)
		}
	}
	Income = Income * 86400 / float64(len(diff))

	return fmt.Sprintf("%.21f", Income)
}
