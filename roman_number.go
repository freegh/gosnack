package main

import (
	"fmt"
)

func main(){
	for i:=1;i<=4000;i++ {
		fmt.Println(i,":",roman_number(i))
	}
}

func roman_number(number int) string {
	var romanNumber string
	for number>0{
		switch {
		case number >= 1000:
			number-=1000
			romanNumber+="M"
		case number >= 900:
			number-=900
			romanNumber+="CM"
		case number >= 500:
			number-=500
			romanNumber+="D"
		case number >= 400:
			number-=400
			romanNumber+="CD"
		case number >= 100:
			number-=100
			romanNumber+="C"
		case number >= 90:
			number-=90
			romanNumber+="XC"
		case number >= 50:
			number-=50
			romanNumber+="L"
		case number >= 40:
			number-=40
			romanNumber+="XL"
		case number >= 10:
			number-=10
			romanNumber+="X"
		case number >= 9:
			number-=9
			romanNumber+="IX"
		case number >= 5:
			number-=5
			romanNumber+="V"
		case number >= 4:
			number-=4
			romanNumber+="IV"
		case number >= 1:
			number-=1
			romanNumber+="I"
		default:
			fmt.Println("ERROR")
		}
	}

	return romanNumber
}