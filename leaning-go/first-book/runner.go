package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Vérification du nombre d'arguments
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run runnner.go <limite_haute>")
        return
    }

    // Conversion de l'argument en entier
    limiteStr := os.Args[1]
    limite, err := strconv.Atoi(limiteStr)
    if err != nil {
        fmt.Println("L'argument doit être un nombre entier.")
        return
    }

    // Boucle for avec limite basse 0 et limite haute limite
    for i := 0; i <= limite; i++ {
        switch {
        case i%3 == 0 && i%5 == 0:
            fmt.Println("baz")
        case i%3 == 0:
            fmt.Println("foo")
        case i%5 == 0:
            fmt.Println("bar")
        }
    }
}