package mock_repo

import (
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"go-mock/models"
	"go-mock/router"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetById(t *testing.T) {
	testCases := []struct {
		State    string
		user     models.Person
		errExist bool
		action   func(mock *MockRepo, person models.Person, id int)
	}{

		{
			State: "OK",
			user: models.Person{
				Id:    1,
				Name:  "Otabek",
				Email: "otabek94_30@mail.ru",
			},
			errExist: false,
			action: func(mock *MockRepo, person models.Person, id int) {
				mock.EXPECT().FindById(id).Times(0).Return(person, nil)
			},
		},
		{
			State: "Not Found",
			user: models.Person{
				Id: -1,
			},
			errExist: true,
			action: func(mock *MockRepo, person models.Person, id int) {
				mock.EXPECT().FindById(id).Times(0).Return(person, errors.New(sql.ErrNoRows.Error()))
			},
		},
	}
	handler := router.Router()
	for _, c := range testCases {
		t.Run(c.State, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer func() {
				ctrl.Finish()
			}()
			mock := NewMockRepo(ctrl)
			req := httptest.NewRequest(http.MethodGet, "http://localhost:8081/get-account/?id="+strconv.Itoa(c.user.Id), nil)
			resp := httptest.NewRecorder()
			c.action(mock, c.user, c.user.Id)
			handler.ServeHTTP(resp, req)
			if resp.Code != http.StatusOK && !c.errExist {
				t.Errorf("got %d status, but wanted %d", resp.Code, http.StatusOK)
			}
		})
	}
	//assert.NoError(t)
}
