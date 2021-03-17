package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Eater struct {
	name  string
	speed int32
}

var hotdogsCount int32 = 1000
var wg = sync.WaitGroup{}
var mutex = sync.Mutex{}

func main() {
	fmt.Println("В соревновании по поеданию хотдогов участвует три едока. " +
		"Введите количество хотдохов, которое поедает каждый участник в секунду, " +
		"чтобы узнать когда закончится соревнование. Всего хотдогов 1000")
	fmt.Println("Участник №1")
	var eaterOneSpeed int32
	fmt.Scan(&eaterOneSpeed)
	eaterOne := Eater{"Участник №1", eaterOneSpeed}
	fmt.Println("Участник №2")
	var eaterTwoSpeed int32
	fmt.Scan(&eaterTwoSpeed)
	eaterTwo := Eater{"Участник №2", eaterTwoSpeed}
	fmt.Println("Участник №3")
	var eaterThreeSpeed int32
	fmt.Scan(&eaterThreeSpeed)
	eaterThree := Eater{"Участник №3", eaterThreeSpeed}

	wg.Add(3)
	go eat(eaterOne)
	go eat(eaterTwo)
	go eat(eaterThree)
	wg.Wait()
}

func eat(eater Eater) {
	ticker := time.NewTicker(1 * time.Second)
	time := 0
	for range ticker.C {
		mutex.Lock()
		time++
		hotdogsCount -= eater.speed
		if hotdogsCount <= 0 {
			fmt.Printf("Соревнование закончилось спустя %v секунд. Последнюю сосиску съел %v",
				time,
				eater.name)
			os.Exit(0)
		}
		fmt.Printf("Осталось сосисок: %v\n", hotdogsCount)
		mutex.Unlock()
	}
}
