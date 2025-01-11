package main

import (
	"fmt"
)

var taskItems = []string{"Read Book","Clean House","Get Groceries"}
func main() {

	fmt.Println("################ Welcome to Checklist App ################")
	fmt.Println()
	printTasks()

}

func printTasks(){

	for index, task := range taskItems {
		fmt.Println(index ,". ",task)
	}
}