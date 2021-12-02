package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/m1stborn/BdayGreet/internal/pkg/model"
)

func TestHandleBdayGreet(t *testing.T) {
	// integration test on http requests to HandleBdayGreet
	old := GetUserByDate
	defer func() { GetUserByDate = old }()

	GetUserByDate = func(month int, day int) ([]model.User, error) {
		return []model.User{
			{
				FirstName: "Peter",
				LastName:  "Wang",
				Gender:    "Male",
				Birth:     time.Date(1950, 12, 22, 0, 0, 0, 0, time.UTC),
				Email:     "peter.wang@linecorp.com",
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

	expected := `<!DOCTYPE html>
<html>
<head>
    <title>Birthday Greeting</title>
</head>
<body>
	
		<div>
		Subject: Happy birthday!
		</div>
		<div>
		Happy birthday, dear Peter
		</div>
		
			<img src="https://imgur.com/NgkfJJv.jpg" width="20%" height="20%">
		
	
</body>
</html>`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}
