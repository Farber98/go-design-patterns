package visitor

import "testing"

type TestHelper struct {
	Received string
}

func (t *TestHelper) Write(p []byte) (int, error) {
	t.Received = string(p)
	return len(p), nil
}

func Test_Overall(t *testing.T) {
	/* 	testHelper := &TestHelper{}
	   	visitor := &MessageVisitor{}

	   	t.Run("MessageA test", func(t *testing.T) {
	   		msg := MessageA{"Hello World", testHelper}
	   	}) */
}
