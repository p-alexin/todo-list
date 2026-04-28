package todo

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
)

type List struct {
	list map[string]Task
}

func NewList() *List {
	return &List{
		list: make(map[string]Task)}
}

func (l *List) AddTask(task Task) {
	l.list[task.Title] = task
}
func (l *List) CheckList() {
	pp.Println(l.list)
}
func (l *List) ChangeStatus(title string) *Task {
	task, ok := l.list[title]
	if ok {
		if task.Status == true {
			task.Status = false
			now := time.Now()
			task.DoneAt = &now
			l.list[title] = task
			return nil
		} else {
			task.Status = true
			now := time.Now()
			task.DoneAt = &now
			l.list[title] = task
			return nil
		}
	} else {
		fmt.Println("Задача не найдена")
	}
	return &task
}
func (l *List) DeleteTask(title string) {
	task, ok := l.list[title]
	if ok {
		l.list[title] = task
		delete(l.list, title)
		fmt.Println("✅ Задача удалена!")
	} else {
		fmt.Println("Задача не найдена")
	}
}
<<<<<<< HEAD
=======

>>>>>>> 49b1929da00b3ef67763d3b6a09675542bc8ac34
