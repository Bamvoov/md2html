package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

// Package-level compiled regular expressions and sets for efficiency and clean code.
var (
	// Embed link pattern: ![[target]] or ![[target|display]]
	embedRegex = regexp.MustCompile(`!\[\[([^\]|]+)(?:\|([^\]]+))?\]\]`)
	// Wikilink pattern: [[target]] or [[target|display]]
	wikilinkRegex = regexp.MustCompile(`\[\[([^\]|]+)(?:\|([^\]]+))?\]\]`)
	// Bold pattern: **bold text**
	boldRegex = regexp.MustCompile(`\*\*(.+?)\*\*`)
	// Italic pattern: *italic text*
	italicRegex = regexp.MustCompile(`\*(.+?)\*`)

	// Set of supported image file extensions (lowercase) for Obsidian image embeds.
	imageExtensions = map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
		".webp": true,
		".svg":  true,
	}
)

// Frontmatter holds the metadata parsed from the optional YAML frontmatter header.
type Frontmatter struct {
	Title string
	Date  string
}

// LineType represents the structural classification of a single Markdown line.
type LineType int

const (
	LineParagraph LineType = iota
	LineHeader
	LineCheckbox
	LinePlainList
	LineBlank
)

// ParsedLine stores the parsed attributes of a line.
type ParsedLine struct {
	Type    LineType
	Text    string
	Level   int  // Used for headers (1 = <h1>, 2 = <h2>, 3 = <h3>)
	Checked bool // Used for checkbox list items (- [ ] vs - [x])
}

// titleCase converts words in a string to Title Case (e.g. "linear algebra" -> "Linear Algebra").
func titleCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			r := []rune(word)
			r[0] = unicode.ToUpper(r[0])
			words[i] = string(r)
		}
	}
	return strings.Join(words, " ")
}

// slugify converts a string target into a clean, lowercase dash-separated filename.
// For example: "Some Note" -> "some-note".
func slugify(s string) string {
	lowered := strings.ToLower(strings.TrimSpace(s))
	words := strings.Fields(lowered)
	return strings.Join(words, "-")
}

// extractFrontmatter checks if the file starts with "---". If so, it looks for a closing "---".
// If a valid frontmatter block is found, it extracts "title" and "date" key-value pairs and
// returns the parsed metadata along with the remaining content.
// If there is no closing "---", the whole file is treated as having no frontmatter (returning empty Frontmatter).
func extractFrontmatter(content string) (Frontmatter, string) {
	var fm Frontmatter
	if !strings.HasPrefix(content, "---") {
		return fm, content
	}

	normalized := strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(normalized, "\n")

	// Ensure the first line is exactly "---"
	if strings.TrimSpace(lines[0]) != "---" {
		return fm, content
	}

	// Look for the closing "---" delimiter
	closingIndex := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			closingIndex = i
			break
		}
	}

	// If no closing "---" line is found, do not consume frontmatter
	if closingIndex == -1 {
		return fm, content
	}

	// Parse key: value lines between the delimiters
	for i := 1; i < closingIndex; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.ToLower(strings.TrimSpace(parts[0]))
			val := strings.TrimSpace(parts[1])
			val = strings.Trim(val, `"'`) // Strip optional quotes surrounding values
			if key == "title" {
				fm.Title = val
			} else if key == "date" {
				fm.Date = val
			}
		}
	}

	// Remaining content starts on the line after the closing "---"
	remaining := strings.Join(lines[closingIndex+1:], "\n")
	return fm, remaining
}

// resolveTitle computes the HTML title following priority order:
// 1. "title" from frontmatter, if present.
// 2. Otherwise, the filename (without extension, dashes/underscores replaced with spaces, title-cased).
func resolveTitle(fm Frontmatter, srcPath string) string {
	if fm.Title != "" {
		return fm.Title
	}

	base := filepath.Base(srcPath)
	ext := filepath.Ext(base)
	nameWithoutExt := strings.TrimSuffix(base, ext)

	replaced := strings.ReplaceAll(nameWithoutExt, "-", " ")
	replaced = strings.ReplaceAll(replaced, "_", " ")

	return titleCase(replaced)
}

