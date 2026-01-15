package db

import (
	"encoding/json"
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

func ReturnDataFromDatabase() []Task {
    file, err := os.Open(GetDatabasePath())
    if err != nil {
        panic(err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    var data []Task
    if err := decoder.Decode(&data); err != nil {
        panic(err)
    }
    return data
}

func CreateDatabase() {
    data := []Task{
        {Title: "Meeting (from database)", Done: false},
        {Title: "Get food (from database)", Done: true},
        {Title: "Update obsidian note (from database)", Done: false},
    }

    file, err := os.Create(GetDatabasePath())
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    if err := encoder.Encode(data); err != nil {
        panic(err)
    }
}

// check if database exist and create it if not
func CheckDatabase() {
    if !isDatabaseExist() {
        // display message on console
        println("Database not found, creating database...")
        CreateDatabase()
        println("Database created !")
    }
}