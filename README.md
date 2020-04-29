
# profanity-filter-go
Profanity filter to check if a message contains blacklisted words

# usage
```go
    package main
    
    import (
        "fmt"
    
        filter "github.com/AccelByte/profanity-filter-go"
    )
    
    func main() {
        isBadWord, badWordFound, err := filter.ProfanityCheck("fuck")
        fmt.Println("Bad words found: ", isBadWord)
        fmt.Println("Bad words tripped: ", badWordFound)
        fmt.Println("Error: ", err)
    }

    
```
### Output
```
> go run main.go
Bad words found:  true
Bad words tripped:  [fuck]
Error:  <nil>
```