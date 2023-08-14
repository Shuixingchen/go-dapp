package realtime

import (
	"fmt"
	"io"
	netURL "net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Shuixingchen/go-dapp/common/pkg/cryptos"
	socketio "github.com/googollee/go-socket.io"
)

const (
	RequestFieldSign      = "sign"
	RequestFieldEIO       = "EIO"
	RequestFieldTransport = "transport"
)

// CheckSign 检查签名
func CheckSign(s socketio.Conn, appDebug bool, configAppIds map[string]string, configAppKey string, supportCoins []string) bool {
	url := s.URL()

	var params []string
	var timestamp int64

	queryParam := url.Query()
	sortedParam := make([]string, 0, len(queryParam))
	for k := range queryParam {
		if k == RequestFieldEIO || k == RequestFieldTransport {
			continue
		}
		sortedParam = append(sortedParam, k)
	}
	sort.Strings(sortedParam)

	nonce := url.Query().Get("nonce")
	appID := url.Query().Get("app_a")
	appName := url.Query().Get("app_b")
	requestTimestamp := url.Query().Get("timestamp")
	coins := url.Query().Get("coins")
	sign := url.Query().Get(RequestFieldSign)
	if requestTimestamp != "" {
		var err error
		timestamp, err = strconv.ParseInt(requestTimestamp, 10, 64)
		if err != nil {
			s.Emit("warning", "invalid timestamp")
			return false
		}
	}
	for _, values := range sortedParam {
		if values != RequestFieldSign {
			for _, value := range queryParam[values] {
				if value != "" {
					params = append(params, value)
				}
			}
		}
	}

	if coins != "" {
		coinList := strings.Split(coins, ",")
		if len(coinList) > len(supportCoins) {
			s.Emit("warning",
				fmt.Sprintf("please set correct query param 'coins',the length of coins beyond the limit %d",
					len(supportCoins)))
			return false
		}

		for _, v := range coinList {
			for k, coin := range supportCoins {
				if strings.EqualFold(coin, v) {
					break
				}
				if k == len(supportCoins)-1 {
					s.Emit("warning",
						fmt.Sprintf("please set correct query param 'coins',the %s is not right", v))
					return false
				}
			}
		}
	}

	if sign == "" && appDebug {
		return true
	}

	if appID == "" {
		s.Emit("warning", "app_a must not be empty")
		return false
	}

	if appName == "" {
		s.Emit("warning", "app_b must not be empty")
		return false
	}

	if timestamp < (time.Now().UTC().Unix() - 300) {
		s.Emit("warning", "invalid timestamp")
		return false
	}

	if len(nonce) < 10 {
		s.Emit("warning", "the length of nonce must >= 10")
		return false
	}

	// 这里需要按配置去区分，如果是 API 服务，则是 API Ids
	// 如果是提供给浏览器的，则是 浏览器对应的 API
	exceptAppID := configAppIds[appName]
	if exceptAppID != appID {
		s.Emit("warning", "invalid app_a")
		return false
	}

	secretKey := cryptos.SHA256(cryptos.SHA256(cryptos.MD5(appID+"|"+configAppKey) + cryptos.MD5(configAppKey)))
	// 浏览器和 API 加密方式按理上说应该不同
	realSign := cryptos.HMACsha256(strings.Join(params, "|"), secretKey)
	realUnescapeSign, err := netURL.QueryUnescape(realSign)
	if err != nil && err != io.EOF {
		s.Emit("warning", "invalid sign")
		return false
	}

	unescapeSign, err := netURL.QueryUnescape(sign)
	if err != nil && err != io.EOF {
		s.Emit("warning", "invalid sign")
		return false
	}

	if sign != realUnescapeSign && sign != realSign && unescapeSign != realUnescapeSign && unescapeSign != realSign {
		s.Emit("warning", "invalid sign")
		return false
	}

	return true
}
