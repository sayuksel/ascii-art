package PKG

import "fmt"

func PrintA(arr *[8]string) {
	for i, line := range arr {
		fmt.Println(line)
		arr[i] = ""
	}
	
}
