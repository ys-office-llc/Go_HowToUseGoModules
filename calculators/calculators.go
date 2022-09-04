package calculators

// 演算インタフェース型
type Calculator interface {

	// 関数の定義
	Calculate(a int, b int) int
}

// 足し算型
type Add struct {
	// フィールドは持たない
}

// Add型にCalculatorインタフェースのCalculate関数を実装
func (x Add) Calculate(a int, b int) int {

	// 足し算
	return a + b
}

// 引き算型
type Sub struct {
	// フィールドは持たない
}

// Sub 型にCalculatorインタフェースのCalculate関数を実装
func (x Sub) Calculate(a int, b int) int {

	// 引き算
	return a - b
}
