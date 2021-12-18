package main

import (
	"fmt"
	"time"
)

//Андрей родился 26 ноября 1993 года. Посчитайте количество дней до его 100-летия — относительно сегодняшнего дня.
//Local — текущий часовой пояс операционной системы

func main() {
	// допишите код здесь
	//newYear := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

	now := time.Now()
	fmt.Println(now)
	birthday := time.Date(1993, time.November, 26, 0, 0, 0, 0, time.Local)
	//daysAlready := now.Sub(birthday)
	resultbirthday := birthday.AddDate(100, 0, 0)
	result := now.Sub(resultbirthday)
	days := result.Hours() / 24.0
	years := days / 365.0
	fmt.Println(days)
	fmt.Println(years)
}
