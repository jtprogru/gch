package fib

func Recursive(n uint) uint {
	if n <= 2 {
		return 1
	}
	return Recursive(n-1) + Recursive(n-2)
}

func Iterative(position uint) uint {
	if position <= 2 {
		return 1
	}

	slc := make([]uint, position)
	slc[0] = 1
	slc[1] = 1

	var result, i uint
	for i = 2; i < position; i++ {
		result = slc[i-1] + slc[i-2]
		slc[i] = result
	}

	return result
}
