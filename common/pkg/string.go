package pkg

import (
	"fmt"
	"math"
	"strconv"
)

const (
	// NULL null
	NULL = "NULL"
)

// StringMultiplication 字符串相乘
func StringMultiplication(a, b string) string {
	floatA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "0"
	}
	floatB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "0"
	}
	return strconv.FormatFloat(floatA*floatB, 'f', -1, 64)
}

// StringAdd 字符串相加
func StringAdd(a, b string) string {
	if a == NULL {
		a = "0"
	}
	if b == NULL {
		b = "0"
	}
	floatA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "0"
	}
	floatB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "0"
	}
	return strconv.FormatFloat(floatA+floatB, 'f', -1, 64)
}

// StringDivision 字符串相除
func StringDivision(a, b string) string {
	floatA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "0"
	}
	floatB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "0"
	}
	return fmt.Sprintf("%.8f", floatA/floatB)
}

// StringSub 字符串相减
func StringSub(a, b string) string {
	floatA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "0"
	}
	floatB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "0"
	}
	return strconv.FormatFloat(floatA-floatB, 'f', -1, 64)
}

// StringSubAbs 字符串相减
func StringSubAbs(a, b string) string {
	floatA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return "0"
	}
	floatB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return "0"
	}
	return strconv.FormatFloat(math.Abs(floatA-floatB), 'f', -1, 64)
}

// StringFormat 对于数字字符串保留 digits 位小数
func StringFormat(str, digits string) string {
	floatA, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return "0"
	}
	_, err = strconv.Atoi(digits)
	if err != nil {
		return "0"
	}
	return fmt.Sprintf("%."+digits+"f", floatA)
}

// StringFormatCeil 对于数字字符串向上取整
func StringFormatCeil(str string) string {
	floatA, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return "0"
	}
	return fmt.Sprintf("%.0f", math.Ceil(floatA))
}

// StringFormatFloor 对于数字字符串向下取整
func StringFormatFloor(str string) string {
	floatA, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return "0"
	}
	return fmt.Sprintf("%.0f", math.Floor(floatA))
}

// StringFormatRound 对于数字字符串四舍五入取整
func StringFormatRound(str string) string {
	floatA, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return "0"
	}
	return fmt.Sprintf("%.0f", math.Floor(floatA+0.5))
}
