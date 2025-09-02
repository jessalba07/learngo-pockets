package main
import "fmt"

func main() {
    greeting := greet("fr")

    fmt.Println(greeting)
}

// language represents the language's code
type language string // declares a type

func greet(l language) string {
    switch l {
    case "en":
        return "Hello, World!"
    case "fr":
        return "Bonjour le monde"     
    default:
        return ""
    }
}