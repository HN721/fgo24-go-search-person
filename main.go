package main

import (
	"search/service"
)

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
	service.SearchParams(users, "Glena")
}
