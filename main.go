package main

import "fmt"

const usdToEur = 0.85
const usdToRub = 83.37

func main() {
	amount := getAmount()
	from := getCurrency("Из валюты (USD/EUR/RUB): ")
	to := getCurrency("В валюту (USD/EUR/RUB): ")

	result := convert(amount, from, to)
	fmt.Printf("%.2f\n", result)
}

func getAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму для конвертации: ")
		_, err := fmt.Scan(&amount)

		if err != nil {
			fmt.Println("Ошибка: введите число. Попробуйте снова.")
			continue
		}
		if amount < 0 {
			fmt.Println("Ошибка: сумма не может быть отрицательной. Попробуйте снова.")
			continue
		}
		return amount
	}

}

func getCurrency(prompt string) string {
	var currency string
	for {
		fmt.Print(prompt)
		fmt.Scan(&currency)

		if currency == "USD" || currency == "EUR" || currency == "RUB" {
			return currency
		}
		fmt.Println("Ошибка: используйте USD, EUR или RUB. Попробуйте снова.")
	}
}

func convert(amount float64, from string, to string) float64 {

	var inUSD float64

	switch from {
	case "USD":
		inUSD = amount
	case "EUR":
		inUSD = amount / usdToEur
	case "RUB":
		inUSD = amount / usdToRub
	}

	switch to {
	case "USD":
		return inUSD
	case "EUR":
		return inUSD * usdToEur
	case "RUB":
		return inUSD * usdToRub
	}
	return 0
}
