package main

import(
	"flag"
	"log"
	"os"

	_ "github.com/jmgrimes/feed-search/matchers"
	"github.com/jmgrimes/feed-search/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var searchTerm string
	flag.StringVar(&searchTerm, "term", "", "the term to search for in the feeds")
	flag.Parse()
	if searchTerm == "" {
		log.Fatal("search term must be specified with -term flag")
	}
	search.Run(searchTerm)
}