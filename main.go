package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Tasks struct {
	Name string
	Task string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	data := map[string][]Tasks{
		"Tasks": {
			{Name: "Ibrahim", Task: "Clean the dishes"},
			{Name: "Mohamed", Task: "Eat Dinner"},
		}}

	tmpl.Execute(w, data)

}

func AddTask(w http.ResponseWriter, r *http.Request) {

	time.Sleep(1 * time.Second)

	name := r.PostFormValue("name")
	task := r.PostFormValue("task")
	// log.Print("HTMX request received")
	// log.Print(r.Header.Get("HX-Request"))

	// htmlStr:=fmt.Sprintf("<div id='tasks'class='bg-light m-1 p-1 border rounded space-between'><p class='display-6'>Task: %s</p><p class='text-muted'><i>created by %s</i> </p></div>",task,name)
	// tmpl,_:=template.New("t").Parse(htmlStr)
	// tmpl.Execute(w, nil)
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	tmpl.ExecuteTemplate(w, "task-element", Tasks{Name: name, Task: task})

}

func main() {

	//handle static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//routes
	http.HandleFunc("/", Home)
	http.HandleFunc("/add-task", AddTask)

	fmt.Println("Starting server...")
	fmt.Println("View app at http://localhost:8000")

	//handle server
	log.Fatal(http.ListenAndServe(":8000", nil))

}
