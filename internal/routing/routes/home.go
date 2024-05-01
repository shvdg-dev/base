package routes

import "fmt"

func Home() string {
	return fmt.Sprintf("/%s", "home")
}
