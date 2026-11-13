package main

import "fmt"

type Table struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func main() {
	tasks := []Table{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}


	found := findTaskByID(tasks, 2)
	if found != nil {
		found.Complete()
	}

	for _, task := range tasks {
		fmt.Println(task.Summary())
	}
}

func (t Table) Summary() string {
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Table) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

func (t *Table) Complete() {
	t.Status = "done"
}

func findTaskByID(tasks []Table, id int) *Table {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}