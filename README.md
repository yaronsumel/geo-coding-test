
### 'GPS' coding challenge by Yaron Sumel

the challenge is to build go program that's reads the csv file and returns the five closest places and five Furthest places to some target point.

## get it 

 `$ go get github.com/yaronsumel/geo-coding-test`

### run it

 `$ go run main.go`
 
##### flags

```
  -file string
        path to csv file (default "data.csv")
```

### expected output

    2017/05/17 01:05:37 Closest Places
    2017/05/17 01:05:37 Place ID 442406 is 0.3338381556847423 Km away
    2017/05/17 01:05:37 Place ID 285782 is 0.5280320575225937 Km away
    2017/05/17 01:05:37 Place ID 429151 is 0.6480104060043963 Km away
    2017/05/17 01:05:38 Place ID 512818 is 0.7405525423045783 Km away
    2017/05/17 01:05:38 Place ID 25182 is 0.8216419667395297 Km away
    2017/05/17 01:05:38 Furthest Places
    2017/05/17 01:05:38 Place ID 7818 is 8776.646278220525 Km away
    2017/05/17 01:05:38 Place ID 382013 is 1810.117975607402 Km away
    2017/05/17 01:05:38 Place ID 381823 is 1758.8482704875876 Km away
    2017/05/17 01:05:38 Place ID 382582 is 1758.0806131200134 Km away
    2017/05/17 01:05:38 Place ID 382693 is 1441.7191536957162 Km away
    2017/05/17 01:05:38 Done!

### cover 
 
 `$ go test -cover ./...`

    ok      github.com/yaronsumel/geo-coding-test/src/dataHandling                  0.029s  coverage: 100.0% of statements
    ok      github.com/yaronsumel/geo-coding-test/src/dataHandling/handlers/csv     0.034s  coverage: 100.0% of statements
    ok      github.com/yaronsumel/geo-coding-test/src/list                          0.026s  coverage: 100.0% of statements
    ok      github.com/yaronsumel/geo-coding-test/src/place                         0.025s  coverage: 100.0% of statements

> if my .csv is very big .. how this app will perform ?

well.. actually the same, many app will fail with reading all into var, 
here its reading line by line.
the list of places that the app holds is limited to size of 10. any append on top of that will pop the middle item after sort.. that's will make us list with Top5 and Bottom5.

> would it be possible to change the app to work with database ? 

yes, really easy. just implement the dataHandling interface which is 

```go
    // Interface - data handler interface
    type Interface interface {
        // Next - read next place into &p
        // return error at EOF
        Next(p *place.Place) bool
        // Close
        // close all related resources to the handler
        Close() error
    }
```

and make sure to register your handler like that

```go
    	// register handler
    	dataHandling.RegHandler("csv", csvHandler)
```     
