package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// doAFlipp create file, open and close it
func doAFlipp(fileName string) {

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File created: %s \n", fileName)

	file, err = os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File opened: %s \n", fileName)

	defer func() {
		err = file.Close()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("File closed: %s \n", fileName)
	}()

}

// ErrorWithTime is error with timestamp
type ErrorWithTime struct {
	text string
	time string
}

// New creates error
func New(text string) error {
	return &ErrorWithTime{
		text: text,
		time: time.Now().String(),
	}
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s\ntime:%s", e.text, e.time)
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Got a panic. Recovered", v)
		}
		doAFlipp("fileTwo.txt")
	}()

	doAFlipp("fileOne.txt")

	var err error

	err = New("my error")
	fmt.Println(err)

	var a int
	fmt.Println(1 / a)

}
