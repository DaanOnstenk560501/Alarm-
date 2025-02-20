package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "strings"
)

func main() {
        var aantal int
        var err error

        reader := bufio.NewReader(os.Stdin)

        for {
                fmt.Print("Voer een positief geheel getal in: ")
                input, _ := reader.ReadString('\n')
                input = strings.TrimSpace(input)

                aantal, err = strconv.Atoi(input)
                if err == nil && aantal > 0 {
                        break
                }

                fmt.Println("Ongeldige invoer. Probeer het opnieuw.")
        }

        for i := 1; i <= aantal; i++ {
                fmt.Printf("Alarm %d!\n", i)
        }

        fmt.Println("Druk op Enter om af te sluiten...")
        reader.ReadString('\n')
}