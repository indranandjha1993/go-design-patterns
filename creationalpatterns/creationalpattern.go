package creationalpatterns

import (
	"design-patterns/creationalpatterns/abstractfactory"
	"design-patterns/creationalpatterns/builder"
	"design-patterns/creationalpatterns/prototype"
	"design-patterns/creationalpatterns/singleton"
	"github.com/fatih/color"
)

var d = color.New(color.FgRed, color.Bold)

func AbstractFactory() {
	d.Printf("#################### This is the %s output ####################\n", "Abstract Factory")
	abstractfactory.Run()
}
func FactoryMethod() {
	d.Printf("#################### This is the %s output ####################\n", "Factory Method")
	abstractfactory.Run()
}
func Prototype() {
	d.Printf("#################### This is the %s output ####################\n", "Prototype")
	prototype.Run()
}
func Singleton() {
	d.Printf("#################### This is the %s output ####################\n", "Singleton")
	singleton.Run()
}
func Builder() {
	d.Printf("#################### This is the %s output ####################\n", "Builder")
	builder.Run()
}
