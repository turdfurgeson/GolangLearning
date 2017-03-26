package main

import "fmt"

func main() {
	for i := 50; i <= 52203; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
		/**Don't use '' around i because it won't point to variable anymore. It will
		just return 105, which is the int32 for i**/

	}
}
