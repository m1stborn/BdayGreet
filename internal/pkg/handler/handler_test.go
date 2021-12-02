package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/m1stborn/BdayGreet/internal/pkg/model"
)

func TestHandleBdayGreet(t *testing.T) {
	// integration test on http requests to HandleBdayGreet
	old := GetUserByDate
	defer func() { GetUserByDate = old }()

	GetUserByDate = func(month int, day int) ([]model.User, error) {
		return []model.User{
			{
				FirstName: "Robert",
				LastName:  "Yen",
				Gender:    "Male",
				Birth:     time.Date(1975, 8, 8, 0, 0, 0, 0, time.UTC),
				Email:     "robert.yen@linecorp.com",
			},
			{
				FirstName: "Sherry",
				LastName:  "Chen",
				Gender:    "Female",
				Birth:     time.Date(1993, 8, 8, 0, 0, 0, 0, time.UTC),
				Email:     "sherry.chen@linecorp.com",
			},
		}, nil
	}

	router := httprouter.New()
	router.GET("/api/bdaygreet/", HandleBdayGreet)

	req, _ := http.NewRequest("GET", "/api/bdaygreet/", nil)
	rr := httptest.NewRecorder()

	HandleBdayGreet(rr, req, httprouter.Params{})

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v", status)
	}
	fmt.Println(rr.Body.String())
	expected := `Subject: Happy birthday!
Happy birthday, dear Yen, Robert!
Subject: Happy birthday!
Happy birthday, dear Chen, Sherry!
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}
