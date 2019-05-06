package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Song struct {
	Title  string
	Author string
	Year   int
	Length time.Duration
}

type byTitle []*Song

func (s byTitle) Len() int {
	return len(s)
}

func (s byTitle) Less(i, j int) bool {
	return s[i].Title > s[j].Title
}

func (s byTitle) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type customSort struct {
	songs []*Song
	less  func(x, y *Song) bool
}

func (c customSort) Len() int {
	return len(c.songs)
}

func (c customSort) Less(i, j int) bool {
	return c.less(c.songs[i], c.songs[j])
}

func (c customSort) Swap(i, j int) {
	c.songs[i], c.songs[j] = c.songs[j], c.songs[i]
}

func main() {
	var songs = []*Song{
		{"Give me love", "wuzhc", 2012, length("3m38s")},
		{"Hello word", "mayun", 2019, length("3m45s")},
		{"Hello word", "tom", 2017, length("6m5s")},
	}

	fmt.Println("byTitle:")
	sort.Sort(byTitle(songs))
	printSongs(songs)

	fmt.Println("byTitle reverse:")
	sort.Sort(sort.Reverse(byTitle(songs)))
	printSongs(songs)

	fmt.Println("customSort:")
	sort.Sort(customSort{songs, func(i, j *Song) bool {
		// 按排序规则,title>author>year>length
		if i.Title != j.Title {
			return i.Title > j.Title
		}
		if i.Author != j.Author {
			return i.Author > j.Author
		}
		if i.Year != j.Year {
			return i.Year > j.Year
		}
		if i.Length != j.Length {
			return i.Length > j.Length
		}
		return false
	}})
	printSongs(songs)
}

func length(t string) time.Duration {
	d, err := time.ParseDuration(t)
	if err != nil {
		panic(t)
	}
	return d
}

func printSongs(songs []*Song) {
	const format = "%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Author", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "----", "------")
	for _, v := range songs {
		fmt.Fprintf(tw, format, v.Title, v.Author, v.Year, v.Length)
	}
	tw.Flush()
	fmt.Println("\n")
}
