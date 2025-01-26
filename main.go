package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// program for generate links in links.txt
// if you want to specify output file use -o flag
// enter ## in place you want to start sequence of number
// go run main.go -o example.txt
// Enter the link: https://dl6.mementoanime.ir/series/2011/Shameless/S09/1080/Shameless%20S9%20-%20##.%5B1080%5D%5BSS%5D%5BAioFilm.com%5D.mkv
// start at: 1
// end: 12
func main() {

	// output
	file := flag.String("o", "links.txt", "output file")

	flag.Parse()

	var link string
	fmt.Print("Enter the link: ")
	fmt.Scan(&link)

	var start int
	fmt.Print("start at: ")
	fmt.Scan(&start)

	var end int
	fmt.Print("end: ")
	fmt.Scan(&end)

	links := geneartor(link, start, end)

	os.WriteFile(*file, []byte(strings.Join(links, "\n")), 0644)

}

func geneartor(link string, first int, last int) []string {
	var links []string
	for i := first; i <= last; i++ {
		l := strings.ReplaceAll(link, "##", fmt.Sprintf("%02d", i))
		links = append(links, l)
	}
	return links
}
