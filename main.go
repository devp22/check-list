package main

import (
	"fmt"
)

func main() {
	var taskItems = []string{"Read Book","Clean House","Get Groceries"}

	fmt.Println("################ Welcome to Checklist App ################")
	fmt.Println()
	printTasks(taskItems)

}

func printTasks(taskItems []string){

	for index, task := range taskItems {
		fmt.Println(index ,". ",task)
	}
}