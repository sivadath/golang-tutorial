package prime

//IsPrime validates if given number is prime.
func IsPrime(num int) bool {
	upperLimit  := num
	for i:=2; i < upperLimit; i ++ {
		if num %i == 0 {
			return false
		}else {
			upperLimit = num/i
		}
	}
	return true
}


type strct struct {
	I int
	i int
}

var I = 10
var i = 20
