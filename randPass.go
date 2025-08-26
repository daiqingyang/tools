package tools

import (
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
