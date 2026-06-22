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

func main() {
	// Init()
	fmt.Println("Welcome to GO Tasks CLI !")
	db.CheckDatabase()
	tasks, err := db.ReturnDataFromDatabase()
	if err != nil {
		log.Fatalf("Error while loading tasks from database: %v", err)
	}
	tasksList = tasks
	ViewAllTasks()
	CommandHandler()
}

var tasksList = []db.Task{}
var reader = bufio.NewReader(os.Stdin)

func createFakeTask(title string, done bool){
	task := db.Task {
		Title: title ,
		Done: done,
	}
	tasksList = append(tasksList, task)
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

func GetTaskByIndex(idx int) db.Task{
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
		tasksList[idx].Done = true
	}
	ViewAllTasks()
	CommandHandler()
}

func untoggleAllTasks(){
	for idx := range tasksList {
		tasksList[idx].Done = false
	}
	db.SaveDataToDatabase(tasksList)
	ViewAllTasks()
	CommandHandler()
}

func UpdateTaskTitle(idx int){
	newTitle, _ := GetInput("New title :", reader)
	tasksList[idx - 1].Title = newTitle
	db.SaveDataToDatabase(tasksList)
	ViewAllTasks()
	CommandHandler()
}

func ToggleTaskState(idx int){
	tasksList[idx - 1].Done = !tasksList[idx - 1].Done
	db.SaveDataToDatabase(tasksList)
	ViewAllTasks()
	CommandHandler()
}

func DeleteTask(idx int){
	tasksList = append(tasksList[:idx - 1], tasksList[idx:]...)
	db.SaveDataToDatabase(tasksList)
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
	task := db.Task{}
	task.Done = false
	
	fmt.Println("Title : ")
	reader := bufio.NewReader(os.Stdin)
	taskTitle, _ := reader.ReadString('\n')
	taskTitle = strings.TrimSpace(taskTitle)
	
	if isValidTitle(taskTitle) {
		fmt.Println("No empty title !")
		CreateTask()
	}else{
		task.Title = taskTitle
		AddTask(task)
	}
	CommandHandler()
}

func AddTask(task db.Task){
	tasksList = append(tasksList, task)
	db.SaveDataToDatabase(tasksList)
	ViewAllTasks()
	CommandHandler()
}

func ViewAllTasks(){
	fmt.Println("Tasks List:")
	for idx, t := range tasksList {
		var checkbox string
		var color string
		if !t.Done {
			color = "\033[31m"
			checkbox = "[ ]"
			}else{
				color = "\033[32m"
			checkbox = "[x]"
		}
		fmt.Println(color, idx + 1, checkbox, t.Title + "\033[0m")
	}
	CommandHandler()
}
