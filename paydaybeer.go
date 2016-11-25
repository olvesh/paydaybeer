package main

import (
  "fmt"
  "time"
)

func main() {

  date := time.Now()
  isPaydayWeek := (25 - date.Day()) < 7
  isFriday := date.Weekday() == time.Friday

  if (isPaydayWeek && isFriday){
    fmt.Println("Yesss! 🍺  it is Pay 🍹 Day 🍸 Beer! 🍻")
  } else {
    fmt.Println("Back 😰 to work... 👷🏼")
  }

}
