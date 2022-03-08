package visitor

import "io"

/*
VISITOR:
- Tries to separate the logic needed to work with a specific object outside the object itself.
- Many different visitors do some things to specific tymes.
- Delegates some logic of an object's type to an external type called the visitor that will visit our object to perform operations on it.

OBJECTIVE:
- Separate the algoithm of some type from its implementation within some other type.
- Improve flexibility of some types by using them with little or no logic at all so all new functionality can be added without altering the object structure

EXAMPLE: log appender
- Visitor that appends different information to the types it visits.

ACCEPTANCE CRITERIA:
- Must have two roles: visitor and visitable.
- Visitor is the type that will act within a visitable type.
- We need two message loggers: MEssage A and MessageB that will apend A or B to the message.
* Need a visitor able to modify the message to be printed.
*/
type MessageA struct {
	Msg    string
	Output io.Writer
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

type Visitable interface {
	Accept(Visitor)
}

type MessageVisitor struct{}
