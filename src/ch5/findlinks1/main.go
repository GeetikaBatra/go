package main

import (
		// "bufio"
		"fmt"
		"os"// read more about what os does


		"golang.org/x/net/html"// for this, you need to get the package `go get golang.org/x/net/html`

)



func main(){
		reader := os.Stdin
		// fmt.Println(text)
	    doc, err := html.Parse(reader)
	    fmt.Println(doc)
	    if err != nil{
	    	fmt.Fprintf(os.Stderr, "findliinks1: %v\n", err)//read about all the formatters once again
	    	os.Exit(1)// check why os.exit 1
	    }
	    for _, link := range visit(nil, doc){
	    	fmt.Println(link)
	    }
}



func visit(links []string, n *html.Node) []string{

	if n.Type == html.ElementNode && n.Data == "a"{
		for _, a := range n.Attr{
			if a.Key == "href"{
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {// check what is the difference between := and =
		links = visit(links, c)
	}

	return links
}