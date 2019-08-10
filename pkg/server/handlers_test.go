package server_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	. "github.com/itmecho/frontdesk/pkg/server"
	"github.com/itmecho/frontdesk/pkg/types"
	"github.com/sirupsen/logrus"
)

type mockStore struct {
	Users []types.User
}

func (s mockStore) Migrate() error { return nil }

func (s mockStore) GetAll() ([]types.User, error) {
	return s.Users, nil
}
func (s mockStore) GetByID(id string) (*types.User, error) {
	return &s.Users[0], nil
}

func (s mockStore) GetByEmail(email string) (*types.User, error) {
	return &s.Users[0], nil
}

func (s mockStore) Create(u *types.User) error {
	return nil
}

func (s mockStore) Update(u *types.User) error {
	return nil
}

func (s mockStore) Delete(u *types.User) error {
	return nil
}

type mockAuthenticator struct{}

func (a mockAuthenticator) GenerateToken(u *types.User) (string, error) {
	return "secret-token", nil
}

func (a mockAuthenticator) CheckToken(t string) error {
	return nil
}

var testUsers = []types.User{
	{
		ID:       "88def0a5-dff8-4d4c-a44d-12131f461020",
		Name:     "John Wick",
		Email:    "john.wick@mailinator.com",
		Password: []byte("$2a$10$VItjGiFXy/bhDEREYu5NX.u18eWF2OzFqpo1G5UDCwkzwvH8rYtyq"),
	},
	{
		ID:       "8973b6c6-c99e-415b-8b0d-57a2aa450c66",
		Name:     "Jason Bourne",
		Email:    "jason.bourne@mailinator.com",
		Password: []byte("$2a$10$VItjGiFXy/bhDEREYu5NX.u18eWF2OzFqpo1G5UDCwkzwvH8rYtyq"),
	},
	{
		ID:       "a9045c2d-b787-463d-a99a-4e5a2d5c775b",
		Name:     "Chuck Bartowski",
		Email:    "chuck.bartowski@mailinator.com",
		Password: []byte("$2a$10$VItjGiFXy/bhDEREYu5NX.u18eWF2OzFqpo1G5UDCwkzwvH8rYtyq"),
	},
}

func getTestServer(users []types.User) Server {
	nullLogger := logrus.New()
	nullLogger.Out = ioutil.Discard
	srv := New(8080, nullLogger, mockStore{users}, mockAuthenticator{})
	return srv
}

func TestHandleGetUsers(t *testing.T) {
	srv := getTestServer(testUsers)
	req := httptest.NewRequest("GET", "/api/users", nil)
	rec := httptest.NewRecorder()
	srv.Handler.ServeHTTP(rec, req)
	resp := rec.Result()
	if resp.StatusCode != 200 {
		t.Error("expected 200 response code, got ", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	var actual []types.User
	decoder.Decode(&actual)

	if len(testUsers) != len(actual) {
		t.Errorf("expected %d users, got %d", len(testUsers), len(actual))
	}

	for i, u := range actual {
		if testUsers[i].ID != u.ID {
			t.Errorf("expected user ID %s, got %s", testUsers[i].ID, u.ID)
		}
	}
}
