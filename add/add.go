package main

import (
    "fmt"

    "example.com/calculators"
)

func main() {

	// Calculatorインタフェースを実装した型の変数を宣言
	var add calculators.Add // 足し算

	// Calculatorインタフェース型の変数を宣言
	var cal calculators.Calculator

	// Add型の値を代入
	cal = add

	// インタフェース経由でメソッドを呼び出す
	fmt.Println("和:", cal.Calculate(1, 2))
}
