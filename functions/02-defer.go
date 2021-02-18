package functions

import "fmt"

func deferExample() {

	defer fmt.Println("I finish second although I appear first in the execution flow, the magic of defer!")
	fmt.Println("I finish first although I appear second in the execution flow, och defer what have u done!")
}
