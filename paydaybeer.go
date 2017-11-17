package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(checkPaydayBeerSituation(time.Now()))
}

func isPaydayWeek(beertime time.Time) bool {
	return (18 - beertime.Day()) < 7 && beertime.Weekday() <= 5
}

func isBeerday(beertime time.Time) bool {
	return beertime.Weekday() == time.Friday
}

func isBeerOClock(beertime time.Time) bool {
	return isBeerday(beertime) && beertime.Hour() >= 16
}

func isWeekend(weekend time.Time) bool {
	return weekend.Weekday() == 6 || weekend.Weekday() == 7
}

func checkPaydayBeerSituation(possibleBeerOClock time.Time) string {
	
	switch {
	case isWeekend(possibleBeerOClock)  :
		return "Have a nice weekend!"
	case isPaydayWeek(possibleBeerOClock) && isBeerday(possibleBeerOClock) :
		return "Yesss! 🍺  it is time for Pay 🍹 Day 🍸 Beer! 🍻"
		
	case isPaydayWeek(possibleBeerOClock) && !isBeerday(possibleBeerOClock):
		return "Patience, good things and money will come... ⏳ "
		
	case isBeerOClock(possibleBeerOClock) :
		return "Weekend here I come! Someones up for a 🍺  ?"
		
	case isBeerday(possibleBeerOClock) :
		return "Getting thirsty for🍺  ?"
		
	default:
		return  "Back 😰 to work... 👷🏼"
	}
}
