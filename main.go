package main

import (
	"bufio"
	"fmt"
	"go-tasks-cli/db"
	"log"
	"os"
	"strconv"
	"strings"
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

func isValidCommand(cmd string) bool{
	validCommands := []string{
		"list", 
		"add", 
		"del", 
		"tog", 
		"togall", 
		"untogall", 
		"upd",
	}
	for _, c := range validCommands {
		if cmd == c {
			return true
		}
	}
	return false
}

func isEmptyArgs(args []string) bool{
	if len(args) == 0 {
		return true
	}else{
		return false
	}
}

func convertStringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Error while converting string to int !")
		return 0
	}
	return num
}

func CommandHandler(){
	fmt.Println("What do you want to do ?")
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(strings.TrimSpace(line))

	if(len(parts) == 0 || isValidCommand(parts[0]) == false) {
		fmt.Println("Enter a valid command (list, add, del, tog, togall,untogall, upd).")
		CommandHandler()
		return
	} else if (parts[0] == "del" || parts[0] == "tog" || parts[0] == "upd") && isEmptyArgs(parts[1:]) == true {
		fmt.Println("Enter an index as argument ! (ex: ", parts[0], " 1)")
		CommandHandler()
		return
	} else if (parts[0] == "del" || parts[0] == "tog" || parts[0] == "upd") && isValidIndex(convertStringToInt(parts[1])) == false {
		fmt.Println("Enter a valid index as argument ! (ex: ", parts[0], " 1)")
		CommandHandler()
		return
	}

	cmd := parts[0]
	args := parts[1:]

	switch cmd {
		case "del":
			DeleteTask(convertStringToInt(args[0]))
		case "tog":
			ToggleTaskState(convertStringToInt(args[0]))
		case "upd":
			UpdateTaskTitle(convertStringToInt(args[0]))
		case "list":
			ViewAllTasks()
		case "add":
			CreateTask()
		case "togall":
			toggleAllTasks()
		case "untogall":
			untoggleAllTasks()
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

func toggleAllTasks(){
	for idx := range tasksList {
		tasksList[idx].done = true
	}
	ViewAllTasks()
	CommandHandler()
}

func untoggleAllTasks(){
	for idx := range tasksList {
		tasksList[idx].done = false
	}
	ViewAllTasks()
	CommandHandler()
}

func UpdateTaskTitle(idx int){
	newTitle, _ := GetInput("New title :", reader)
	tasksList[idx - 1].title = newTitle
	ViewAllTasks()
	CommandHandler()
}

func ToggleTaskState(idx int){
	tasksList[idx - 1].done = !tasksList[idx - 1].done
	ViewAllTasks()
	CommandHandler()
}

func DeleteTask(idx int){
	tasksList = append(tasksList[:idx - 1], tasksList[idx:]...)
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
