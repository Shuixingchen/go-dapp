package date

import (
	"regexp"
	"strconv"
	"time"

	"github.com/Shuixingchen/go-dapp/common/consts"
	"github.com/Shuixingchen/go-dapp/common/pkg"
	log "github.com/sirupsen/logrus"
)

// IsDateFormat 判断时间格式是否为 "20060102" 或 "2006-01-02"
func IsDateFormat(date string) (match bool, format string) {
	field1, err := regexp.MatchString("(([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]|[0-9][1-9][0-9]{2}|[1-9][0-9]{3})"+
		"(((0[13578]|1[02])(0[1-9]|[12][0-9]|3[01]))|"+
		"((0[469]|11)(0[1-9]|[12][0-9]|30))|(02(0[1-9]|[1][0-9]|2[0-8]))))|"+
		"((([0-9]{2})(0[48]|[2468][048]|[13579][26])|"+
		"((0[48]|[2468][048]|[3579][26])00))0229)$ ", date)
	if err != nil {
		log.WithFields(log.Fields{"func": "IsDateFormate", "input": date}).Error("date match failed!", err.Error())
	}
	if field1 {
		return true, consts.UTCDay
	}

	matched, err := regexp.MatchString("(([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]|[0-9][1-9][0-9]{2}|[1-9][0-9]{3})"+
		"-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|"+
		"((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8]))))|"+
		"((([0-9]{2})(0[48]|[2468][048]|[13579][26])|"+
		"((0[48]|[2468][048]|[3579][26])00))-02-29)$ ", date)
	if err != nil {
		log.WithFields(log.Fields{"func": "IsDateFormate", "input": date}).Error("date match failed!", err.Error())
	}
	if matched {
		return true, consts.UTCDatetimeDay
	}

	return false, ""
}

// Timestamp2CurrentTime 转换时间戳为当前时间，并判断是不是今天
// 对大于当前时间的情况，均转换为当前时间
func Timestamp2CurrentTime(timestamp string) (match bool, curTimestamp string) {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		log.Fatal(err)
		return true, strconv.Itoa(int(time.Now().UTC().Unix()))
	}

	tm := time.Unix(i, 0).Format(consts.UTCDay)
	now := time.Now().UTC()
	if tm != now.Format(consts.UTCDay) && i < now.Unix() {
		return false, timestamp
	}

	return true, strconv.FormatInt(now.Unix(), 10)
}

// Timestamp2CurrentZeroTime 转换时间戳为当天0时0分0秒
func Timestamp2CurrentZeroTime(timestamp string) (curZeroTimestamp string, err error) {
	tm, err := ParserTimestamp(timestamp)
	if err != nil {
		return "", err
	}

	currentDateZeroTime := tm.Format(consts.UTCDatetimeOnlyDay)
	currentDateZero, err := time.Parse(consts.UTCDatetimeOnlyDay, currentDateZeroTime)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(currentDateZero.Unix(), 10), nil
}

// Timestamp2CurrentMaxTime 转换时间戳为当天23时59分59秒
func Timestamp2CurrentMaxTime(timestamp string) (curMaxTimestamp string, err error) {
	tm, err := Timestamp2CurrentZeroTime(timestamp)
	if err != nil {
		return "", err
	}

	return pkg.StringAdd(tm, "86399"), nil
}

// ParserTimestamp 转换时间戳为 time 格式
func ParserTimestamp(timestamp string) (t time.Time, err error) {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		log.Fatal(err)
		return
	}
	t = time.Unix(i, 0).UTC()
	return t, nil
}

func DateMode2TimeFormat(coin, mode, format string) (startTime, endTime string) {
	aTime, bTime := DateMode2Timestamp(coin, mode)
	start, err := ParserTimestamp(aTime)
	if err != nil {
		log.WithFields(log.Fields{"coin": coin, "func": "DateMode2UTCDay",
			"startTime": startTime, "endTime": endTime}).Error(err.Error())
		return
	}
	end, err := ParserTimestamp(bTime)
	if err != nil {
		log.WithFields(log.Fields{"coin": coin, "func": "DateMode2UTCDay",
			"startTime": startTime, "endTime": endTime}).Error(err.Error())
		return
	}
	startTime = start.Format(format)
	endTime = end.Format(format)
	return startTime, endTime
}

// DateMode2Timestamp
func DateMode2Timestamp(coin, mode string) (startTime, endTime string) {
	switch mode {
	case "hour12", "half_day":
		startTime = strconv.Itoa(int(time.Now().Add(-12 * time.Hour).UTC().Unix()))
		endTime = GetTimestampUnix(0, 0, 0)
	case "day":
		startTime = GetTimestampUnix(0, 0, -1)
		endTime = GetTimestampUnix(0, 0, 0)
	case "day2":
		startTime = GetTimestampUnix(0, 0, -2)
		endTime = GetTimestampUnix(0, 0, 0)
	case "day3":
		startTime = GetTimestampUnix(0, 0, -3)
		endTime = GetTimestampUnix(0, 0, 0)
	case "week":
		startTime = GetTimestampUnix(0, 0, -7)
		endTime = GetTimestampUnix(0, 0, 0)
	case "day12":
		startTime = GetTimestampUnix(0, 0, -12)
		endTime = GetTimestampUnix(0, 0, 0)
	case "month":
		startTime = GetTimestampUnix(0, -1, 0)
		endTime = GetTimestampUnix(0, 0, 0)
	case "month3":
		startTime = GetTimestampUnix(0, -3, 0)
		endTime = GetTimestampUnix(0, 0, 0)
	case "half_year":
		startTime = GetTimestampUnix(0, -6, 0)
		endTime = GetTimestampUnix(0, 0, 0)
	case "year":
		startTime = GetTimestampUnix(-1, 0, 0)
		endTime = GetTimestampUnix(0, 0, 0)
	case "year3":
		startTime = GetTimestampUnix(-3, 0, 0)
		endTime = GetTimestampUnix(0, 0, 0)
	case "year6":
		startTime = GetTimestampUnix(-6, 0, 0)
		endTime = GetTimestampUnix(0, 0, 0)
	case "all":
		var tm time.Time
		switch coin {
		case consts.CoinTypeBitcoin:
			tm, _ = time.Parse(consts.UTCDay, "20090103")
		case consts.CoinTypeBitcoinCash:
			tm, _ = time.Parse(consts.UTCDay, "20090103")
		case consts.CoinTypeLitecoin:
			tm, _ = time.Parse(consts.UTCDay, "20111007")
		default:
			tm, _ = time.Parse(consts.UTCDay, "20090103")
		}
		startTime = strconv.Itoa(int(tm.UTC().Unix()))
		endTime = GetTimestampUnix(0, 0, 0)
	default:
		startTime = GetTimestampUnix(0, 0, -3)
		endTime = GetTimestampUnix(0, 0, 0)
	}

	return startTime, endTime
}

func GetTimestampUnix(x, y, z int) string {
	return strconv.Itoa(int(time.Now().AddDate(x, y, z).UTC().Unix()))
}
