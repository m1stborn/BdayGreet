package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/m1stborn/BdayGreet/internal/pkg/model"

	"github.com/julienschmidt/httprouter"
)

type greetMsg struct {
	Title   string
	Content string
}

var GetUserByDate = model.DB.GetUserByDate

func HandleBdayGreetJson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var msgs []greetMsg

	//users, _ := model.DB.GetUserByDate(8, 8)
	users, _ := GetUserByDate(8, 8)

	for _, user := range users {
		title := fmt.Sprintf("Subject: Happy birthday!")
		content := fmt.Sprintf("Happy birthday, dear %s!", user.FirstName)
		msgs = append(msgs, greetMsg{Title: title, Content: content})
	}
	if len(msgs) == 0 {
		msgs = make([]greetMsg, 0)
	}

	js, err := json.Marshal(msgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func HandleBdayGreetXml(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var msgs []greetMsg

	//users, _ := model.DB.GetUserByDate(8, 8)
	users, _ := GetUserByDate(8, 8)

	for _, user := range users {
		title := fmt.Sprintf("Subject: Happy birthday!")
		content := fmt.Sprintf("Happy birthday, dear %s!", user.FirstName)
		msgs = append(msgs, greetMsg{Title: title, Content: content})
	}
	if len(msgs) == 0 {
		msgs = make([]greetMsg, 0)
	}

	x, err := xml.MarshalIndent(msgs, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
