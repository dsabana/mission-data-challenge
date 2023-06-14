package internal_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"mission-data-challenge/internal"
	internal_test "mission-data-challenge/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func beforeTests(t *testing.T) (*mux.Router, *internal_test.MockRepository) {
	t.Helper()

	mCtrl := gomock.NewController(t)
	mr := internal_test.NewMockRepository(mCtrl)

	s := internal.NewService(mr)

	r := internal.SetupRouter(s)

	return r, mr
}

func TestCompaniesHandler_AddCompany(t *testing.T) {
	handler, mr := beforeTests(t)
	mockId := "someID"
	newJournalMock := internal.Journal{
		Name: "test journal",
		Id:   &mockId,
	}

	mockErr := struct {
		Message string
	}{
		Message: "error saving journal",
	}

	t.Run("successfully creates journal", func(t *testing.T) {
		// given
		var mockReq = internal.Journal{
			Name: "test journal",
		}

		requestBody := new(bytes.Buffer)
		err := json.NewEncoder(requestBody).Encode(mockReq)
		assert.NoError(t, err)

		w, r := SetupRequest(t, http.MethodPost, "http://localhost:8000/journals", requestBody.Bytes())

		// when
		mr.EXPECT().SaveJournal(gomock.Any(), mockReq).Times(1).Return(&newJournalMock, nil)

		handler.ServeHTTP(w, r)

		// then
		var response internal.Journal
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, newJournalMock, response)
	})

	t.Run("returns 500 while creates journal", func(t *testing.T) {
		// given
		var mockReq = internal.Journal{
			Name: "test journal",
		}

		requestBody := new(bytes.Buffer)
		err := json.NewEncoder(requestBody).Encode(mockReq)
		assert.NoError(t, err)

		w, r := SetupRequest(t, http.MethodPost, "http://localhost:8000/journals", requestBody.Bytes())

		// when
		mr.EXPECT().SaveJournal(gomock.Any(), mockReq).Times(1).Return(nil, fmt.Errorf("some error"))

		handler.ServeHTTP(w, r)

		// then
		var response struct {
			Message string
		}
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, mockErr, response)
	})
}

func SetupRequest(t *testing.T, httpMethod string, target string, data []byte) (*httptest.ResponseRecorder, *http.Request) {
	t.Helper()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(
		httpMethod,
		target,
		bytes.NewReader(data),
	)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "some-great-token")

	return w, r
}
