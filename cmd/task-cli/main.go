package main 


import "fmt"


type Task struct {
    ID       int
    Title    string
    Status   string
    Priority int
}

type TaskStore interface {
    Add(title string, priority int) (Task, error)
    Get(id int) (Task, error)
    Complete(id int) error
    Delete(id int) error
    List() []Task
}

type MemoryStore struct {
    tasks  []Task
    nextID int
}


func main() {
	fmt.Println()
}

func NewMemoryStore() *MemoryStore {
    return &MemoryStore{nextID: 1}
}





