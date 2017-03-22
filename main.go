package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"golang.org/x/tools/imports"

	"github.com/arstd/light/domain"
	"github.com/arstd/light/parse"
	"github.com/arstd/light/prepare"
	"github.com/arstd/light/util"
	"github.com/arstd/log"
)

func usage() {
	fmt.Fprintln(os.Stderr, `usage: light [flags] [file.go]
	//go:generate light [flags] [file.go]
examples:
	light -force -dbvar=db.DB -dbpath=github.com/arstd/light/example/mapper
	light -force -dbvar=db2.DB -dbpath=github.com/arstd/light/example/mapper`)
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	// log.SetLevel(log.Lwarn)
	log.SetFormat("2006-01-02 15:04:05.999 info examples/main.go:88 message")

	dbVar := flag.String("dbvar", "db", "variable of db to open transaction and execute SQL statements")
	dbPath := flag.String("dbpath", "", "path of db to open transaction and execute SQL statements")
	force := flag.Bool("force", false, "force to regenerate, even sourceimpl.go file newer than source.go file")
	version := flag.Bool("v", false, "variable of db to open transaction and execute SQL statements")
	flag.Usage = usage

	flag.Parse()
	if *version {
		fmt.Println("0.5.0")
		return
	}

	goFile := os.Getenv("GOFILE")
	if goFile == "" {
		if flag.NArg() > 1 {
			goFile = flag.Arg(0)
			if !strings.HasSuffix(goFile, ".go") {
				fmt.Println("file suffix must match *.go")
				return
			}
		} else {
			flag.Usage()
		}
	}
	fmt.Printf("Found  go file: %s\n", goFile)

	outFile := goFile[:len(goFile)-3] + "impl.go"
	if !*force {
		// TODO
		// if sourceimpl.go newer than source.go
		// do nothing, skip, return
		// return
	}

	pkg := &domain.Package{
		Source:  goFile,
		DBVar:   *dbVar,
		Imports: map[string]string{},
	}
	if *dbPath != "" {
		ss := strings.Split(*dbVar, ".")
		if len(ss) != 2 {
			fmt.Println("arg 'dbvar' must be <package-name>:<variable-name")
			flag.Usage()
			return
		}
		pkg.Imports[ss[0]] = strings.Trim(*dbPath, `'"`)
	}

	parse.ParseGoFile(pkg)

	prepare.Prepare(pkg)
	// log.JSONIndent(pkg)

	paths := strings.Split(os.Getenv("GOPATH"), string(filepath.ListSeparator))
	tmplFile := filepath.Join(paths[0], "src", "github.com/arstd/light", "templates/pq.gotemplate")

	funcMap := template.FuncMap{
		"timestamp": func() string { return time.Now().Format("2006-01-02 15:04:05") },
	}

	tmpl, err := template.New("pq.gotemplate").Funcs(funcMap).ParseFiles(tmplFile)
	util.CheckError(err)

	// out, err := os.OpenFile(goFile[:len(goFile)-3]+"impl.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	// util.CheckError(err)
	// err = tmpl.Execute(out, pkg)
	// util.CheckError(err)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, pkg)
	util.CheckError(err)

	pretty, err := imports.Process(outFile, buf.Bytes(), nil)
	util.CheckError(err)
	err = ioutil.WriteFile(outFile, pretty, 0644)
	util.CheckError(err)

	fmt.Printf("Generated file: %s\n", outFile)
}
