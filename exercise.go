package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
 
// Main function to display the menu and handle user input
 
// Create a structure Task having name and done
 
type Task struct {
    name        string
    description string
    isDone      bool
}
 
func AddTask(tasks *[]Task, reader *bufio.Reader) {
    fmt.Print("Enter the title of the Task: ")
    title, _ := reader.ReadString('\n')
 
    fmt.Print("Enter the description of the Task: ")
    desc, _ := reader.ReadString('\n')
 
    newTask := Task{
        name:        title,
        description: desc,
        isDone:      false,
    }
 
    *tasks = append(*tasks, newTask)
}
 
func ListTasks(tasks []Task) {
    // Iterate over all tasks and print
 
    for i, value := range tasks {
        fmt.Println("Task id: ", i)
        fmt.Print("Task name: ", value.name)
        fmt.Println("Task desc: ", value.description)
 
    }
}
func MarkASDone(tasks []Task,reader *bufio.Reader){
	fmt.Print("Enter the description of the Task: ")
    input, _ := reader.ReadString('\n')
	id,_ := strconv.Atoi(strings.TrimSpace(input))
	tasks[id].isDone=true
}
 
func main() {
 
    reader := bufio.NewReader(os.Stdin)
    var tasks []Task
    for {
        fmt.Println("1. Add Task")
        fmt.Println("2. List Tasks")
        fmt.Println("3. Mark Task as Done")
        fmt.Println("4. Delete Task")
        fmt.Println("5. Exit")
        fmt.Print("Enter choice: ")
 
        input, _ := reader.ReadString('\n')
        choice, err := strconv.Atoi(strings.TrimSpace(input))
 
        switch choice {
        case 1:
 
            AddTask(&tasks, reader)
        case 2:
            ListTasks(tasks)
		case 3:
			MarkAsDone()
 
        }
 
        if err != nil {
            fmt.Println("Invalid choice, please try again.")
            continue
        }
 
    }
}