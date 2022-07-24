package main

import (
	"fmt"
	"log"

	print "github.com/guozheng/go_package_example/formatter"
	"github.com/guozheng/go_package_example/math"
	"github.com/shopspring/decimal"
)

func printDouble(n int) {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}

func calcTotal(price string, percent string) {
	priceDecimal, err := decimal.NewFromString(price)
	if err != nil {
		log.Fatal(err)
	}

	percentDecimal, err := decimal.NewFromString(percent)
	if err != nil {
		log.Fatal(err)
	}
	percentDecimal = percentDecimal.Div(decimal.NewFromInt(100))

	total := priceDecimal.Mul(percentDecimal.Add(decimal.NewFromInt(1))).Round(2)
	fmt.Println(print.Space(80, price, percent, total.StringFixed(2)))
}

func main() {
	// printDouble(2)
	calcTotal("25", "10")
}
