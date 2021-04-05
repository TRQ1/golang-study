package main


func IsVaildLottery(numbers ...int) bool {
	if !IsVaildLength(numbers...) || !IsVaildRange(numbers...) || !IsAscOrder(numbers...) {
		return false
	}

	return true
}

func IsAscOrder(numbers ...int) bool {
	BeForeNumber := 0
	for _, n := range numbers {
		if BeForeNumber >= n {
			return false
		}
  	 BeForeNumber = n 
   	}
  	return true
}

func IsVaildRange(numbers ...int) bool {
 	for _, n := range numbers {
		 if n < 1 || n > 45 {
			 return false
		 }
    }
  	return true
}

func IsVaildLength(numbers ...int) bool {
	return len(numbers) == 6
}