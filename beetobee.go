package beetobee

import (
	"fmt"

	"github.com/telcomate/beetobee/v2/bumblebee"
)

var Text = "all your base are belong to us"

func Bees() {
	fmt.Println(Text)
}

func Bumblebees() {
	fmt.Println(bumblebee.Text)
}
