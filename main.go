package main

import (
	"fmt"
)

func main() {
	fmt.Println("Какое население города на данный момент?")

	var popAll int
	fmt.Scan(&popAll)

	popMen := popAll / 2

	fmt.Println("Так получилось, что количество мужчин и женщин в этом городе исторически равное.")
	fmt.Println("Поэтому мужчин тут", popMen, "человек и всем нужен уход за бородой.")

	fmt.Println("Сколько барберов работает в городе на данный момент?")

	var barbers int
	fmt.Scan(&barbers)

	barberShift := 8                                          // сколько часов длится смена барбера
	barberCapacity := barberShift * 1                         // количество клиентов (1) за час работы барбера
	trimRequest := 1                                          // сред. количество посещений в месяц одного мужчины
	clientCapacity := popMen / (30 / trimRequest)             // ежедневный поток мужчин на стрижку (среднемесячный)
	barberNeed := (clientCapacity / barberCapacity) - barbers // количество необходимых барберов, чтобы покрыть спрос

	if barberNeed <= 0 {
		fmt.Println("В городе достаточно барберов на данный момент. Все довольны.")
	}

	if barberNeed > 0 {
		fmt.Println("Чтобы покрыть спрос, нужно еще нанять:", barberNeed, "барберов")
	}

}
