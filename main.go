package main

import (
	"design-patterns/creationalpatterns"
	"design-patterns/structuralpatterns"
	"github.com/fatih/color"
)

var d = color.New(color.FgHiYellow, color.Bold)

func main() {
	d.Println("############################################################ Creational Design Pattern ############################################################")
	creationalpatterns.AbstractFactory()
	creationalpatterns.FactoryMethod()
	creationalpatterns.Prototype()
	creationalpatterns.Singleton()
	creationalpatterns.Builder()

	d.Println("############################################################ Structural Design Pattern ############################################################")
	structuralpatterns.Adapter()
	structuralpatterns.Bridge()
	structuralpatterns.Composite()
	structuralpatterns.Decorator()
	structuralpatterns.Facade()
	structuralpatterns.Flyweight()
	structuralpatterns.Proxy()
}
