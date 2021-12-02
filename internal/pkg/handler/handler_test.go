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

var mockUser = []model.User{
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
}

func TestHandleBdayGreetJson(t *testing.T) {
	// integration test on http requests to HandleBdayGreet
	old := GetUserByDate
	defer func() { GetUserByDate = old }()

	GetUserByDate = func(month int, day int) ([]model.User, error) {
		return mockUser, nil
	}

	router := httprouter.New()
	router.GET("/api/bdaygreet/json", HandleBdayGreetJson)

	req, _ := http.NewRequest("GET", "/api/bdaygreet/", nil)
	rr := httptest.NewRecorder()

	HandleBdayGreetJson(rr, req, httprouter.Params{})

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v", status)
	}
	fmt.Println(rr.Body.String())
	expected := `[{"Title":"Subject: Happy birthday!","Content":"Happy birthday, dear Robert!"},{"Title":"Subject: Happy birthday!","Content":"Happy birthday, dear Sherry!"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}

func TestHandleBdayGreetXml(t *testing.T) {
	// integration test on http requests to HandleBdayGreet
	old := GetUserByDate
	defer func() { GetUserByDate = old }()

	GetUserByDate = func(month int, day int) ([]model.User, error) {
		return mockUser, nil
	}

	router := httprouter.New()
	router.GET("/api/bdaygreet/xml", HandleBdayGreetXml)

	req, _ := http.NewRequest("GET", "/api/bdaygreet/", nil)
	rr := httptest.NewRecorder()

	HandleBdayGreetXml(rr, req, httprouter.Params{})

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v", status)
	}
	fmt.Println(rr.Body.String())
	expected := `<greetMsg>
	<Title>Subject: Happy birthday!</Title>
	<Content>Happy birthday, dear Robert!</Content>
</greetMsg>
<greetMsg>
	<Title>Subject: Happy birthday!</Title>
	<Content>Happy birthday, dear Sherry!</Content>
</greetMsg>`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}
