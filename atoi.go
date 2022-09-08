package piscine

func Atoi(s string) int {
	rep := 0
	coef_mul := 1
	n := 1
	atoi := []rune{}
	for _, character := range s {
		if len(s) == 1 {
			if character == '-' || character == '+' {
				return 0
			} else {
				return int(character)
			}
		} else if character == '-' && s[1] != '-' {
			coef_mul = -1
		} else if character == '+' && s[1] != '+' {
			coef_mul = 1
		} else if character == '+' && s[1] == '+' || character == '-' && s[1] == '-' {
			return 0
		} else if character != '0' && character != '1' && character != '2' && character != '3' && character != '4' && character != '5' && character != '6' && character != '7' && character != '8' && character != '9' {
			return 0
		} else {
			atoi = append(atoi, rune(character))
		}
	}
	for i := len(atoi) - 1; len(atoi) != 0; i-- {
		rep += int(atoi[i]-48) * n
		atoi = atoi[:len(atoi)-1]
		n *= 10
	}
	return (rep * coef_mul)
}
