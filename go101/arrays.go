package main

import "fmt"

func main() {

	var tempArr [5]int
	fmt.Println("temp:", tempArr)

	tempArr[1] = 100
	fmt.Println("set:", tempArr)
	fmt.Println("get [1]. element :", tempArr[1])
	fmt.Println("get [4]. element:", tempArr[4])

	fmt.Println("len:", len(tempArr))

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	array := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	array1 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	//Before
	fmt.Println("i1a:", array[0])
	fmt.Println("i1b:", array1[0])

	arrayOp(array, &array1) // Pass reference B and non reference A as array

	//After
	fmt.Println("i3a:", array[0])
	fmt.Println("i3b:", array1[0])

	// B result will change due to passing reference (Pass by reference)
	// A result will be same after arrayOp scope finished(Pass by value)

}

func arrayOp(a [8]int, b *[8]int) {
	a[0] = 2
	b[0] = 3
	fmt.Println("i2a:", a[0])
	fmt.Println("i2b:", b[0])
}
