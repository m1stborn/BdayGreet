package handler

import (
	"fmt"
	"net/http"

	"github.com/m1stborn/BdayGreet/internal/pkg/model"

	"github.com/julienschmidt/httprouter"
)

var GetUSerByData = model.DB.GetUserByDate

func HandleBdayGreet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	maleMsg := "We offer special discount 20%% for the following items:\nWhite Wine, iPhone X\n"
	femaleMsg := "We offer special discount 50%% for the following items:\nCosmetic, LV Handbags\n"

	//users, _ := model.DB.GetUserByDate(8, 8)
	users, _ := GetUSerByData(8, 8)

	for _, user := range users {
		replyMsg := fmt.Sprintf("Subject: Happy birthday!\nHappy birthday, dear %s!\n", user.FirstName)
		if user.Gender == "Male" {
			fmt.Fprintf(w, replyMsg+maleMsg)
		} else {
			fmt.Fprintf(w, replyMsg+femaleMsg)
		}
	}
}