// applyObsidianLinks handles Obsidian-flavored embeds (![[...]]) and wikilinks ([[...]]).
// Non-obvious design detail: Embeds (![[...]]) MUST be matched and replaced BEFORE plain wikilinks ([[...]])
// because every embed syntax (![[target]]) contains a wikilink pattern ([[target]]) with a leading exclamation mark.
func applyObsidianLinks(text string) string {
	// 1. Process embeds: ![[target]] or ![[target|display]]
	text = embedRegex.ReplaceAllStringFunc(text, func(match string) string {
		submatches := embedRegex.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		target := strings.TrimSpace(submatches[1])
		display := ""
		if len(submatches) >= 3 {
			display = strings.TrimSpace(submatches[2])
		}

		ext := strings.ToLower(filepath.Ext(target))
		if imageExtensions[ext] {
			// Image embed -> <img src="target" alt="target">
			return fmt.Sprintf(`<img src="%s" alt="%s">`, target, target)
		}

		// Non-image embed -> convert target to slugified html link
		linkText := display
		if linkText == "" {
			linkText = target
		}
		href := slugify(target) + ".html"
		return fmt.Sprintf(`<a href="%s">%s</a>`, href, linkText)
	})

	// 2. Process plain wikilinks: [[target]] or [[target|display]]
	text = wikilinkRegex.ReplaceAllStringFunc(text, func(match string) string {
		submatches := wikilinkRegex.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}
		target := strings.TrimSpace(submatches[1])
		display := ""
		if len(submatches) >= 3 {
			display = strings.TrimSpace(submatches[2])
		}

		linkText := display
		if linkText == "" {
			linkText = target
		}
		href := slugify(target) + ".html"
		return fmt.Sprintf(`<a href="%s">%s</a>`, href, linkText)
	})

	return text
}

// applyInline applies inline formatting (Obsidian links, bold, italic) inside headers, paragraphs, and list items.
// Non-obvious design detail: Bold (**bold**) MUST be checked before italic (*italic*) so the double asterisk **
// is not mistakenly matched and mangled as a single asterisk italic syntax.
func applyInline(text string) string {
	// 1. Obsidian wikilinks and embeds first
	text = applyObsidianLinks(text)

	// 2. Bold (**bold**) before italic (*italic*)
	text = boldRegex.ReplaceAllString(text, "<strong>$1</strong>")
	text = italicRegex.ReplaceAllString(text, "<em>$1</em>")

	return text
}

// isCheckbox checks if a trimmed line is a checkbox list item (- [ ] or - [x] or - [X]).
func isCheckbox(s string) (text string, checked bool, ok bool) {
	if strings.HasPrefix(s, "- [ ] ") {
		return strings.TrimPrefix(s, "- [ ] "), false, true
	}
	if s == "- [ ]" {
		return "", false, true
	}
	if strings.HasPrefix(s, "- [x] ") {
		return strings.TrimPrefix(s, "- [x] "), true, true
	}
	if s == "- [x]" {
		return "", true, true
	}
	if strings.HasPrefix(s, "- [X] ") {
		return strings.TrimPrefix(s, "- [X] "), true, true
	}
	if s == "- [X]" {
		return "", true, true
	}
	return "", false, false
}

// convertLine classifies a single line according to strict precedence order:
// 1. Checkboxes (- [ ] or - [x]) MUST be checked before plain list items (- text) because "- [ ]" starts with "- ".
// 2. Plain unordered list items (- text).
// 3. Headers (# , ## , ### ).
// 4. Anything else is paragraph text or blank.
func convertLine(line string) ParsedLine {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return ParsedLine{Type: LineBlank}
	}

	// 1. Checkboxes first
	if text, checked, ok := isCheckbox(trimmed); ok {
		return ParsedLine{Type: LineCheckbox, Text: text, Checked: checked}
	}

	// 2. Plain list items second
	if strings.HasPrefix(trimmed, "- ") {
		return ParsedLine{Type: LinePlainList, Text: strings.TrimPrefix(trimmed, "- ")}
	}

	// 3. Headers third
	if strings.HasPrefix(trimmed, "### ") {
		return ParsedLine{Type: LineHeader, Level: 3, Text: strings.TrimPrefix(trimmed, "### ")}
	}
	if strings.HasPrefix(trimmed, "## ") {
		return ParsedLine{Type: LineHeader, Level: 2, Text: strings.TrimPrefix(trimmed, "## ")}
	}
	if strings.HasPrefix(trimmed, "# ") {
		return ParsedLine{Type: LineHeader, Level: 1, Text: strings.TrimPrefix(trimmed, "# ")}
	}

	// 4. Paragraph lines
	return ParsedLine{Type: LineParagraph, Text: line}
}

// convertChecklistItem formats a checkbox list item into HTML:
// <li><input type="checkbox" disabled [checked]> text</li>
func convertChecklistItem(text string, checked bool) string {
	checkedAttr := ""
	if checked {
		checkedAttr = " checked"
	}
	return fmt.Sprintf("<li><input type=\"checkbox\" disabled%s> %s</li>", checkedAttr, applyInline(text))
}

