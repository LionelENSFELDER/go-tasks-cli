package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"go-tasks-cli/db"
)

// flow: 
// Init: load config, init const values
// Check / create database and return data from database
// view all tasks
// wait for a command

type Task struct {
	title string
	done bool
}

func main() {
	Init()
	db.CheckDatabase()
	ViewAllTasks()
	CommandHandler()
}

var tasksList = []Task {}
var reader = bufio.NewReader(os.Stdin)

func createFakeTask(title string, done bool){
	task := Task {
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

	switch userInput {
		case "list":
			ViewAllTasks()
		case "add":
			CreateTask()
		case "del":
			DeleteTask()
		case "tog":
			ToggleTaskState()
		case "togall":
			toggleAllTasks()
		case "untogall":
			untoggleAllTasks()
		case "upd":
			UpdateTaskTitle()
		default:
			fmt.Println("Enter a valid command (list, add, del, tog, togall,untogall, upd).")
			CommandHandler()
	}
}

func GetTaskByIndex(idx int) Task{
	return tasksList[idx - 1]
}

func isValidIndex(index int) bool {
	fmt.Println("tasksList length = ", len(tasksList), index > len(tasksList))
	if index > len(tasksList) || index < 1{
		return false
	}else{
		return true
	}
}

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("Error while reading input!")
	}
	return strings.TrimSpace(input), err
}

func GetInputInt(prompt string, r *bufio.Reader) (int, error) {
	fmt.Println(prompt)
	var input int
	_, err := fmt.Fscanf(r, "%d\n", &input)
	if err != nil {
		log.Fatal("Error while reading input!")
		return 0, err
	}
	return input, err
}

func Scan(r rune) {
	panic("unimplemented")
}

func UpdateTaskTitle(){
	userInput, err := GetInputInt("Index :", reader)
	if err != nil {
		fmt.Println("Error when read index, entrer a valid index !")
		UpdateTaskTitle()
	}

	isValidIndex := isValidIndex(userInput)
	taskIndex := userInput - 1
	if isValidIndex == false {
		fmt.Println("Index is not valid !")
		UpdateTaskTitle()
	}

	newTitle, _ := GetInput("New title :", reader)
	tasksList[taskIndex].title = newTitle

	ViewAllTasks()
	CommandHandler()
}

func ToggleTaskState(){
	taskIndex, err := GetInputInt("index", reader)
	if err != nil {
		fmt.Println("Error when read index, entrer a valid index !")
		ToggleTaskState()
	}

	isValidIndex := isValidIndex(taskIndex)
	if isValidIndex == false {
		fmt.Println("Index is not valid !")
		ToggleTaskState()
	}

	tasksList[taskIndex - 1].done = !tasksList[taskIndex - 1].done
	ViewAllTasks()
	CommandHandler()
}

func toggleAllTasks(){
	for idx := range tasksList {
		tasksList[idx].done = true
	}
	CommandHandler()
}

func untoggleAllTasks(){
	for idx := range tasksList {
		tasksList[idx].done = false
	}
	CommandHandler()
}


func DeleteTask(){
	taskIndex, err := GetInputInt("Index", reader)
	if err != nil {
		fmt.Println("Error when read index, entrer a valid index !")
		DeleteTask()
	}
	isValidIndex := isValidIndex(taskIndex)
	if isValidIndex == false {
		fmt.Println("Index is not valid !")
		DeleteTask()
	}
	index := taskIndex - 1
	tasksList = append(tasksList[:index], tasksList[index+1:]...)
	ViewAllTasks()
	CommandHandler()
}

func isValidTitle(title string) bool{
	if len(title) <= 1 {
		fmt.Println(len(title))
		return true
	}else{
		return false
	}
}

func CreateTask() {
	task := Task{}
	task.done = false
	
	fmt.Println("Title : ")
	reader := bufio.NewReader(os.Stdin)
	taskTitle, _ := reader.ReadString('\n')
	taskTitle = strings.TrimSpace(taskTitle)
	
	if isValidTitle(taskTitle) {
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
	for idx, t := range tasksList {
		var checkbox string
		var color string
		if !t.done {
			color = "\033[31m"
			checkbox = "[ ]"
			}else{
				color = "\033[32m"
			checkbox = "[x]"
		}
		fmt.Println(color, idx + 1, checkbox, t.title + "\033[0m")
	}
	CommandHandler()
}
