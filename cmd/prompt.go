package cmd

import "fmt"

func confirm(q string) bool {
	a := ""
	for {
		fmt.Printf(q)
		fmt.Scanln(&a)
		switch a {
		case "y", "Y", "yes":
			return true
		case "n", "N", "no":
			return false
		default:
			continue
		}
	}
}
