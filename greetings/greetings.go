package greetings

import "fmt"

func Hello(name string) string {
	//Return a greeting that emnbeds the name in a message.
	message := fmt.Sprintf("hi, %v. Welcome!", name)
	return message

}
