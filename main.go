package main          

import (
    "fmt"             
    "os"    
	"strings"  
	"regexp"        
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
		return "<h3>" +strings.TrimPrefix(line,"###") + "</h3>"
	}
	if strings.HasPrefix(line, "##") {
		return "<h2>" + strings.TrimPrefix(line, "##") + "</h2>"
	}
	if strings.HasPrefix(line, "#"){
		return "<h1>" + strings.TrimPrefix(line, "#") +"</h1>"
	}
	return line
}


func main() {         
	data , err  := os.ReadFile("test.md")
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
			output = append(output, "<p>"+text+"</p>")
			paragraph = nil
		}
	}

	flushList := func ()  {
		if len(list)>0{
			output = append(output, "<u1>")
			for _, item := range list{
				output = append(output, "<li>"+applyInline(item)+"</l1>")
			}
			output = append(output, "</u1>")
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
			output = append(output, convertLines((trimmed)))
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
	err = os.WriteFile("output.html", []byte(html), 0644)
	if err != nil {
		fmt.Println("error writing file:", err)
		return
	}
	fmt.Println("wrote output.html")
}