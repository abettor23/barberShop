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

	barberShift := 8                              // сколько часов длится смена барбера
	barberCapacity := barberShift * 1             // количество клиентов (1) за час работы барбера
	trimRequest := 1                              // сред. количество посещений в месяц одного мужчины
	clientCapacity := popMen / (30 / trimRequest) // ежедневный поток мужчин на стрижку (среднемесячный)

	if popMen%(30/trimRequest) > 0 {
		clientCapacity += 1
	} //если получилось число с точкой, то увеличиваем поток на 1

	barberNeed := clientCapacity / barberCapacity // количество необходимых барберов в городе

	if clientCapacity%barberCapacity > 0 {
		barberNeed += 1
	} // если число с точкой, то увеличиваем количество нужных барберов на 1

	if barberNeed*barberCapacity*30 < popMen {
		barberNeed++
	} // проверка, сколько клиентов будет обслужено за месяц (из-за дробных чисел), если меньше положенного, то +1 барбер

	fmt.Println("Необходимое число барберов:", barberNeed)

	if barbers > barberNeed {
		fmt.Println("В городе достаточно барберов на данный момент. Все довольны.")
	} else if barbers == barberNeed {
		fmt.Println("В городе ровно столько барберов, сколько требуется!")
	} else {
		fmt.Println("Чтобы покрыть спрос, нужно нанять:", barberNeed-barbers, "барберов")
	}

}
