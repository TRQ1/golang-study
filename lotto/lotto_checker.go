package main


func IsVaildLottery(numbers ...int) bool {
	if !IsVaildLength(numbers...) {
		return false
	}

	BeForeNumber := 0
	for _, n := range numbers {
		if !IsVaildRange(n) {
			return false
		}
		if BeForeNumber >= n {
			return false
		}
		BeForeNumber = n
	}
	return true
}

func IsVaildRange(number int) bool {
	return 1 <= number && number <=45
}

func IsVaildLength(numbers ...int) bool {
	return len(numbers) == 6
}