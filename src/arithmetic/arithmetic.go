package arithmetic

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	if (y == 0)	{
		if (x == 0) {
			return 1
		} else if (x < 0){
			return -0x80000000
		} else {
			return  0x7FFFFFFF
		}
	}
	return x / y
}