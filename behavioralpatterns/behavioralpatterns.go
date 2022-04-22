package behavioralpatterns

import (
	"design-patterns/behavioralpatterns/chainofresponsibility"
	"design-patterns/behavioralpatterns/command"
	"design-patterns/behavioralpatterns/iterator"
	"design-patterns/behavioralpatterns/mediator"
	"design-patterns/behavioralpatterns/memento"
	"design-patterns/behavioralpatterns/observer"
	"design-patterns/behavioralpatterns/state"
	"design-patterns/behavioralpatterns/strategy"
	"design-patterns/behavioralpatterns/templatemethod"
	"design-patterns/behavioralpatterns/visitor"
	"github.com/fatih/color"
)

var d = color.New(color.FgRed, color.Bold)

func ChainOfResponsibility() {
	d.Printf("#################### This is the %s output ####################\n", "Chain Of Responsibility")
	chainofresponsibility.Run()
}

func Command() {
	d.Printf("#################### This is the %s output ####################\n", "Command")
	command.Run()
}

func Iterator() {
	d.Printf("#################### This is the %s output ####################\n", "Iterator")
	iterator.Run()
}

func Mediator() {
	d.Printf("#################### This is the %s output ####################\n", "Mediator")
	mediator.Run()
}

func Memento() {
	d.Printf("#################### This is the %s output ####################\n", "Memento")
	memento.Run()
}

func Observer() {
	d.Printf("#################### This is the %s output ####################\n", "Observer")
	observer.Run()
}

func State() {
	d.Printf("#################### This is the %s output ####################\n", "State")
	state.Run()
}

func Strategy() {
	d.Printf("#################### This is the %s output ####################\n", "Strategy")
	strategy.Run()
}

func TemplateMethod() {
	d.Printf("#################### This is the %s output ####################\n", "Template method")
	templatemethod.Run()
}

func Visitor() {
	d.Printf("#################### This is the %s output ####################\n", "Visitor")
	visitor.Run()
}
