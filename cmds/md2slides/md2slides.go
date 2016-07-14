//
// md2slides.go - A simple command line utility that uses Markdown
// to generate a sequence of HTML5 pages that can be used for presentations.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	// 3rd Part packages
	"github.com/russross/blackfriday"
)

const (
	version = "v0.0.2"
	license = `
%s

 Copyright (c) 2016, Caltech
 All rights not granted herein are expressly reserved by Caltech.
 
 Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
 
 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 
 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 
 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
 
 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 
`
)

// Slide is the metadata about a slide to be generated.
type Slide struct {
	CurNo   int
	PrevNo  int
	NextNo  int
	FirstNo int
	LastNo  int
	FName   string
	Title   string
	Content string
	CSSPath string
}

var (
	showHelp          bool
	showVersion       bool
	showLicense       bool
	presentationTitle string
	defaultHTML       = `<!DOCTYPE html>
<html>
<head>
   {{- if .Title -}}<title>{{ .Title }}</title>{{- end}}
   {{- if .CSSPath -}}
   <link href="{{ .CSSPath }}" rel="stylesheet" />
   {{else}}
   <style>
body {
	max-width: 100%;
	max-height: 100%;
	margin: 10%;
	padding: 0;
	font-size: 24px;
	font-family: sans-serif;
}

ul {
	list-style: circle;
	text-indent: 0.25em;
}

nav {
	position: absolute;
   	bottom: 1em; 
	margin:0;
	padding:0.24em;
	width: 100%;
	height: 4em;
	text-align: left;
	font-size: 60%;
}

section, p {
	max-width: 85%;
	padding: 0.24em;
	margin: 0.24em;
}
   </style>
   {{- end }}
</head>
<body>
	<section>{{ .Content }}</section>
	<nav>
{{ if ne .CurNo .FirstNo -}}
<a href="{{printf "%02d-%s.html" .FirstNo .FName}}">Home</a>
{{- end}}
{{ if gt .CurNo .FirstNo -}} 
<a href="{{printf "%02d-%s.html" .PrevNo .FName}}">Prev</a>
{{- end}}
{{ if lt .CurNo .LastNo -}} 
<a href="{{printf "%02d-%s.html" .NextNo .FName}}">Next</a>
{{- end}}
	</nav>
</body>
</html>
`
	cssPath       string
	templateFName string
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.StringVar(&presentationTitle, "title", "", "Presentation title")
	flag.StringVar(&cssPath, "css", cssPath, "Specify the CSS file to use")
	flag.StringVar(&templateFName, "template",
		templateFName, "Specify an HTML template to use")
}

func makeSlide(tmpl *template.Template, slide *Slide) {
	sname := fmt.Sprintf(`%02d-%s.html`, slide.CurNo, slide.FName)
	fp, err := os.Create(sname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s\n", sname, err)
		os.Exit(1)
	}
	defer fp.Close()
	err = tmpl.Execute(fp, slide)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s", sname, err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s\n", sname)
}

func main() {
	appname := path.Base(os.Args[0])

	flag.Parse()

	if showHelp == true {
		fmt.Printf(`
 USAGE: %s [OPTIONS] [FILENAME]

 OPTIONS:

`, appname)

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("\t-%s\t(defaults to %s) %s\n", f.Name, f.Value, f.Usage)
		})

		fmt.Printf("\n\n Version %s\n", version)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Printf("\n Version %s\n", version)
		os.Exit(0)
	}
	if showLicense == true {
		fmt.Printf(license, appname)
		os.Exit(0)
	}

	if templateFName != "" {
		src, err := ioutil.ReadFile(templateFName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s %s\n", templateFName, err)
			os.Exit(1)
		}
		defaultHTML = string(src)
	}

	var fname string
	args := flag.Args()
	if len(args) > 0 {
		//fname, args = args[0], args[1:]
		fname = args[0]
	}
	if fname == "" {
		fmt.Fprintf(os.Stderr, "Missing filename")
		os.Exit(1)
	}
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fname = strings.TrimSuffix(path.Base(fname), path.Ext(fname))
	tmpl, err := template.New("slide").Parse(defaultHTML)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Note: handle legacy CR/LF endings as well as normal LF line endings
	if bytes.Contains(src, []byte("\r\n")) {
		src = bytes.Replace(src, []byte("\r\n"), []byte("\n"), -1)
	}
	slides := bytes.Split(src, []byte("--\n"))

	fmt.Printf("Slide count: %d\n", len(slides))
	lastSlide := len(slides) - 1
	for i, s := range slides {
		data := blackfriday.MarkdownCommon(s)
		makeSlide(tmpl, &Slide{
			FName:   fname,
			CurNo:   i,
			PrevNo:  (i - 1),
			NextNo:  (i + 1),
			FirstNo: 0,
			LastNo:  lastSlide,
			Title:   presentationTitle,
			Content: string(data),
			CSSPath: cssPath,
		})
	}
}
