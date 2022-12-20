package colors

import "fmt"

const RESET = "\033[0m"

// also bolded
const GREEN = "\033[31;92m"
const ORANGE = "\033[31;91m"
const YELLOW = "\033[31;93m"
const TEAL = "\033[31;96m"

func Color(s, color string) string {
	return fmt.Sprintf("%s%s%s", color, s, RESET)
}
