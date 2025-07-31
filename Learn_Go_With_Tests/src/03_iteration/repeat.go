package iteration

import "strings"

const defaultRepeatCount = 5

func Repeat(character string, repeatTimes int) string {
	if repeatTimes <= 0 {
		repeatTimes = defaultRepeatCount
	}

	var repeated strings.Builder
	for i := 0; i < repeatTimes; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
