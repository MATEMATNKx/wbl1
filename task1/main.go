package main

import "fmt"

type Human struct {
	Name string
}

func (h Human) getName() string {
	return h.Name
}

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Action struct {
	Human
}

func main() {
	human_ := Human{Name: "Vladimir Vseponyatno"}
	action_ := Action{Human: human_}
	fmt.Println(action_.getName())
}
