package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"git.sr.ht/~kiba/gdn/gmi"
	"github.com/gorilla/feeds"
)

func main() {
	baseString := flag.String("b", "", "Base URL. This is required and should be where your gemfeed is located. Ex: gemini://example.org/gemlog/")
	feedLength := flag.Int("n", 0, "Number of most recent items to include in the atom feed. All items from the gemfeed are included by default.")
	feedTitle := flag.String("t", "", "Feed title. Defaults to the first top level heading found in the gemfeed.")
	outPath := flag.String("o", "", "Where to save the converted atom feed. If not provided, prints to stdout.")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 || *baseString == "" {
		fmt.Printf("Usage: %v -b base-url gemfeed\n", os.Args[0])
		flag.PrintDefaults()
		return
	}
	baseURL, err := url.Parse(*baseString)
	if err != nil {
		panic(err)
	}
	gemfeedPath := args[0]

	gemfeedFile, err := os.Open(gemfeedPath)
	if err != nil {
		panic(err)
	}
	defer gemfeedFile.Close()

	feed := &feeds.Feed{
		Title:   *feedTitle,
		Link:    &feeds.Link{Href: *baseString},
		Updated: time.Now(),
	}

	items := make([]*feeds.Item, 0)

	scanner := gmi.NewScanner(gemfeedFile)
	for scanner.Scan() {
		if scanner.Type() == gmi.Head1 && feed.Title == "" {
			feed.Title = scanner.Text()
		}
		if scanner.Type() == gmi.Link {
			desc := strings.SplitN(strings.Trim(scanner.Text(), " "), " ", 2)
			if desc[0] == "" {
				continue
			}
			postDate, err := time.Parse("2006-01-02", desc[0])
			if err != nil {
				continue
			}
			postTitle := desc[1]
			u, err := url.Parse(scanner.URL())
			if err != nil {
				continue
			}
			newItem := &feeds.Item{
				Title:   postTitle,
				Link:    &feeds.Link{Href: baseURL.ResolveReference(u).String()},
				Created: postDate,
			}
			items = append(items, newItem)
		}
	}

	feed.Items = items
	feed.Sort(func(a, b *feeds.Item) bool { return a.Created.After(b.Created) })
	if *feedLength > 0 {
		feed.Items = feed.Items[:*feedLength]
	}
	atom, err := feed.ToAtom()
	if err != nil {
		panic(err)
	}
	if *outPath != "" {
		err := os.WriteFile(*outPath, []byte(atom), 0664)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Print(atom)
	}
}
