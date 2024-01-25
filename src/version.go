package src

import "fmt"

const vnersion = "1.0.0"

const name = "go-curd"

func GetVersion() string {
	return fmt.Sprintf("%s\t%s", name, vnersion)
}
