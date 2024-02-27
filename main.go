package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

type Task struct {
	id int 
	title string
	description string
	done bool
}

var tasksList = []Task {}

func main() {
	Init()
	CommandHandler()
}

func CreateTask() {
	task := Task{}

	fmt.Print("Enter task title: ")
	reader := bufio.NewReader(os.Stdin)
	taskTitle, err := reader.ReadString('\n')
	// TODO : wait for user input
	if err != nil{
		fmt.Print("Enter task title: ")
		return
	}
	task.title = taskTitle

	task.description = "description"
	task.id = 3
	task.done = false

	AddTask(task)
}

func AddTask(task Task){
	tasksList = append(tasksList, task)
	ViewAllTasks()
}

func Init(){
	fmt.Printf("Todo List \n")
	task1 := Task {
		id: 1,
		title: "Meeting" ,
		description: "start meeting without Lionel",
		done: true,
	}
	task2 := Task {
		id: 2,
		title: "Cooking" ,
		description: "cook smoked chicken !",
		done: false,
	}
	tasksList = append(tasksList, task1, task2)
}

func CommandHandler(){
	var userInput string
	fmt.Println("What do you want to do ?")
	fmt.Scan(&userInput)
	if userInput == "list"{
		ViewAllTasks()
		}else if userInput == "add"{
			CreateTask()
			}else{
				fmt.Print("Enter a valid command (list, add, update, delete).")
			}
}

func ViewAllTasks(){
	fmt.Println("Tasks List:")
	for _, t := range tasksList {
		fmt.Println("ID:", t.id)
		fmt.Println("Title:", t.title)
		fmt.Println("Description:", t.description)
		fmt.Println("Done:", t.done)
		fmt.Println()
	}
}
