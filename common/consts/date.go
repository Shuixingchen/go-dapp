package consts

const (
	// UTCDatetime : 2006-01-02 15:04:05
	UTCDatetime = "2006-01-02 15:04:05"
	UTCMonth    = "200601"
	// UTCDay : 20060102
	UTCDay             = "20060102"
	UTCHour            = "2006010215"
	UTCDatetimeDay     = "2006-01-02"
	UTCDatetimeMinute  = "2006-01-02 15:04:00"
	UTCDatetimeHour    = "2006-01-02 15:00:00"
	UTCDatetimeOnlyDay = "2006-01-02 00:00:00"
)

var DateMap = map[string]string{
	"hour12":    "",
	"half_day":  "",
	"day":       "",
	"day3":      "",
	"week":      "",
	"day12":     "",
	"month":     "",
	"month3":    "",
	"half_year": "",
	"year":      "",
	"year3":     "",
	"year6":     "",
	"all":       "",
}

var DateFrequencyMap = map[string]int{
	"day":    6,
	"day3":   5,
	"week":   4,
	"month":  3,
	"month3": 2,
	"year":   1,
	"all":    0,
}

// CoinStartDateMap 币种首次爆块日期
var CoinStartDateMap = map[string]string{
	"btc": "2009-01-03",
	// 2017-08-01
	"bch": "2009-01-03",
	// 2018-11-15
	"bsv": "2009-01-03",
	"ltc": "2011-10-07",
	"eth": "2015-07-30",
	"etc": "2015-06-30",
}
