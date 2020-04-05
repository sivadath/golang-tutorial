package rules


/*var a = struct {
	str string
	I 	int
}{"Non exported string",1}
*/

type st struct {
	str string
	I 	int
}

var a  = st{"Non exported string",1}

func GetA() interface{} {
	return a
}