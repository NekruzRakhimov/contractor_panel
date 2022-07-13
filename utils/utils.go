package utils

import "fmt"

func ReverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func FloatToMoneyFormat(numIn float64) (numOut string) {
	counter := 0
	temp := fmt.Sprintf("%.2f", numIn)
	//fmt.Println(temp)
	temp = ReverseString(temp)

	for i := 0; i < len(temp); i++ {
		if i < 3 {
			numOut = string(temp[i]) + numOut
			continue
		}

		if counter == 3 {
			numOut = " " + numOut
			counter = 0
		}

		numOut = string(temp[i]) + numOut
		counter++

		//fmt.Println("i: ", i)
		//fmt.Println("	temp[i]: ", string(temp[i]))
		//fmt.Println("	counter: ", counter)
		//fmt.Println("	numOut: ", numOut)
	}

	return numOut
}
