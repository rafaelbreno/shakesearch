package search

import (
	"bufio"
	"index/suffixarray"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Works struct {
	Data         []byte
	CompleteWork string
	SuffixArray  *suffixarray.Index
	Contents     []Content
	Buf          *bufio.Reader
}

type Content struct {
	Title     string
	LineStart int
	LineEnd   int
}

// Loading File
func Load(filename string) (*Works, error) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	w := Works{
		Data:         data,
		CompleteWork: string(data),
		SuffixArray:  suffixarray.New(data),
	}

	w.GetContents()

	return &w, nil
}

func openFile() {

}
func (w *Works) GetContents() {
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

	file, err := os.Open("completeworks.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	/*
	 * Scanning the first 130 lines because I'd a look
	 * into the file, and the contents/summary ended before
	 * reaching 130
	**/
	for i := 0; i < 130; i++ {
		// Scanning the line by line
		scanner.Scan()

		// Getting the line string
		str := scanner.Text()

		// Matching the regexp with a line
		a := r.FindString(str)
		if a != "" {

			// Appending new content into Content
			w.Contents = append(w.Contents, Content{
				Title: strings.Trim(a, " "),
			})
		}
	}

	for scanner.Scan() {
		str := scanner.Text()

		if str == "" {

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (w *Works) Search(query string) ([]string, error) {
	idxs := w.SuffixArray.Lookup([]byte(query), -1)

	results := []string{}

	for _, idx := range idxs {
		results = append(results, w.CompleteWork[idx:idx])
	}

	return results, nil
}
