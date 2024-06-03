package main

import "fmt"

type dog struct{}

func (d dog) getSound() string {
	return "Waw- waw"
}

type cat struct{}

func (c cat) getSound() string {
	return "Miu-miu"
}

func main() {
	dog := dog{}
	whatSoundDoesADogMake(dog)
	cat := cat{}
	whatSoundDoesACatMake(cat)

}

func whatSoundDoesADogMake(d dog) {
	fmt.Printf("Это животное говорит \"%s\"", d.getSound())
}

func whatSoundDoesACatMake(c cat) {
	fmt.Printf("Это животное говорит \"%s\"", c.getSound())
}
