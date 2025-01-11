package main

import (
	"fmt"
)

func main() {
	var taskItems = []string{"Read Book","Clean House","Get Groceries"}

	printTasks(taskItems,"################ Welcome to Checklist App ################")
	taskItems = addTask(taskItems,"Go for a run")
	printTasks(taskItems,"######### Updated List #########")
}

func printTasks(taskItems []string, header string){

	fmt.Println(header)
	for index, task := range taskItems {
		fmt.Printf("%d. %s\n",index+1,task)
	}
}

func addTask(taskItems []string, newTask string) ([]string) {
	var updatedTaskItems = append(taskItems, newTask)
	fmt.Println("Task added successfully :)")
	return updatedTaskItems
}