# Go-Heroicons

a wrapper around heroicons that are of the template.HTML type, so they can be used in the go templating language.

this should be used in conjunction with tailwindcss.

```go
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
		Icon: heroicons.Home,
		Name: "Home",
		Link: "/",
	}, {
		Icon: heroicons.User,
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

```

```html
<ul>
  {{ range $i, $v := .}}
  <li>{{$v.Icon}} - {{$v.Name}}</li>
  {{end }}
</ul>
```
