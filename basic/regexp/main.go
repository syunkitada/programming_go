package main

import (
	"fmt"
	"regexp"
)

func main() {
	matchString, err := regexp.MatchString("hoge([a-z]+)piyo", "hogetestpiyo")
	fmt.Println("regexp.MatchString", matchString, err)

	regexpHoge := regexp.MustCompile("hoge([a-z]+)piyo")
	regexpHogeMatchString := regexpHoge.MatchString("hogetestpiyo")
	fmt.Println("regexp.MatchString", regexpHogeMatchString, err)
}
