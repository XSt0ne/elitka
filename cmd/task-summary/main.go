package main

import (
	"fmt"
)


func main() {
	statuses := []string{"new", "in_progress", "done","done"}
	for i,v:=range statuses {
		switch v {
			case "new":
				fmt.Println(i+1,statusLabel(v))
			case "in_progress":
				fmt.Println(i+1,statusLabel(v))
			case "done":
				fmt.Println(i+1,statusLabel(v))
			default:
				fmt.Println(i+1,statusLabel(v))
		}
	}
	count_tasks:=""
	fmt.Println("Какие задачи вы хотие подсчитать:",count_tasks)
	fmt.Scan(&count_tasks)
	println("Количество",countStatus(count_tasks,statuses))
	counts := make(map[string]int)
	for _, status := range statuses {
    	counts[status]++
	}
	fmt.Println("Статистика:",counts)
}


func statusLabel(status string) string{
	res:=""
	switch status {
		case "done":
			res = "Готово к сдаче"
		case "in_progress":
			res = "В процессе выполнения"
		case "new":
			res = "Новая задача"
		default:
			res = "Ошибка обработки статуса"
	}
	return res
}


func countStatus (status string, statuses []string) int {
	count:=0
	for _,v:= range statuses {
		if v== status {
			count++
		}
	}
	return count
}