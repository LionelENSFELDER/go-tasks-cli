package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	id int 
	title string
	done bool
}

func main() {
	Init()
	CommandHandler()
}

var tasksList = []Task {}

func createId() int{
	return len(tasksList) + 1
}

func createFakeTask(title string, done bool){
	task := Task {
		id:  createId(),
		title: title ,
		done: done,
	}
	tasksList = append(tasksList, task)
}

func Init(){
	createFakeTask("Meeting", false)
	createFakeTask("Get food", false)
	createFakeTask("Call Lionel Ensfelder", true)
	createFakeTask("Launch with team", false)
	createFakeTask("Learn GO !", true)
}

func CommandHandler(){
	var userInput string
	fmt.Println("What do you want to do ?")
	fmt.Scanln(&userInput)
	if userInput == "list"{
		ViewAllTasks()
		}else if userInput == "add"{
			CreateTask()
			}else if userInput == "del"{
				DeleteTask()
			}else{
				fmt.Println("Enter a valid command (list, add, update, delete).")
				CommandHandler()
			}
}

func DeleteTask(){
	var id int
	fmt.Print("id : ")
	fmt.Scan(&id)
	for i, v := range tasksList {
		if v.id == id {
			tasksList = append(tasksList[:i], tasksList[i+1:]...)
			ViewAllTasks()
			CommandHandler()
		}
	}
}

func CreateTask() {
	task := Task{}
	task.id =  createId()
	task.done = false
	
	fmt.Println("Title : ")
	reader := bufio.NewReader(os.Stdin)
	taskTitle, _ := reader.ReadString('\n')
	fmt.Println("task.title and taskTitle lenght : ", len(task.title), len(taskTitle))
	task.title = taskTitle

	if len(task.title) == 2 {
		fmt.Println("task haven't title !")
		CreateTask()
	}else{
		AddTask(task)
	}
	CommandHandler()
}

func AddTask(task Task){
	tasksList = append(tasksList, task)
	ViewAllTasks()
	CommandHandler()
}

func ViewAllTasks(){
	fmt.Println("Tasks List:")
	for _, t := range tasksList {
		var checkbox string
		var color string
		if !t.done {
			color = "\033[31m"
			checkbox = "[ ]"
			}else{
				color = "\033[32m"
			checkbox = "[x]"
		}
		fmt.Println(color + checkbox, t.id, t.title + "\033[0m")
	}
	CommandHandler()
}
