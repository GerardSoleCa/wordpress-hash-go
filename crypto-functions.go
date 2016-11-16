package wphash

import (
	"crypto/md5"
	"math"
	"math/rand"
	"strings"
	"time"
)

const ITOA64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const ITERATION_COUNT_LOG2 = 8
const POSSIBLE_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func encode64(input string, count int) string {
	output := ""
	i := 0

	for ok := true; ok; ok = (i < count) {
		value := uint(input[i])

		i++

		output += string(ITOA64[value&0x3f])

		if i < count {
			value |= uint(input[i]) << 8
		}

		output += string(ITOA64[(value>>6)&0x3f])

		if i >= count {
			break
		}

		i++

		if i < count {
			value |= uint(input[i]) << 16
		}

		output += string(ITOA64[(value>>12)&0x3f])

		if i >= count {
			break
		}

		i++

		output += string(ITOA64[(value>>18)&0x3f])
	}

	return output
}

func generatePrivateSalt(input string) string {
	output := "$P$"
	output += string(ITOA64[int(math.Min(ITERATION_COUNT_LOG2+5, 30))])
	output += encode64(input, 6)
	return output
}

func cryptPrivate(password string, setting string) string {
	output := "*0"

	if setting[0:2] == output {
		output = "*1"
	}

	if setting[0:3] != "$P$" {
		return output
	}

	countLog2 := uint(strings.Index(ITOA64, string(setting[3])))
	if countLog2 < 7 || countLog2 > 30 {
		return output
	}

	count := 1 << countLog2

	salt := setting[4:12]

	if len(salt) != 8 {
		return output
	}

	hash := hashString(salt + password)
	for count != 0 {
		hash = hashString(hash + password)
		count--
	}

	output = setting[0:12]
	output += encode64(hash, 16)

	return output
}

func hashString(data string) string {
	return hash([]byte(data))
}

func hash(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	return string(hasher.Sum(nil))
}

func sixCharRandom() string {
	text := ""
	i := 0

	for i < 6 {
		text += string(POSSIBLE_CHARS[int(math.Floor(rand.Float64()*float64(len(POSSIBLE_CHARS))))])
		i++
	}

	return text
}
