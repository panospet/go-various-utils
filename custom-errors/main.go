package main

import "fmt"

type ErrCustom struct {
	Message string
}

func (e ErrCustom) Error() string {
	return fmt.Sprintf("this is the message: %s", e.Message)
}

func main() {
	err := Test(true)
	switch err.(type) {
	case ErrCustom:
		fmt.Println("custom error caught!")
	default:
		panic(err)
	}

	err = Test(false)
	switch err.(type) {
	case ErrCustom:
		panic(err)
	default:
		fmt.Println("default error caught!")
	}
}

func Test(giveIt bool) error {
	if giveIt {
		return ErrCustom{Message: "custom error thrown"}
	}
	return fmt.Errorf("another")
}
