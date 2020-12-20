package search

import (
	"bufio"
	"errors"
	"fmt"
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
	Body      []string
}

var currentLine int

func (w *Works) GetContentByKeyword(key string) (Content, error) {
	for i := range w.Contents {
		if strings.Contains(strings.ToLower(w.Contents[i].Title), strings.ToLower(key)) {
			return w.Contents[i], nil
		}
		for _, value := range strings.Split(key, " ") {
			if strings.Contains(strings.ToLower(w.Contents[i].Title), strings.ToLower(value)) {
				return w.Contents[i], nil
			}
		}
	}
	return Content{}, errors.New(fmt.Sprintf(`Not found any: %s`, key))
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
		currentLine++

		// Scanning the line by line
		w.Buf.Scan()

		// Getting the line string
		str := w.Buf.Text()

		// Matching the regexp with a line
		if a := r.FindString(str); a != "" {
			// Appending new content into Content
			w.Contents = append(w.Contents, Content{
				Title: strings.Trim(a, " "),
			})
		}
	}
}

func checkTitles(cont map[string]int, str string) (int, bool, string) {
	if v, found := cont[str]; found {
		return v, found, str
	}

	if v, found := cont["THE TRAGEDY OF "+str]; found {
		return v, found, "THE TRAGEDY OF " + str
	}

	return -1, false, ""
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
		currentLine++

		// Getting current line
		line := w.Buf.Text()

		// Saving the body
		if w.Contents[0].LineStart != 0 {
			w.Contents[i-1].Body = append(w.Contents[i-1].Body, line)
		}

		if line != "" {
			if line[0] == ' ' && i != max {
				continue
			}
		}

		str := strings.Trim(line, " ")
		str = strings.Trim(str, ":")

		if a := regexp.MustCompile(`[A-Z\s\'\\’;]{1,}$`).FindString(str); a == "" {
			continue
		}

		// Only parse non-empty lines
		if a := regexp.MustCompile(`[^.]{1,}$`).FindString(str); str == "" && a == "" {
			continue
		}

		// Appending each line to the content

		// Conditional to avoid checkTitles()
		if i != max {
			// Checking if current line exists in Content List
			if v, found, key := checkTitles(contents, str); found {

				// First content
				if i == 0 {
					w.Contents[v].LineStart = currentLine
				} else {
					// Defining number of the last line from the previous content
					w.Contents[v-1].LineEnd = currentLine - 1
					// Defining the start from the current content
					w.Contents[v].LineStart = currentLine

					w.Contents[i-1].Body = w.Contents[i-1].Body[:len(w.Contents[i-1].Body)-1]
				}

				delete(contents, key)
				i++
			}
		}

		// Stop loop when reaches max number of contents
		if str == "FINIS" && i == max {
			w.Contents[i-1].LineEnd = currentLine
			break
		}

	}

	if err := w.Buf.Err(); err != nil {
		log.Fatal(err)
	}
}
