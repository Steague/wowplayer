package main

import "fmt"
import ps "github.com/mitchellh/go-ps"

func main() {
	test, err := ps.Processes()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %v", err.Error()))
	}

	for i := range test {
		if test[i].Executable() != "notepad.exe" {
			continue
		}
		fmt.Println(test[i].Pid(), test[i].Executable())
	}
}
