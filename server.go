/*package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

//var tmpl = template.Must(template.ParseFiles("index.html", "assets/upcomingMovie.html", "header.html", "footer.html"))

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "secretory.html", nil)
}
func up(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "admin_registor.html", nil)
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/admin", up)

	http.ListenAndServe(":8080", nil)
}*/
package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

//var tmpl = template.Must(template.ParseFiles("index.html", "assets/upcomingMovie.html", "header.html", "footer.html"))

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "fieldman.html", nil)
}
func up(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "resourse.html", nil)
}
func sec(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "second.html", nil)
}
func sear(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "searchadmin.html", nil)
}
func secrotor(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "secretory.html", nil)
}
func admreg(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "admin_register.html", nil)
}
func searchal(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "search.html", nil)
}
func pay(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "payment.html", nil)
}
func teach(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "teacher.html", nil)
}
func secrpay(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "secretarypayment.html", nil)
}
func studentroun(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "roundstudent.html", nil)
}
func roundpay(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "roundpayment.html", nil)
}
func eachroun(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "eachround.html", nil)
}
func editroun(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "editround.html", nil)
}
func generate(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "generatedata.html", nil)
}
func courseconf(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "courseconfig.html", nil)
}
func populate(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "populatexlsx.html", nil)
}
func categoryconf(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "categoryconfig.html", nil)
}
func main() {
	fs := http.FileServer(http.Dir("css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fsj := http.FileServer(http.Dir("js/"))
	http.Handle("/js/", http.StripPrefix("/js/", fsj))
	fsk := http.FileServer(http.Dir("img/"))
	http.Handle("/img/", http.StripPrefix("/img/", fsk))
	http.HandleFunc("/", index)
	http.HandleFunc("/resourse", up)
	http.HandleFunc("/second", sec)
	http.HandleFunc("/search", sear)
	http.HandleFunc("/secrotory", secrotor)
	http.HandleFunc("/adminregistor", admreg)
	http.HandleFunc("/searchall", searchal)
	http.HandleFunc("/paymentt", pay)
	http.HandleFunc("/teacher", teach)
	http.HandleFunc("/secrotorypay", secrpay)
	http.HandleFunc("/studentround", studentroun)
	http.HandleFunc("/roundpaymen", roundpay)
	http.HandleFunc("/eachround", eachroun)
	http.HandleFunc("/editround", editroun)
	http.HandleFunc("/generatedata", generate)
	http.HandleFunc("/courseconfig", courseconf)
	http.HandleFunc("/populatexlsx", populate)
	http.HandleFunc("/categoryconfig", categoryconf)

	http.ListenAndServe(":8989", nil)
}