// convertMarkdown parses Markdown body text into HTML body content.
// Consecutive paragraph lines join with a space into <p>...</p>.
// Consecutive list items (checkboxes or plain) accumulate into the same <ul>...</ul> block.
func convertMarkdown(markdown string) string {
	normalized := strings.ReplaceAll(markdown, "\r\n", "\n")
	lines := strings.Split(normalized, "\n")

	var sb strings.Builder
	var pLines []string
	inList := false

	flushParagraph := func() {
		if len(pLines) > 0 {
			joined := strings.Join(pLines, " ")
			sb.WriteString(fmt.Sprintf("<p>%s</p>\n", applyInline(joined)))
			pLines = nil
		}
	}

	flushList := func() {
		if inList {
			sb.WriteString("</ul>\n")
			inList = false
		}
	}

	for _, rawLine := range lines {
		parsed := convertLine(rawLine)

		switch parsed.Type {
		case LineBlank:
			flushParagraph()
			flushList()

		case LineHeader:
			flushParagraph()
			flushList()
			sb.WriteString(fmt.Sprintf("<h%d>%s</h%d>\n", parsed.Level, applyInline(parsed.Text), parsed.Level))

		case LineCheckbox, LinePlainList:
			flushParagraph()
			if !inList {
				sb.WriteString("<ul>\n")
				inList = true
			}
			if parsed.Type == LineCheckbox {
				sb.WriteString(convertChecklistItem(parsed.Text, parsed.Checked) + "\n")
			} else {
				sb.WriteString(fmt.Sprintf("<li>%s</li>\n", applyInline(parsed.Text)))
			}

		case LineParagraph:
			flushList()
			pLines = append(pLines, strings.TrimSpace(parsed.Text))
		}
	}

	flushParagraph()
	flushList()

	return sb.String()
}

// buildHTMLDocument constructs the full HTML document string containing <!DOCTYPE html>, <head>, and <body>.
func buildHTMLDocument(title string, date string, bodyContent string) string {
	var sb strings.Builder
	sb.WriteString("<!DOCTYPE html>\n")
	sb.WriteString("<html lang=\"en\">\n")
	sb.WriteString("<head>\n")
	sb.WriteString("    <meta charset=\"UTF-8\">\n")
	sb.WriteString(fmt.Sprintf("    <title>%s</title>\n", title))
	sb.WriteString("    <link rel=\"stylesheet\" href=\"style.css\">\n")
	sb.WriteString("</head>\n")
	sb.WriteString("<body>\n")
	sb.WriteString(fmt.Sprintf("    <h1>%s</h1>\n", applyInline(title)))
	if date != "" {
		sb.WriteString(fmt.Sprintf("    <p><em>%s</em></p>\n", applyInline(date)))
	}
	sb.WriteString(bodyContent)
	sb.WriteString("</body>\n")
	sb.WriteString("</html>\n")
	return sb.String()
}

// convertFile handles converting a single Markdown source file to HTML and saving it to destDir.
func convertFile(srcPath string, outDir string) (string, error) {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return "", fmt.Errorf("reading source file: %w", err)
	}

	fm, bodyMarkdown := extractFrontmatter(string(data))
	title := resolveTitle(fm, srcPath)

	bodyHTML := convertMarkdown(bodyMarkdown)
	fullDoc := buildHTMLDocument(title, fm.Date, bodyHTML)

	// Determine output directory and output filename
	baseName := filepath.Base(srcPath)
	nameWithoutExt := strings.TrimSuffix(baseName, filepath.Ext(baseName))
	outFileName := nameWithoutExt + ".html"

	var destDir string
	if outDir != "" {
		destDir = outDir
	} else {
		destDir = filepath.Dir(srcPath)
	}

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return "", fmt.Errorf("creating output directory: %w", err)
	}

	outPath := filepath.Join(destDir, outFileName)
	if err := os.WriteFile(outPath, []byte(fullDoc), 0644); err != nil {
		return "", fmt.Errorf("writing output file: %w", err)
	}

	return outPath, nil
}

