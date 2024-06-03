package utils

func onlyNumbers(s string) string {
	var numbers string
	for _, char := range s {
		if char >= '0' && char <= '9' {
			numbers += string(char)
		}
	}
	return numbers
}

func IsValidCNPJ(cnpj string) bool {
	cnpj = onlyNumbers(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	sum := 0
	for i := 0; i < 12; i++ {
		sum += int(cnpj[i]-'0') * (i%8 + 2)
	}
	digit1 := (sum * 10) % 11
	if digit1 == 10 {
		digit1 = 0
	}

	if int(cnpj[12]-'0') != digit1 {
		return false
	}

	sum = 0
	for i := 0; i < 13; i++ {
		sum += int(cnpj[i]-'0') * ((i + 1) % 8)
	}
	digit2 := (sum * 10) % 11
	if digit2 == 10 {
		digit2 = 0
	}

	if int(cnpj[13]-'0') != digit2 {
		return false
	}

	return true
}
