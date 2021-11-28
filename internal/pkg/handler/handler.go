package handler

import (
	"fmt"
	"net/http"

	"github.com/m1stborn/BdayGreet/internal/pkg/model"

	"github.com/julienschmidt/httprouter"
)

func HandleBdayGreet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, _ := model.MongoDB.FindByBday(8, 8)
	for _, user := range users {
		replyMsg := fmt.Sprintf("Subject: Happy birthday!\nHappy birthday, dear %s!\n", user["firstname"])
		fmt.Fprintf(w, replyMsg)
	}
}
