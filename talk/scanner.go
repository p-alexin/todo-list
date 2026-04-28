package talk

import (
	"bufio"
	"fmt"
	"os"
	"project-todo/todo"
	"strings"

	"github.com/k0kubun/pp"
)

type sr struct {
	todolist *todo.List
}

var event []string

func Scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	tl := &sr{
		todolist: todo.NewList(),
	}

	PrintPromt()

	for scanner.Scan() {
		inputString := scanner.Text()
		field := strings.Fields(inputString)
		if inputString == "" {
			fmt.Println("unkown command")
			continue
		}
		if field[0] == "add" {
			if len(field) >= 3 {
				text := strings.Join(field[2:], "")
				task := todo.NewTask(field[1], text)
				tl.todolist.AddTask(task)
				fmt.Println("✅ Задача добавлена!")
				event = append(event, inputString)
			} else {
				fmt.Println("❌ Ошибка: формат 'add <заголовок> <описание>")
			}
		} else if field[0] == "list" {
			tl.todolist.CheckList()
			event = append(event, inputString)
		} else if field[0] == "del" {
			if len(field) == 2 {
				tl.todolist.DeleteTask(field[1])
				event = append(event, inputString)
			} else {
				fmt.Println("❌ Ошибка: формат 'del <заголовок>")
			}
		} else if field[0] == "done" {
			if len(field) == 2 {
				tl.todolist.ChangeStatus(field[1])
				fmt.Println("✅ Статус изменен!")
				event = append(event, inputString)
			} else {
				fmt.Println("❌ Ошибка: формат 'done <заголовок>")
			}
		} else if inputString == "help" {
			PrintHelp()
			event = append(event, inputString)
		} else if inputString == "exit" {
			PrintExit()
			event = append(event, inputString)
			return
		} else if inputString == "events" {
			event = append(event, inputString)
			pp.Println(event)
		} else {
			fmt.Println("unkown command")
			continue
		}

	}
}
