package messages

import "fmt"

func Greet(name string) string { // public function
	return fmt.Sprintf("Hello, %v!\n", name)
}

func depart(name string) string { // private function
	return fmt.Sprintf("Bye, %v\n", name)
}
