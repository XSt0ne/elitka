package main


import "fmt"


func main() {
	course := "Бэкенд разработка на Go"
	name :=  "Александр"
 	group := "ЭФБО-05-26"
	completed := 1
	var ready bool = true


	fmt.Println("Участник:", name)
	fmt.Printf("Группа: %d\n", group)
	fmt.Printf("Курс: %d\n", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к backend: %t\n", ready)
}