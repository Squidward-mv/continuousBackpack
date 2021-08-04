package main

import "fmt"

type item struct{
	price float64
	weight float64
}

func main() {

	var n int
	var volume float64
	fmt.Scan(&n)
	fmt.Scan(&volume)

	arrOfItems := make([]item, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arrOfItems[i].price)
		fmt.Scan(&arrOfItems[i].weight)
	}

	for i := 0; i < n; i++{
		for j := 1; j < n; j++ {
			if (arrOfItems[j].price / arrOfItems[j].weight) < (arrOfItems[j-1].price / arrOfItems[j-1].weight) {
				temp := arrOfItems[j-1]
				arrOfItems[j-1] = arrOfItems[j]
				arrOfItems[j] = temp
			}
		}
	}

	var maxPrice float64

	for n > 0 {
		if volume > arrOfItems[n-1].weight {
			volume -= arrOfItems[n-1].weight
			maxPrice += arrOfItems[n-1].price
		}else {
			maxPrice += arrOfItems[n-1].price * (volume / arrOfItems[n-1].weight)
			break
		}

		n -= 1
	}

	fmt.Printf("%3f",maxPrice)
}