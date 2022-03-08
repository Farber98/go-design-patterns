package interpreter

import (
	"strconv"
	"strings"
)

/*
INTERPRETER:
- Widely used to solve business cases where it's useful to have language to perform common operations.

OBJECTIVE:
- Provide a syntax for very common operations in some scope.
- Have an intermediate language to translate actions between two systems.
- Ease the use of some operations in an easier-to-use syntax.


EXAMPLE: polish notation calculator.

ACCEPTANCE CRITERIA:
- Create a language that allows making common arithm operations (sum, sub, mul, div).
- It muste be done using reverse polish notation
- The user must be able to write as many operations in a row as they want.
- The operations must be performed from left to right.

*/
const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

type PolishNotationStack []int

func (p *PolishNotationStack) Push(s int) {
	*p = append(*p, s)

}
func (p *PolishNotationStack) Pop() int {
	lenght := len(*p)
	if lenght > 0 {
		temp := (*p)[lenght-1]
		*p = (*p)[:lenght-1]
		return temp
	}
	return 0
}

func Calculate(o string) (int, error) {
	stack := PolishNotationStack{}
	operators := strings.Split(o, " ")
	for _, operatorString := range operators {
		if IsOperator(operatorString) {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := getOperationFunc(operatorString)
			res := mathFunc(left, right)
			stack.Push(res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				return 0, err
			}
			stack.Push(val)
		}
	}
	return int(stack.Pop()), nil
}

func IsOperator(o string) bool {
	if o == SUM || o == SUB || o == MUL || o == DIV {
		return true
	}
	return false
}

func getOperationFunc(o string) func(a, b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	}
	return nil
}
