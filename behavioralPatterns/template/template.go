package template

import "strings"

/*
TEMPLATE:
- Usefule when writing libraries and frameworks.
- The idea is to provide a user some way to execute code within an algorithm.
- Lets the user write a part of an algorithm while the rest is executed by the abstraction.

OBJECTIVE:
- All about reusability and giving responsibilities to the user.
- Defer a part of an algorithm of the library to the user.
- Improve reusability by abstracting the parts of the code that are not common between executions.

EXAMPLE: a simple algorithm with a deferred step.
- Write an algorithm that is composed of three steps, each of them returns a message.
- First and third stepts are controlled by the template and just the second step is deferred to the user.

ACCEPTANCE CRITERIA:
- Eache step in the algorithm must return a string
- The first step is a method called first() and returns the string hello.
- The third step is a method called third() and returns the string template.
- The second step is whatever string the user wants to return but it's defined by the MessageRetriever interface that has a Message() string method
- The algorithm is executed sequentially by a method called ExecuteAlgorithm and returns the strings returned by each step joined in a single string by a space.
*/

type MessageRetriever interface {
	Message() string
}

type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type TemplateImpl struct {
}

func (t *TemplateImpl) first() string {
	return "hello"
}
func (t *TemplateImpl) third() string {
	return "template"
}
func (t *TemplateImpl) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, "")
}

// --------------------------------------------
type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}

// --------------------------------------------

type TemplateAdapter struct {
	MyFunc func() string
}

func (a *TemplateAdapter) Message() string {
	if a.MyFunc != nil {
		return a.MyFunc()
	}
	return ""
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &TemplateAdapter{MyFunc: f}
}
