package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"text/template"
)

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

var outputFilename *string = flag.String("o", "", "output filename (default: stdout)")

func main() {
	flag.Parse()

	var output io.WriteCloser
	if outputFilename == nil || *outputFilename == "" {
		output = os.Stdout
	} else {
		var err error
		output, err = os.Create(*outputFilename)
		if err != nil {
			fatalf("Can't open output file: %s\n", err)
		}
		defer output.Close()
	}

	templates := template.New("")
	templates.Funcs(map[string]interface{}{
		"for":  For,
		"list": List,
	})

	var err error
	for _, filename := range flag.Args() {
		templates, err = templates.ParseFiles(filename)
		if err != nil {
			fatalf("Can't open template file %q: %s\n", filename, err)
		}
	}

	for _, filename := range flag.Args() {
		err = templates.ExecuteTemplate(output, filename, struct{}{})
		if err != nil {
			fatalf("Can't execute template %q: %s\n", filename, err)
		}
	}
}
