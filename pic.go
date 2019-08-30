package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	showPic := [][]uint8{}
	for i := 0; i < dy; i++ {
		row := []uint8{}

		for j := 0; j < dx; j++ {
			row = append(row, uint8(1))
		}

		showPic = append(showPic, row)
	}
	
	return showPic
}

func main()  {
	pic.Show(Pic)
}



// package main

// import (
// 	"golang.org/x/tour/pic"
// )

// func Pic(dx, dy int) [][]uint8 {

// 	showPic := make([][]uint8,dy)
// 	for i := range showPic {
// 		showPic[i] = make([]uint8, dx)
// 	}
	
// 	return showPic
// }

// func main()  {
// 	pic.Show(Pic)
// }