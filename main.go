package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Info struct {
	Action     string
	Method     string
	Url        string
	Params     string
	RetSuccess string
	RetError   string
}

func main() {

	file := flag.String("file", "", "The text file to parse")
	flag.Parse()

	r, _ := regexp.Compile("(?s)\\/\\*\\nHTTP API.*?\\*\\/")
	content, _ := regexp.Compile(": .*\n")

	ac, _ := regexp.Compile(".*Action : .*\n")
	method, _ := regexp.Compile(".*Method : .*\n")
	url, _ := regexp.Compile(".*Url : .*\n")
	params, _ := regexp.Compile(".*Params : .*\n")
	suc, _ := regexp.Compile(".*Return success : .*\n")
	er, _ := regexp.Compile(".*Return error : .*\n")

	b, err := ioutil.ReadFile(*file) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	allstr := r.FindAllString(str, -1)

	var infos []*Info

	for _, v := range allstr {

		i := &Info{}
		i.Action = strings.TrimSpace(strings.TrimPrefix(content.FindString(ac.FindString(v)), ": "))
		i.Method = strings.TrimSpace(strings.TrimPrefix(content.FindString(method.FindString(v)), ": "))
		i.Url = strings.TrimSpace(strings.TrimPrefix(content.FindString(url.FindString(v)), ": "))
		i.Params = strings.TrimSpace(strings.TrimPrefix(content.FindString(params.FindString(v)), ": "))
		i.RetSuccess = strings.TrimSpace(strings.TrimPrefix(content.FindString(suc.FindString(v)), ": "))
		i.RetError = strings.TrimSpace(strings.TrimPrefix(content.FindString(er.FindString(v)), ": "))

		infos = append(infos, i)
	}

	wtf, err := os.Create(*file + ".md")
	if err != nil {
		log.Fatal(err)
	}
	defer wtf.Close()

	w := bufio.NewWriter(wtf)

	fmt.Fprintln(w, *file)
	fmt.Fprintln(w, "==========================================================")
	fmt.Fprintln(w, "")

	for _, v := range infos {

		fmt.Fprintln(w, "**"+v.Url+"** ["+v.Method+"]")
		fmt.Fprintln(w, "----------------------------")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "	"+v.Action)
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "	- Params : "+v.Params)
		fmt.Fprintln(w, "	- Success : "+v.RetSuccess)
		fmt.Fprintln(w, "	- Error : "+v.RetError)
		fmt.Fprintln(w, "")
	}

	w.Flush()
}
