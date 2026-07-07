package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

)

//inline bold and  italic

var boldPattern = regexp.MustCompile(`\*\*(.+?)\*\*`)
var italicPattern = regexp.MustCompile(`\*(.+?)\*`)

func applyInline(text string) string{
	text = boldPattern.ReplaceAllString(text ,  "<strong>$1</strong>")
    text = italicPattern.ReplaceAllString(text, "<em>$1</em>")
	return text 

}



func convertLines(line string) string{
	if strings.HasPrefix(line, "###"){
		return "<h3>" +applyInline(strings.TrimPrefix(line,"###")) +"</h3>"
	}
	if strings.HasPrefix(line, "##") {
		return "<h2>" + applyInline(strings.TrimPrefix(line, "##")) + "</h2>"
	}
	if strings.HasPrefix(line, "#"){
		return "<h1>" + applyInline(strings.TrimPrefix(line, "#")) +"</h1>"
	}
	return line
}


func main() {  
	if len(os.Args)<2{
		fmt.Println(("usage : md to html <file.md>"))
		return
		}     
	filename := os.Args[1]

	data , err  := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading file : ", err)
		return
	}


	lines := strings.Split(string(data), "\n")

	var output []string
	var paragraph []string
	var list []string

	flush := func(){
		if len(paragraph) >0 {
			text := strings.Join(paragraph, " ")
			output = append(output, "<p>"+applyInline(text)+"</p>")
			paragraph = nil
		}
	}

	flushList := func ()  {
		if len(list)>0{
			output = append(output, "<ul>")
			for _, item := range list{
				output = append(output, "<li>"+applyInline(item)+"</li>")
			}
			output = append(output, "</ui>")
			list = nil
		
			}		
	}
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == ""{
			flush()
			flushList()
			continue
		}
		if strings.HasPrefix(trimmed, "#"){
			flush()
			flushList()
			output = append(output, convertLines(trimmed))
			continue
		}

		if strings.HasPrefix(trimmed, "- "){
			flush()
			list = append(list, strings.TrimPrefix(trimmed ,"- "))
			continue
		}
		flushList() // close any open list before starting/continuing a paragraph
		paragraph = append(paragraph, trimmed)
	}
	flush()
	flushList()

		html := strings.Join(output, "\n")
	fullHTML := "<html><head><title>Converted Markdown</title></head><body>" + html + "</body></html>"

	err = os.WriteFile("output.html", []byte(fullHTML), 0644) // FIXED: was writing "html" (unwrapped fragment) instead of "fullHTML" (the wrapped version) — file and server output didn't match
	if err != nil {
		fmt.Println("error writing file:", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fullHTML)
	})

	fmt.Println("serving at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}