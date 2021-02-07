package main

import (
	"fmt"
	"net/http"
	"html/template"
	)

type User struct {
	Name string // imya~
	Age uint16 // uint- ne mozhet bit otritsatelnim
	Money int16 // balans
	Avg_grades, Happiness float64 // 1) sredmiy ball ego otsenok
	Hobbies []string
} // type

func (u User) getAllInfo() string {
	return fmt.Sprintf("Username is : %s. He is %d and he has money" +
		" equal :%d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string){
	u.Name = newName
}


func home_page(page http.ResponseWriter,r *http.Request)  { // 2+ - parametr. 1) obrashaemsya stranicke chto libo pokazivat polzivatelyu.2)Vsegda Peredaetsya. Otlidit podklyuchenie k stranitsel;
	// var bob User = ....
	bob:= User{"Boboy", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}} // type
//	bob.setNewName("Hakim")
//	fmt.Fprintf(page, `<h1>Main</h1>
//		<b>main</b>`)
	tmpl, _ := template.ParseFiles("html/home_page.html")
	tmpl.Execute(page, bob)

}

func contacts_page(page http.ResponseWriter,r *http.Request)  { // 2+ - parametr. 1) obrashaemsya stranicke chto libo pokazivat polzivatelyu.2)Vsegda Peredaetsya. Otlidit podklyuchenie k stranitsel;

	fmt.Fprintf(page, "Contacts page!") // fprint- formatirovanaya stroka. 1) kuda budem pisat -> page;
}

func handlerRequest()  {
	http.HandleFunc("/", home_page) // otsledit perekhod URL adress. /- Glavnaya stranitsa. 2-parametr;
	http.HandleFunc("/contacts/", contacts_page) // contact_page
	http.ListenAndServe(":8800", nil) // Local server. lichnie nastroyki dlya samogo servera.nill ->pustota;
}

func main()  {

handlerRequest() // zapusk servera

}
