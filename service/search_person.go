package service

import (
	"fmt"
	"strings"
)

func SearchParams(users []string, name string) {
	found := false

	for _, user := range users {
		if strings.Contains(strings.ToLower(user), strings.ToLower(name)) {
			result := []string{user}
			fmt.Println(result)
			found = true
		}
	}
	if !found {
		result := []string{}
		fmt.Println(result)
	}

}
