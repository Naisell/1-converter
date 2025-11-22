package main

import "fmt"

const usdToEur = 0.85
const usdToRub = 83.37

func main() {
	eurToRub := usdToRub / usdToEur
	fmt.Printf("%.2f RUB\n", eurToRub)
}
