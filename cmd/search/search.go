package search

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type Works struct {
	Contents    []Content
	CurrentLine int
	Buf         *bufio.Scanner
}

type Content struct {
	Title     string
	LineStart int
	LineEnd   int
}

// Loading File
func Load(filename string) (*Works, error) {
	file, err := os.Open("completeworks.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	buf := bufio.NewScanner(file)

	w := Works{
		Buf:         buf,
		CurrentLine: 0,
	}

	w.SetContents()

	return &w, nil
}

func (w *Works) SetContents() {
	w.getContentTitles()
	w.getContentBody()
}

func (w *Works) getContentTitles() {
	// Counter to store the current line in file

	/*
	 * / {15}[^~]{1,}$/
	 *
	 * / -> regexp init
	 *   {15} -> 15 whitespaces at the beginning of the string
	 *       [^~] -> any string without the char '~'
	 *  		 {1,} -> can have any amount of characters
	 * 				 $ -> End of the line
	 * 				  / -> end of the regexp
	**/
	r := regexp.MustCompile(` {15}[^~]{1,}$`)

	/*
	 * Scanning the first 130 lines because I'd a look
	 * into the file, and the contents/summary ended before
	 * reaching 130
	**/
	for i := 0; i < 130; i++ {
		w.CurrentLine++

		// Scanning the line by line
		w.Buf.Scan()

		// Getting the line string
		str := w.Buf.Text()

		// Matching the regexp with a line
		a := r.FindString(str)
		if a != "" {
			// Appending new content into Content
			w.Contents = append(w.Contents, Content{
				Title: strings.Trim(a, " "),
			})
		}
	}
}

func (w *Works) getContentBody() {
	contents := make(map[string]int, len(w.Contents))

	for key, value := range w.Contents {
		contents[value.Title] = key
	}

	i := 0
	max := len(w.Contents)

	// Itering each line
	for w.Buf.Scan() {

		// Getting current line
		str := strings.Trim(w.Buf.Text(), " ")

		// Only parse non-empty lines
		if str != "" {

			// Checking if current line exists in Content List
			if v, found := contents[str]; found {

				// First content
				if i == 0 {
					w.Contents[v].LineStart = w.CurrentLine
				} else {
					// Defining number of the last line from the previous content
					w.Contents[v-1].LineEnd = w.CurrentLine - 1
					// Defining the start from the current content
					w.Contents[v].LineStart = w.CurrentLine
				}

				i++
			}

			// Stop loop when reaches max number of contents
			if i == max {
				break
			}

			w.CurrentLine++
		}
	}

	if err := w.Buf.Err(); err != nil {
		log.Fatal(err)
	}
}

func (w *Works) Search(query string) ([]string, error) {
	return nil, nil
}
