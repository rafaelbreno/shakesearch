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
	w.getContentTitles().GetContentBody()
}

func (w *Works) getContentTitles() *Works {
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

	return w
}

func (w *Works) GetContentBody() {
	i := 0
	max := len(w.Contents)

	for w.Buf.Scan() {

		str := w.Buf.Text()

		if str == w.Contents[i].Title {

			// First content
			if i == 0 {
				w.Contents[i].LineStart = w.CurrentLine
			} else {
				// Defining number of the last line from the previous content
				w.Contents[i-1].LineEnd = w.CurrentLine - 1
				// Defining the start from the current content
				w.Contents[i].LineStart = w.CurrentLine
			}

			i++
		}

		if i == max {
			break
		}

		w.CurrentLine++
	}

	if err := w.Buf.Err(); err != nil {
		log.Fatal(err)
	}
}

func (w *Works) Search(query string) ([]string, error) {
	return nil, nil
}
