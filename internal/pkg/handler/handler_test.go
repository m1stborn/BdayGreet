package handler

import (
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleBdayGreet(t *testing.T) {
	// integration test on http requests to HandleBdayGreet
	old := FindByBday
	defer func() { FindByBday = old }()

	FindByBday = func(month int, day int) ([]bson.M, error) {
		//return []model.User{
		//	{
		//		FirstName: "Robert",
		//		LastName:  "Yen",
		//		Gender:    "Male",
		//		Birth:     time.Date(1975, 8, 8, 0, 0, 0, 0, time.UTC),
		//		Email:     "robert.yen@linecorp.com",
		//	},
		//	{
		//		FirstName: "Sherry",
		//		LastName:  "Chen",
		//		Gender:    "Female",
		//		Birth:     time.Date(1993, 8, 8, 0, 0, 0, 0, time.UTC),
		//		Email:     "sherry.chen@linecorp.com",
		//	},
		//}, nil
		return []bson.M{
			map[string]interface{}{
				"firstname": "Robert",
			},
			map[string]interface{}{
				"firstname": "Sherry",
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

	expected := `Subject: Happy birthday!
Happy birthday, dear Robert!
Subject: Happy birthday!
Happy birthday, dear Sherry!
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}
