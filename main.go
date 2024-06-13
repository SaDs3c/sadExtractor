package main

import (
    "bufio"
    "flag"
    "fmt"
    "net/http"
    "os"
    "io"
    "golang.org/x/net/html"
    "strings"
)

const version = "sadExtractor v1.0.0"

func main() {
    // Define flags
    singleDomain := flag.String("d", "", "URL of the single domain to parse")
    urlListFile := flag.String("l", "", "Text file containing a list of URLs to parse")
    outputFile := flag.String("o", "", "Output file to save the extracted links")
    showVersion := flag.Bool("v", false, "Show version information")
    
    // Custom usage message
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [-d <url> | -l <file>] [-o <output>] [-v]\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "\nFlags:\n")
        flag.PrintDefaults()
        fmt.Fprintf(os.Stderr, "\nExamples:\n")
        fmt.Fprintf(os.Stderr, "  Single URL: %s -d https://example.com -o output.txt\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "  URL List File: %s -l urls.txt -o output.txt\n", os.Args[0])
    }
    
    // Parse the flags
    flag.Parse()
    
    // Handle the version flag
    if *showVersion {
        fmt.Println(version)
        return
    }

    // Check if both or none of the URL flags are provided
    if (*singleDomain == "" && *urlListFile == "") || (*singleDomain != "" && *urlListFile != "") {
        flag.Usage()
        return
    }

    // Open the output file if specified
    var output *os.File
    var err error
    if *outputFile != "" {
        output, err = os.Create(*outputFile)
        if err != nil {
            fmt.Println("Error creating output file:", err)
            return
        }
        defer output.Close()
    }

    if *singleDomain != "" {
        // Fetch and parse a single URL
        processSingleURL(*singleDomain, output)
    } else if *urlListFile != "" {
        // Fetch and parse a list of URLs from a text file
        processURLList(*urlListFile, output)
    }
}

// processSingleURL processes a single URL
func processSingleURL(url string, output *os.File) {
    body, err := fetchWebPage(url)
    if err != nil {
        fmt.Println("Error fetching the URL:", err)
        return
    }
    defer body.Close()

    doc, err := html.Parse(body)
    if err != nil {
        fmt.Println("Error parsing the HTML:", err)
        return
    }

    links := extractLinks(nil, doc)
    for _, link := range links {
        if output != nil {
            fmt.Fprintln(output, link)
        } else {
            fmt.Println(link)
        }
    }
}

// processURLList processes a list of URLs from a text file
func processURLList(filename string, output *os.File) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        url := scanner.Text()
        processSingleURL(url, output)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}

// fetchWebPage fetches the content of the given URL and returns the response body
func fetchWebPage(url string) (body io.ReadCloser, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("failed to fetch URL: %s, status code: %d", url, resp.StatusCode)
    }
    return resp.Body, nil
}

// extractLinks recursively traverses the HTML nodes and extracts all href links
func extractLinks(links []string, n *html.Node) []string {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, attr := range n.Attr {
            if attr.Key == "href" {
                link := attr.Val
                if strings.HasPrefix(link, "http") {
                    links = append(links, link)
                }
                break
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = extractLinks(links, c)
    }
    return links
}
