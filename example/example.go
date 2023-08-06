package main

import (
	"html/template"
	"os"

	"github.com/Ol1BoT/go-heroicons"
)

type Navigation struct {
	Icon template.HTML
	Name string
	Link string
}

var (
	Nav = []Navigation{{
		Icon: heroicons.HomeSolid,
		Name: "Home",
		Link: "/",
	}, {
		Icon: heroicons.UserSolid,
		Name: "Profile",
		Link: "/profile",
	}}
)

func main() {

	tmpl, err := template.ParseFiles("example.gohtml")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(os.Stdout, Nav)

}
