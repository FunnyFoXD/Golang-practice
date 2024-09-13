package main

import (
	"fmt"
	"calc"
	"github.com/bojanz/currency"
)

func main() {
    sum := calc.AddInts(1, 2, 3, 4, 5)

	//Перевод в копейки
	total, err := currency.NewAmountFromInt64(int64(sum * 100), "RUB")
	if err != nil {
		panic(err)
	}

	//Справа будет выводится ru
	locale := currency.NewLocale("ru")
	formatter := currency.NewFormatter(locale)

	fmt.Println("Балданс:", formatter.Format(total))
}