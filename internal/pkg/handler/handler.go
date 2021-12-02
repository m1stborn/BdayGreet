package handler

import (
	"fmt"
	ht "html/template"
	"net/http"
	"time"

	"github.com/m1stborn/BdayGreet/internal/pkg/model"

	"github.com/julienschmidt/httprouter"
)

var ImageTemplate string = `<!DOCTYPE html>
<html>
<head>
    <title>Birthday Greeting</title>
</head>
<body>
	{{ range .Users }}
		<div>
		Subject: Happy birthday!
		</div>
		<div>
		Happy birthday, dear {{.Name}}
		</div>
		{{ if .AgeGt49 }}
			<img src="https://imgur.com/NgkfJJv.jpg" width="20%" height="20%">
		{{ end }}
	{{ end }}
</body>
</html>`

type HtmlUser struct {
	Name    string
	AgeGt49 bool
}

var GetUserByDate = model.DB.GetUserByDate

func HandleBdayGreet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//users, _ := model.DB.GetUserByDate(12, 22)
	users, _ := GetUserByDate(12, 22)

	var names []HtmlUser
	var gt49 bool

	for _, user := range users {
		now := time.Now()
		dur := now.Sub(user.Birth)
		age := dur.Seconds() / 31207680
		if age > 49 {
			gt49 = true
		} else {
			gt49 = false
		}
		names = append(names, HtmlUser{
			Name:    user.FirstName,
			AgeGt49: gt49,
		})
	}
	t, tErr := ht.New("webpage").Parse(ImageTemplate)
	if tErr != nil {
		fmt.Println(tErr)
	}
	tErr = t.Execute(w, struct {
		Users []HtmlUser
	}{
		names,
	})

}
