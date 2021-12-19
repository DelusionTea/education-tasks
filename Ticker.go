package main

import (
	"fmt"
	"time"
)

func main() {
	//допишите код здесь
	//Используя Ticker, напишите программу,
	//которая каждые 2 секунды 10 раз подряд выводит разницу в секундах между текущим временем и временем запуска программы.
	//Выводить нужно только целую часть секунд.
	//Обратите внимание на пример в разделе Sleep и на методы Sub и Seconds. Привести к целому числу можно с помощью int(seconds).
	startTime := time.Now()
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		anytime := time.Now()
		duration := anytime.Sub(startTime)
		fmt.Println("iteration", i+1, duration)
	}
}
