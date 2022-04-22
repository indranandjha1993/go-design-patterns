package structuralpatterns

import (
	"design-patterns/structuralpatterns/adapter"
	"design-patterns/structuralpatterns/bridge"
	"design-patterns/structuralpatterns/composite"
	"design-patterns/structuralpatterns/decorator"
	"design-patterns/structuralpatterns/facade"
	"design-patterns/structuralpatterns/flyweight"
	"design-patterns/structuralpatterns/proxy"
	"github.com/fatih/color"
)

var d = color.New(color.FgRed, color.Bold)

func Adapter() {
	d.Printf("#################### This is the %s output ####################\n", "Adapter")
	adapter.Run()
}

func Bridge() {
	d.Printf("#################### This is the %s output ####################\n", "Bridge")
	bridge.Run()
}

func Composite() {
	d.Printf("#################### This is the %s output ####################\n", "Composite")
	composite.Run()
}

func Decorator() {
	d.Printf("#################### This is the %s output ####################\n", "Decorator")
	decorator.Run()
}

func Facade() {
	d.Printf("#################### This is the %s output ####################\n", "Facade")
	facade.Run()
}

func Flyweight() {
	d.Printf("#################### This is the %s output ####################\n", "Flyweight")
	flyweight.Run()
}

func Proxy() {
	d.Printf("#################### This is the %s output ####################\n", "Proxy")
	proxy.Run()
}