// ensureStyleCSS checks if style.css exists in targetDir. If missing, it writes a default clean blog stylesheet.
// It never overwrites an existing style.css file to preserve user customizations.
func ensureStyleCSS(targetDir string) error {
	if targetDir == "" {
		targetDir = "."
	}
	cssPath := filepath.Join(targetDir, "style.css")

	// Check if style.css already exists
	if _, err := os.Stat(cssPath); err == nil {
		return nil
	}

	const defaultCSS = `/* Clean, readable typography for Obsidian converted blog posts */
body {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
    max-width: 760px;
    margin: 40px auto;
    padding: 0 20px;
    line-height: 1.6;
    color: #222222;
    background-color: #ffffff;
}

h1, h2, h3 {
    line-height: 1.25;
    color: #111111;
    margin-top: 1.5em;
    margin-bottom: 0.5em;
}

h1 {
    font-size: 2.2rem;
    border-bottom: 1px solid #eaeaea;
    padding-bottom: 0.3em;
}

h2 {
    font-size: 1.6rem;
}

h3 {
    font-size: 1.25rem;
}

p {
    margin-bottom: 1.2em;
}

a {
    color: #0066cc;
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}

img {
    max-width: 100%;
    height: auto;
    border-radius: 4px;
    margin: 1em 0;
}

ul {
    padding-left: 24px;
    margin-bottom: 1.2em;
}

li {
    margin-bottom: 0.4em;
}

/* Disabled checkbox input styling for list items */
input[type="checkbox"] {
    margin-right: 8px;
    vertical-align: middle;
    cursor: default;
}
`
	return os.WriteFile(cssPath, []byte(defaultCSS), 0644)
}

// main is the CLI entry point. It parses flags, checks arguments, and orchestrates
// single-file mode, batch-directory mode conversion, and automatic local web server hosting.
func main() {
	dirFlag := flag.String("dir", "", "Directory containing .md files to convert in batch mode")
	outFlag := flag.String("out", "", "Output folder for generated HTML and CSS files")
	noServerFlag := flag.Bool("no-serve", false, "Disable starting the local HTTP web server")
	portFlag := flag.Int("port", 8080, "Port for the local HTTP web server (default 8080)")

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  Single file mode: mdtohtml [flags] <file.md>")
		fmt.Println("  Batch mode:       mdtohtml [flags] -dir <folder>")
		fmt.Println()
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}

	flag.Parse()

	// If no file argument and no -dir flag are given, print usage message and exit cleanly (status 0).
	if flag.NArg() == 0 && *dirFlag == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Ensure style.css exists in output folder (if -out specified) or current directory
	cssDir := "."
	if *outFlag != "" {
		cssDir = *outFlag
	}
	if err := os.MkdirAll(cssDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "error creating directory %s: %v\n", cssDir, err)
		os.Exit(1)
	}
	if err := ensureStyleCSS(cssDir); err != nil {
		fmt.Fprintf(os.Stderr, "error writing style.css: %v\n", err)
	}

	var convertedFiles []string

	// Batch Mode: process directory
	if *dirFlag != "" {
		entries, err := os.ReadDir(*dirFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading directory %s: %v\n", *dirFlag, err)
			os.Exit(1)
		}

		convertedCount := 0
		for _, entry := range entries {
			if entry.IsDir() || strings.ToLower(filepath.Ext(entry.Name())) != ".md" {
				continue
			}

			srcPath := filepath.Join(*dirFlag, entry.Name())
			outPath, err := convertFile(srcPath, *outFlag)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error converting %s: %v\n", srcPath, err)
				continue
			}

			fmt.Printf("converted: %s -> %s\n", srcPath, outPath)
			convertedFiles = append(convertedFiles, outPath)
			convertedCount++
		}

		fmt.Printf("converted %d files\n", convertedCount)
	} else if flag.NArg() > 0 {
		// Single File Mode: process positional file argument
		srcPath := flag.Arg(0)
		outPath, err := convertFile(srcPath, *outFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("converted: %s -> %s\n", srcPath, outPath)
		convertedFiles = append(convertedFiles, outPath)
	}

	// Unless -no-serve is explicitly passed, start local HTTP web server and print full URL links
	if !*noServerFlag {
		serveDir := cssDir
		port := *portFlag

		fmt.Println()
		fmt.Println("Local Web Server running!")

		if len(convertedFiles) == 1 {
			fileName := filepath.Base(convertedFiles[0])
			escapedFileName := url.PathEscape(fileName)
			fmt.Printf("  Preview your page at: http://localhost:%d/%s\n", port, escapedFileName)
		} else {
			fmt.Printf("  Preview directory at: http://localhost:%d/\n", port)
		}

		fmt.Println("Press Ctrl+C to stop the server.")

		addr := fmt.Sprintf(":%d", port)
		if err := http.ListenAndServe(addr, http.FileServer(http.Dir(serveDir))); err != nil {
			fmt.Fprintf(os.Stderr, "server error: %v\n", err)
			os.Exit(1)
		}
	}
}