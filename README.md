
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

### Generates list of profanity list from json file to array in go language code
1. Run file main.go under directory `/profanity-filter-go/profanitylist/generator/main.go`
   ```go run main.go```
2. This script unmarshal the json files and generates as array variable in golang format.
   This script diplays the result in terminal.
   Example result:
    ```go
            var frenchCAProfanityList = profanityList{
                    Country:     "frenchCA",
                    BadWordList: []string{"noune", "osti", "criss", "crisse", "calice", "tabarnak", "viarge"},
            }
    
            var klingonProfanityList = profanityList{
                    Country:     "klingon",
                    BadWordList: []string{"ghuy'cha'", "QI'yaH", "Qu'vatlh"},
            }
    
    ```
3. Copy the result as array variable profanity list in `profanity-filter-go/profanitylist.go`