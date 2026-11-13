// Command snyk-demo-go-2 uses public modules only: no dependencies on other brightsheets repos.
package main

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println("snyk-demo-go-2", uuid.NewString(), language.Make("en-US").String())
}
