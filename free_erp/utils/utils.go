package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func StrToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return int64(0)
	}
	return num
}

func StrToInt(str string) int {
	return int(StrToInt64(str))
}

func StrToFloat64(str string) float64 {
	if num, err := strconv.ParseFloat(str, 64); err != nil {
		return 0.0
	} else {
		return num
	}
}

/*
* 对sql语句转义，防止SQL注入攻击
 */
func EscapeStringBackslash(v string) string {
	pos := 0
	buf := make([]byte, len(v)*2)

	for i := 0; i < len(v); i++ {
		c := v[i]
		switch c {
		case '\x00':
			buf[pos] = '\\'
			buf[pos+1] = '0'
			pos += 2
		case '\n':
			buf[pos] = '\\'
			buf[pos+1] = 'n'
			pos += 2
		case '\r':
			buf[pos] = '\\'
			buf[pos+1] = 'r'
			pos += 2
		case '\x1a':
			buf[pos] = '\\'
			buf[pos+1] = 'Z'
			pos += 2
		case '\'':
			buf[pos] = '\\'
			buf[pos+1] = '\''
			pos += 2
		case '"':
			buf[pos] = '\\'
			buf[pos+1] = '"'
			pos += 2
		case '\\':
			buf[pos] = '\\'
			buf[pos+1] = '\\'
			pos += 2
		default:
			buf[pos] = c
			pos += 1
		}
	}

	return string(buf[:pos])
}

/*
 * 生成sql语句，已包含字符串转义，防止SQL注入攻击
 * 使用方法同 fmt.Sprintf(format, a...)
 */
func GetEscapeSqlClause(format string, a ...interface{}) string {
	if len(a) <= 0 {
		return fmt.Sprintf(format, a...)
	}

	args := make([]interface{}, 0, len(a))
	for _, arg := range a {
		switch arg.(type) {
		case string:
			newArg := EscapeStringBackslash(arg.(string))
			args = append(args, newArg)
		default:
			args = append(args, arg)
		}
	}

	return fmt.Sprintf(format, args...)
}

func GetMd5Value(cipherText string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(cipherText))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
