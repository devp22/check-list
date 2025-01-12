package main

import (
	"fmt"
	"net/http"
)
var taskItems = []string{"Read Book","Clean House","Get Groceries"}

func main() {
	//creating an http server endpoint

	http.HandleFunc("/",greet)
	http.HandleFunc("/show",showTasks)

	http.ListenAndServe(":8080",nil) // first param is to specify port of server, second is to specify what if it did not worked (then 'nil', means do nothing)

}

func greet(w http.ResponseWriter, r *http.Request) {
	var message = "############## Welcome to to-do list app ##############"
	fmt.Fprintf(w,"%s", message)
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	for index, task := range taskItems {
		fmt.Fprintf(w,"%d. %s\n",index+1,task)
	}
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