package main

import "fmt"

/*
COMMAND:
- When you need to connect types that are really unrelated.

OBJECTIVE:
- We are trying to encapsulate some sort of action or information in a light package that must be processed somewhere else.
- Put some information into a box. Just the receiver will open the box and know its contents.
- Delegate some action somewhere else.

EXAMPLE: Simple queue.
- We will put some information into a command implementer and we will have a queue.
- We will create many instances of a type implementing a Command pattern and we will pass them to a queue that will store the commands until three of them are in the queue, at which time it will process them.


ACCEPTANCE CRITERIA:
- Command should reflect somehow the creation of a box that can accept urelated types and the execution of the Command itself.
- We need a constructor of console printing commands. When we using this constructor with a string, it will return a command that will print it. In this case, the handler inside is the command that acts as a bos and as a handler.
- We need a data structure that stores incoming commands in a queue and prints them once the queue reaches the length of three.
*/

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	Message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.Message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command.")
	return &ConsoleOutput{
		Message: s,
	}
}

type CommandeQueue struct {
	Queue []Command
}

func (p *CommandeQueue) AddCommand(c Command) {
	p.Queue = append(p.Queue, c)

	if len(p.Queue) == 3 {
		for _, command := range p.Queue {
			command.Execute()
		}
	}
	p.Queue = make([]Command, 3)
}

func main() {
	queue := CommandeQueue{}

	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))
	queue.AddCommand(CreateCommand("Forth message"))
	queue.AddCommand(CreateCommand("Fifth message"))
}
