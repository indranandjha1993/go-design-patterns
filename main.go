package main

import (
	"design-patterns/creationalpatterns"
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
}
