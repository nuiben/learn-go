

package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

//makes use of multi-modules found in the golang example directory 'stringutil'

func main() {
	fmt.Println(stringutil.ToUpper("Hello"))
}
