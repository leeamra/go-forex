package goforex


	// int -> float (0.7675814602)
	x := int64(7675814602)
	y := float64(x) / 10000000000
	fmt.Printf("%T %+v\n", x, x)
	fmt.Printf("%T %+v\n", y, y)

	// float -> int (0.7675814602)
	x2 := float64(0.7675814602)
	y2 := x2 * 10000000000    // ???
	fmt.Printf("%T %+v\n", x2, x2)
	fmt.Printf("%T %+v\n", y2, y2)
