# golang-project

This is my first project in Golang. We had 3 things to do:

### Calculator

Make a program that takes for first argument in CLI the operator (+, -, *, /) and numbers separated by a space.
The program then makes the operation as follows:
`go run calc/calc.go + 4 6 2 7` will do `4 + 6 + 2+ 7` and return `19`
**Note :** To use the operator `*`, you have to call it in quotes (`"*"`)

### File Stats

Make a program that takes for argument a file.
It will return its name, its size in bytes, and count which letter is the least and the most recurrent in the file.

### Web scraper

Make a program that takes for argument a file, in which each line corresponds to a URL.
The program will then count the number of `"@"` present in the HTML.
For a list of more than 1 URL, use goroutines to use multi-thread to find emails in all the URLS in concurrency.

**Bonus :** Fetch every email adress present in the HTML thanks to regexp, sort them, and display them
