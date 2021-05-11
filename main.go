package main

import (
	model "Groupie-Tracker/src/models"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", hom)
	http.HandleFunc("/test", test)                                                             // the func redirect to the hom function when we go to localhost:6082/                                                       // the func redirect to the read function when	we go to localhost:6082/read
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) // we handle /assets/ in our localhost
	fmt.Printf("Starting server at port 6085\n")
	if err := http.ListenAndServe(":6085", nil); err != nil { // we open the serve and we have and we have an error handling
		log.Fatal(err)
	}
}

type PageData struct {
	Image   []string
	Name    []string
	Members []string
}

func hom(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HOM")
	Artists := model.LoadData()
	TooPrint := make([]string, 0)
	TooPrintName := make([]string, 0)
	TooPrintMembers := make([]string, 0)
	tpl := template.Must(template.ParseFiles("mygptrack/Select.html"))
	TooPrintName = append(TooPrintName, " a")
	for i := range Artists {
		//Image
		TooPrint = append(TooPrint, Artists[i].Image)
		//Name
		TooPrintName = append(TooPrintName, Artists[i].Name)
		TooPrintName = append(TooPrintName, "a")
		//Members
		// TooPrintMembers = append(TooPrintMembers, Artists[i].Members)
	}
	TooPrintName = append(TooPrintName, " ")

	data := PageData{
		Image:   TooPrint,
		Name:    TooPrintName,
		Members: TooPrintMembers,
	}
	err := tpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}

}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HOM")
	Artists := model.LoadData()
	TooPrint := make([]string, 0)
	tpl := template.Must(template.ParseFiles("mygptrack/test.html"))
	for i := range Artists {
		TooPrint = append(TooPrint, Artists[i].Image)
	}
	data := PageData{
		Image: TooPrint,
	}
	err := tpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}