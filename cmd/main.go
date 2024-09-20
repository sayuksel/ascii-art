package main

import (
	"ascii/PKG"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//enting file
	if len(os.Args) == 1 {
		log.Fatal("Wrong number of parameters!!!! ")
		return
	}
	arr := os.Args[1:]
	if len(arr) > 2 {
		log.Fatal("Wrong number of parameters!!!! ")
	}
	if len(arr) == 1 {
		arr = append(arr, "..//banners/standard.txt")
	} else {
		switch arr[1] {
		case "standard":
			arr[1] = ("..//banners/standard.txt")
		case "shadow":
			arr[1] = ("..//banners/shadow.txt")
		case "thinkertoy":
			arr[1] = ("..//banners/thinkertoy.txt")
		default:
			fmt.Println("Wrong banner name !!!!")
			return
		}
	}

	arr[0], arr[1] = arr[1], arr[0]

	//entry of txt
	for _, v := range arr[1] {
		if v < 32 || v > 126 {
			log.Fatal("wrong text entered!!! ")
		}
	}

	if len(arr[1]) == 0 {
		return
	}

	sections := strings.Split(arr[1], "\\n")
	if len(sections) >= 2 && sections[0] == "" && sections[1] == "" {
		fmt.Println()
		sections = sections[2:]

	}
	for _, s := range sections {
		if s == "" {

			fmt.Println()
			continue
		}
		cache := [8]string{}
		for _, r := range s {
			PKG.Strings(arr[0], r, &cache)
		}
		PKG.PrintA(&cache)
	}
}
