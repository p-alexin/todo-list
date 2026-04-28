package talk

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func PrintPromt() {
	pp.Println("--------------------------TODO LIST--------------------------")
	fmt.Print("ВВЕДИТЕ КОМАНДУ: ")
}

func PrintHelp() {
	pp.Println("add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов}")
	fmt.Println("эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("")
	pp.Println("list")
	fmt.Println("эта команда позволяет получить полный список всех задач")
	fmt.Println("")
	pp.Println("del {заголовок существующей задачи}")
	fmt.Println("эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("")
	pp.Println("done {заголовок существующей задачи}")
	fmt.Println("эта команда позволяет отметить задачу как выполненную")
	fmt.Println("")
	pp.Println("events")
	fmt.Println("эта команда позволяет получить список всех событий")
	fmt.Println("")
	pp.Println("exit")
	fmt.Println("эта команда позволяет завершить выполнение программы")
}

func PrintExit() {
	fmt.Println("-------Завершение______________________программы--------")
}
