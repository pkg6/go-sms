package gosms

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var char = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

// RandString 生成随机字符串
func RandString(lenNum int) string {
	str := strings.Builder{}
	length := len(char)
	for i := 0; i < lenNum; i++ {
		l := char[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

// RandLetterString 生成随机字母
func RandLetterString(lenNum int) string {
	str := strings.Builder{}
	length := 52
	for i := 0; i < lenNum; i++ {
		str.WriteString(char[rand.Intn(length)])
	}
	return str.String()
}

// RandNumString 生成随机数字
func RandNumString(lenNum int) string {
	str := strings.Builder{}
	length := 10
	for i := 0; i < lenNum; i++ {
		str.WriteString(char[52+rand.Intn(length)])
	}
	return str.String()
}

// Uniqid 基于以微秒计的当前时间,生成一个唯一的ID
func Uniqid(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}

// Rand 随时数
func Rand(min, max int) int {
	if min > max {
		panic("min: min cannot be greater than max")
	}
	if int31 := 1<<31 - 1; max > int31 {
		panic("max: max can not be greater than " + strconv.Itoa(int31))
	}
	if min == max {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max+1-min) + min
}

// RandomInt random_int()
func RandomInt(min, max int) (int, error) {
	if min > max {
		panic("argument #1 must be less than or equal to argument #2")
	}
	if min == max {
		return min, nil
	}
	nb, err := crand.Int(crand.Reader, big.NewInt(int64(max+1-min)))
	if err != nil {
		return 0, err
	}
	return int(nb.Int64()) + min, nil
}

// RandomBytes random_bytes()
func RandomBytes(length int) ([]byte, error) {
	bs := make([]byte, length)
	_, err := crand.Read(bs)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// Sleep sleep()
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// Usleep usleep()
func Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}

// Time time()
func Time() int64 {
	return time.Now().Unix()
}

// TimeString Time time()
func TimeString() string {
	return strconv.FormatInt(Time(), 10)
}
