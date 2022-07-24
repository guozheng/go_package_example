package main

import (
	"fmt"

	print "github.com/guozheng/go_package_example/formatter"
	"github.com/guozheng/go_package_example/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
