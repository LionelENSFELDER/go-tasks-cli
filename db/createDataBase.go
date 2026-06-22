package db

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

// struct of task
type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// get current os
func GetOS() string {
	return runtime.GOOS
}

// get current user
func GetUser() string {
    if runtime.GOOS == "windows" {
        return os.Getenv("USERNAME")
    }
	return os.Getenv("USER")
}

// define path of database (windows, linux, mac)
var databasePathWindows = "C:/Users/" + GetUser() + "/AppData/Roaming/go-tasks-cli-data.json"
var databasePathLinux = "/home/" + GetUser() + "/.config/go-tasks-cli-data.json"
var databasePathMac = "/Users/" + GetUser() + "/Library/Application Support/go-tasks-cli-data.json"

func GetDatabasePath() string {
    currentOS := GetOS()
    // get database path depending on the os
    switch currentOS {
        case "windows":
            return databasePathWindows
        case "linux":
            return databasePathLinux
        case "darwin":
            return databasePathMac
		default:
			return ""
	}
}


// check if database exist in path define previously
func isDatabaseExist() bool {
	currentOS := GetOS()
    // check if database exist depending on the os
    switch currentOS {
        case "windows":
            if _, err := os.Stat(databasePathWindows); os.IsNotExist(err) {
                return false
            }
        case "linux":
            if _, err := os.Stat(databasePathLinux); os.IsNotExist(err) {
                return false
            }
        case "darwin":
            if _, err := os.Stat(databasePathMac); os.IsNotExist(err) {
                return false
            }
		default:
			return true
	}
    return true
}

func ReturnDataFromDatabase() ([]Task, error) {
    file, err := os.Open(GetDatabasePath())
    if err != nil {
        return nil, fmt.Errorf("Could not open database for reading: %w", err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    var data []Task
    if err := decoder.Decode(&data); err != nil {
        return nil, fmt.Errorf("Could not decode database: %w", err)
    }
    return data, nil
}

func CreateDatabase() error {
    data := []Task{
        {Title: "Create your first task", Done: false},
        {Title: "Update a task", Done: true},
        {Title: "Toggle a task", Done: false},
    }

    file, err := os.Create(GetDatabasePath())
    if err != nil {
        return fmt.Errorf("Could not create database: %w", err)
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    if err := encoder.Encode(data); err != nil {
        return fmt.Errorf("Could not encode database: %w", err)
    }
    return nil
}

// check if database exist and create it if not
func CheckDatabase() {
    fmt.Println("Checking database...")
    currentOS := GetOS()
    fmt.Printf("Current operating system: %s\n", currentOS)
    if currentOS != "windows" && currentOS != "linux" && currentOS != "darwin" {
        println("Unsupported operating system. This application only supports Windows, Linux, and macOS.")
    }
    // check if database exist and create it if not
    isDatabaseExist := isDatabaseExist()
    fmt.Printf("Database exist: %t, Database path: %s\n", isDatabaseExist, GetDatabasePath())
    if !isDatabaseExist {
        // display message on console
        println("Database not found, creating database...")
        CreateDatabase()
        println("Database created !")
    }
}

func SaveDataToDatabase(tasks []Task) error {
    file, err := os.Create(GetDatabasePath())
    if err != nil {
        return fmt.Errorf("Could not open database for writing: %w", err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    return encoder.Encode(tasks)
}