package main

import "fmt"

func main() {
	users := []string{
		"Leanne Graham",
		"Ervin Howell",
		"Clementine Bauch",
		"Patricia Lebsack",
		"Chelsey Dietrich",
		"Mrs. Dennis Schulist",
		"Kurtis Weissnat",
		"Nicholas Runofsdottir V",
		"Glena Reichert",
		"Clementina DuBuque",
	}
	searchParams(users, "Glena ")
}
func searchParams(users []string, name string) {
	for x := range users {
		if users[x] == name {
			result := []string{users[x]}

			fmt.Println(result)
		}
	}
	result := []string{}
	fmt.Println(result)
}
