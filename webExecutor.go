package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"html/template"
)
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func execute(command string) string{
	var c1 string
	var c2 string
	is:=false
	for i:=0;i<len(command);i+=1{
		if command[i] == ' '{
			c1 = command[0:i]
			c2 = command[i+1:len(command)]
			is = true
			break

		}
	}
	fmt.Println(c1)
	fmt.Println(c2)
	if is {
		out, err := exec.Command(c1, c2).Output()
		if err != nil {
			log.Fatal(err)
		}
		return string(out)
	} else{
		out, err := exec.Command(command).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)}

}
func execHandler(w http.ResponseWriter, r *http.Request){
	body := r.FormValue("body")
	fmt.Fprintf(w,"%s",execute(body))

	//http.Redirect(w, r, "/comm/", http.StatusFound)
}
func commandHandler(w http.ResponseWriter, r *http.Request){
	title := "Execute"
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "comm", p)
}

func main() {
	http.HandleFunc("/comm/", commandHandler)
	http.HandleFunc("/exec/", execHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}