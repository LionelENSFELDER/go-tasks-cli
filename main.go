package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
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

func RandomId() int {
	return rand.Int() / 100000000000000000
}

func createId() int {
	// TODO : fix duplicate id by display by index and delete id
	id := RandomId()
	isIdExist := GetTaskById(id)
	if isIdExist != -1 {
		return id
	}else{
		return RandomId()
	}
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
	} else if userInput == "add"{
		CreateTask()
	} else if userInput == "del"{
		DeleteTask()
	} else if userInput == "tog"{
		ToggleTaskState()
	} else if userInput == "upd"{
		UpdateTaskTitle()
	}else{
		fmt.Println("Enter a valid command (list, add, del, tog, upd).")
		CommandHandler()
	}
}

func GetTaskById(id int) int{
	for i, v := range tasksList {
		if v.id == id {
			return i
		}
	}
	return -1
}

func UpdateTaskTitle(){
	// TODO : fix bug user cannot type new title
	var id int
	
	fmt.Println("id : ")
	fmt.Scan(&id)
	taskIndex := GetTaskById(id)
	fmt.Println("=> Current title : ", tasksList[taskIndex].title)
	
	fmt.Println("New title : ")
	reader := bufio.NewReader(os.Stdin)
	newTitle, _ := reader.ReadString('\n')
	fmt.Println("=> New title : ",  newTitle, len(newTitle))

	if len(newTitle) < 2 {
		fmt.Println("No empty title !")
		UpdateTaskTitle()
	}else{
		tasksList[taskIndex].title = newTitle
		ViewAllTasks()
		CommandHandler()
	}

}

func ToggleTaskState(){
	var id int
	fmt.Print("id : ")
	fmt.Scan(&id)
	taskIndex := GetTaskById(id)
	tasksList[taskIndex].done = !tasksList[taskIndex].done
	ViewAllTasks()
	CommandHandler()
}

func DeleteTask(){
	var id int
	fmt.Print("id : ")
	fmt.Scan(&id)
	taskIndex := GetTaskById(id)
	tasksList = append(tasksList[:taskIndex], tasksList[taskIndex+1:]...)
	ViewAllTasks()
	CommandHandler()
}

func CreateTask() {
	task := Task{}
	task.id =  createId()
	task.done = false
	
	fmt.Println("Title : ")
	reader := bufio.NewReader(os.Stdin)
	taskTitle, _ := reader.ReadString('\n')
	
	if len(task.title) == 2 {
		fmt.Println("No empty title !")
		CreateTask()
	}else{
		task.title = taskTitle
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
