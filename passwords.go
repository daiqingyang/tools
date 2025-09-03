package tools

import (
	"math"
	"math/rand"
)

// 生成随机密码
func GenPassword() (rst string) {
	letters := "abcdefghigklmnopqrstuvwxwzABCDEFGHIGKLMNOPQRSTUVWXWZ0123456789!@$%^-_=+[{}]:,./?"
	for {
		var lower, upper, number, other bool
		var count int
		b := make([]byte, 12)
		for i := range b {
			idx := rand.Intn(len(letters))
			if idx < 26 {
				lower = true
			} else if idx >= 26 && idx < 52 {

				upper = true
			} else if idx >= 52 && idx < 62 {
				number = true
			} else {
				other = true
			}

			b[i] = letters[idx]
		}

		if lower {
			count++
		}
		if upper {
			count++
		}
		if number {
			count++
		}
		if other {
			count++
		}
		//密码至少必须包含大写字母、小写字母、数字和特殊字符（!@$%^-_=+[{}]:,./?）中的三种
		if count >= 3 {
			rst = string(b)
			break
		}
	}
	return
}

// 计算 Shannon 熵
// 熵检测 = 判断字符串随机程度
// 随机性高 → 熵值大。
// 重复性强/模式明显 → 熵值小
// H=−i∑​p(xi​)log2​p(xi​)
// 所以我们通常设置一个 熵阈值（比如 4.0），超过就认为可能是敏感信息
// /ˈʃænən ˈɛntrəpi/
// 👉 可以读作 “SHAN-nən EN-trə-pee”。
func ShannonEntropy(s string) float64 {
	if len(s) == 0 {
		return 0.0
	}

	// 统计每个字符出现次数
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}

	// 计算熵
	var entropy float64
	length := float64(len(s))
	for _, count := range freq {
		p := float64(count) / length
		entropy -= p * math.Log2(p)
	}

	return entropy
}
