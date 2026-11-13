package main


import "fmt"


func main() {
	minut:=0
	fmt.Println("Программа для перевода минут в часы")
	fmt.Println("Введите количество минут:")
	fmt.Scan(&minut)
	hours:=minut/60
	ostat_minut:=minut-hours*60
	fmt.Println("Количество полных часов:",hours)
	fmt.Println("Количество оставшихся минут:",ostat_minut)
}