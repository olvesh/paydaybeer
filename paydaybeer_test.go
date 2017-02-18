package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jinzhu/now"
)

var someSaturydayAfterPayday = now.MustParse("2017-02-25 15:04")
var someNormalWorkday = now.MustParse("2017-02-15 15:04")
var someBoringSaturyday = now.MustParse("2017-02-18 15:04")
var somePaydayWeek = now.MustParse("2017-02-21 15:04")
var somePaydayFridayAt15 = now.MustParse("2017-02-24 15:04")
var somePaydayFridayAt16 = now.MustParse("2017-02-24 16:04")
var someFridayAt15 = now.MustParse("2017-02-17 15:04")
var someFridayAt16 = now.MustParse("2017-02-17 16:04")

func TestCheckBeerSituation(t *testing.T)  {
	assert.Contains(t, checkPaydayBeerSituation(someBoringSaturyday), "Have a nice weekend")
	assert.Contains(t, checkPaydayBeerSituation(someNormalWorkday), "to work")
	assert.Contains(t, checkPaydayBeerSituation(someSaturydayAfterPayday), "Have a nice weekend")
	assert.Contains(t, checkPaydayBeerSituation(someFridayAt15), "Getting thirsty")
	assert.Contains(t, checkPaydayBeerSituation(someFridayAt16), "Weekend here I come")
	assert.Contains(t, checkPaydayBeerSituation(somePaydayFridayAt15), "Yesss!")
	assert.Contains(t, checkPaydayBeerSituation(somePaydayFridayAt16), "Yesss!")
}

func TestWeekend(t *testing.T) {
	assert.True(t, isWeekend(someSaturydayAfterPayday), "%s", someSaturydayAfterPayday)
	assert.False(t, isWeekend(someNormalWorkday), "%s", someSaturydayAfterPayday)
}
func TestSaturdayAfterPayday(t *testing.T) {
	assert.False(t, isPaydayWeek(someSaturydayAfterPayday), "%s", someSaturydayAfterPayday)
}

func TestSomeBoringSaturday(t *testing.T) {
	assert.False(t, isPaydayWeek(someBoringSaturyday), "%s", someBoringSaturyday)
	assert.False(t, isBeerday(someBoringSaturyday), "%s", someBoringSaturyday)
	assert.False(t, isBeerOClock(someBoringSaturyday), "%s", someBoringSaturyday)
}

func TestSomeFridayJustBeforeBeertime(t *testing.T) {
	assert.True(t, isBeerday(someFridayAt15), "%s", someFridayAt15)
	assert.False(t, isBeerOClock(someFridayAt15), "%s", someFridayAt15)
}

func TestSomeFridayJustAfterBeertime(t *testing.T) {
	assert.True(t, isBeerday(someFridayAt16), "%s", someFridayAt16)
	assert.True(t, isBeerOClock(someFridayAt16), "%s", someFridayAt16)
}

func TestPayWeekFridayJustAfterBeertime(t *testing.T) {
	assert.False(t, isPaydayWeek(someFridayAt16), "%s", someFridayAt16)
	assert.True(t, isBeerday(someFridayAt16), "%s", someFridayAt16)
	assert.True(t, isBeerOClock(someFridayAt16), "%s", someFridayAt16)
}

func TestPayWeek(t *testing.T) {
	assert.True(t, isPaydayWeek(somePaydayWeek), "%s", somePaydayWeek)
	assert.False(t, isBeerday(somePaydayWeek), "%s", somePaydayWeek)
	assert.False(t, isBeerOClock(somePaydayWeek), "%s", somePaydayWeek)
}





