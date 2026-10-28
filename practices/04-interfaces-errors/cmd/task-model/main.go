package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	return "#" + fmt.Sprint(t.ID) + " " + fmt.Sprint(t.Title) + " [" + fmt.Sprint(t.Status) + "], приоритет " + fmt.Sprint(t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	if 1 <= t.Priority && t.Priority <= 3 {
		return true
	}
	return false
}

func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	ptr := findTaskByID(tasks, 2)
	if ptr != nil {
		ptr.Complete()
		ptr.Priority = 3
	}
	tasks_done := 0
	for _, task := range tasks {
		fmt.Println(task.Summary())
		if task.Status == "done" {
			tasks_done++
		}
	}
	fmt.Println("Задач завершено:", tasks_done)
	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}
}
