package regs

import (
	"regexp"
	"strconv"
)

var IsStuNum = func(i string) bool { return regexp.MustCompile(`^\d{13}$`).MatchString(i) }
var IsEmail = func(i string) bool {
	return egexp.MustCompile(`^\w+(?:[-+.]\w+)*@\w+(?:[-.]\w+)*\.\w+(?:[-.]\w+)*$`).MatchString(i)
}
var IsIdNumber = func(i string) bool { return regexp.MustCompile(`^(:?\d{14}|\d{17,17})[\dxX]$`).MatchString(i) }
var IsPhone = func(i string) bool { return regexp.MustCompile(`^1\d{10}$`).MatchString(i) }
var IsQQ = func(i string) bool { return regexp.MustCompile(`^\d{5,10}$`).MatchString(i) }

func IsMale(id string) bool {
	if len(id) == 15 {
		sig := id[len(id)-1 : len(id)]
		if sig == "X" {
			return true
		} else {
			n, _ := strconv.Atoi(sig)
			if n%2 == 0 {
				return false
			} else {
				return true
			}
		}

	} else {
		sig := id[len(id)-2 : len(id)-1]

		n, _ := strconv.Atoi(sig)
		if n%2 == 0 {
			return false
		} else {
			return true
		}

	}
}
