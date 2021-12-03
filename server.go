package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func SplitLines(s string) [][]byte {
	var count int

	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && s[i+1] == 'n' {
			count++
		}
	}
	splitString := []byte(s)
	splitLines := make([][]byte, count+1)

	j := 0

	for i := 0; i < len(splitLines); i++ {
		for j < len(splitString) {
			if splitString[j] == 'n' && splitString[j-1] == '\\' {
				j++
				splitLines[i] = splitLines[i][:len(splitLines[i])-1]
				break
			}
			splitLines[i] = append(splitLines[i], splitString[j])
			j++
		}
	}
	return splitLines
}

func ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userBanner := r.FormValue("font")
	var userString string = r.FormValue("uString")

	if strings.Contains(userString, "\n") {
		userString = strings.Replace(userString, "\r\n", " ", -1)
	}

	splitLines := SplitLines(userString)

	lines, err := ReadLines(userBanner + ".txt")
	if err != nil {
		log.Fatalf("ReadLines: %s", err)
	}

	/*The line below uses the make method to make a map
	and uses a start point of 32 to match up the ascii values
	of each character to the ascii version of the character*/
	charMap := make(map[int][]string)

	start := 32

	for i := 0; i < len(lines); i++ {
		// Tells it to add to start every 9 to match the chars
		if len(charMap[start]) == 9 {
			start++
		}
		charMap[start] = append(charMap[start], lines[i])
	}

	var eString []string

	for j, val := range splitLines {
		for i := 1; i < 9; i++ {
			for k := 0; k < len(val); k++ {
				eString = append(eString, charMap[int(splitLines[j][k])][i])
			}
			eString = append(eString, "\n")
		}
	}

	sAscii := strings.Join(eString, "")
	fmt.Fprintf(w, sAscii)
	// fmt.Fprintf(w, userBanner)

	d := struct {
		Banner string
		String string
		sAscii string
	}{
		Banner: userBanner,
		String: userString,
		sAscii: sAscii,
	}

	tpl.ExecuteTemplate(w, "ascii.gohtml", d)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ascii", ascii)

	http.ListenAndServe(":8080", nil)
}
