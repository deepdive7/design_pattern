package behavioral

import "fmt"

/**
 * 策略接口
 */
type Strategier interface {
	Compute(num1, num2 int) int
}

/**
 * 加法
 */
type Addition struct{}

func (p Addition) Compute(num1, num2 int) int {
	return num1 + num2
}

/**
 * 减法
 */
type Subtraction struct{}

func (p Subtraction) Compute(num1, num2 int) int {
	return num1 - num2
}

type Multiplication struct{}

func (p Multiplication) Compute(num1, num2 int) int {
	return num1 * num2
}

type Division struct{}

func (p Division) Compute(num1, num2 int) int {
	defer func() {
		if f := recover(); f != nil {
			fmt.Println(f)
			return
		}
	}()

	if num2 == 0 {
		panic("num2 must not be 0!")
	}

	return num1 / num2
}
func NewStrategy(t string) (res Strategier) {
	switch t {
	case "s": // 减法
		res = Subtraction{}
	case "m": // 乘法
		res = Multiplication{}
	case "d": // 除法
		res = Division{}
	case "a": // 加法
		fallthrough
	default:
		res = Addition{}
	}

	return
}

type Computer struct {
	Num1, Num2 int
	strate     Strategier
}

func (p *Computer) SetStrategy(strate Strategier) {
	p.strate = strate
}

func (p Computer) Do() int {
	defer func() {
		if f := recover(); f != nil {
			fmt.Println(f)
		}
	}()

	if p.strate == nil {
		panic("Strategier is null")
	}

	return p.strate.Compute(p.Num1, p.Num2)
}
