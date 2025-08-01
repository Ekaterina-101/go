package printer

import "fmt"

// func PrintHello(name string) string {
//     return fmt.Sprintf("Hello, %s!", name)
// }

var names = make(map[string]string)

func PrintHello(name string) string {
    names[name] = name
    return fmt.Sprintf("Hello, %s!", name)
}