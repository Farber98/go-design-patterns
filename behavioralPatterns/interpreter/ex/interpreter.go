package main

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

type Interpreter interface {
	Read() int
}

type Value int

func (v *Value) Read() int {
	return int(*v)
}

type OperationSum struct {
	Left  Interpreter
	Right Interpreter
}

func (s *OperationSum) Read() int {
	return s.Left.Read() + s.Right.Read()
}

type OperationSub struct {
	Left  Interpreter
	Right Interpreter
}

func (s *OperationSub) Read() int {
	return s.Left.Read() - s.Right.Read()
}

type OperationMul struct {
	Left  Interpreter
	Right Interpreter
}

func (s *OperationMul) Read() int {
	return s.Left.Read() * s.Right.Read()
}

type OperationDiv struct {
	Left  Interpreter
	Right Interpreter
}

func (s *OperationDiv) Read() int {
	return s.Left.Read() / s.Right.Read()
}

func OperatorFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &OperationSum{left, right}
	case SUB:
		return &OperationSub{left, right}
	case MUL:
		return &OperationMul{left, right}
	case DIV:
		return &OperationDiv{left, right}
	}
	return nil
}

type PolishNotationStack []Interpreter

func (p *PolishNotationStack) Push(s Interpreter) {
	*p = append(*p, s)

}
func (p *PolishNotationStack) Pop() Interpreter {
	lenght := len(*p)
	if lenght > 0 {
		temp := (*p)[lenght-1]
		*p = (*p)[:lenght-1]
		return temp
	}
	return nil
}

func main() {
	stack := PolishNotationStack{}
	operators := strings.Split("5 3 sub 8 mul 4 sum 5 div", " ")

	for _, operatorString := range operators {
		if operatorString == SUM || operatorString == SUB || operatorString == DIV || operatorString == MUL {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := OperatorFactory(operatorString, left, right)
			res := Value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				panic(err)
			}

			temp := Value(val)
			stack.Push(&temp)
		}
	}

	println(int(stack.Pop().Read()))
}
