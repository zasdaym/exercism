package airportrobot

import "fmt"

type Greeter interface {
	Greet(name string) string
	LanguageName() string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

type Italian struct{}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

func (i Italian) LanguageName() string {
	return "Italian"
}

type Portuguese struct{}

func (i Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}
