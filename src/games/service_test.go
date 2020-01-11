package games

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KayacChang/API_Server/pg"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func send(mathod string, url string, body io.Reader) *http.Response {
	r := httprouter.New()

	Serve(r, pg.New("test"))

	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(mathod, url, body)

	r.ServeHTTP(rec, req)

	return rec.Result()
}

func TestGetGames(t *testing.T) {

	res := send("GET", "/games", nil)

	// === Test ===
	assert.Equal(t, http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("%s\n", string(data))
}

func TestGetGameByID(t *testing.T) {

	res := send("GET", "/games/1", nil)

	// === Test ===
	assert.Equal(t, http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("%s\n", string(data))
}

func TestPostGame(t *testing.T) {

	input := `
	{
		"name": "test",
		"href": "http://test.com"
	}
	`

	res := send("POST", "/games", bytes.NewBuffer([]byte(input)))

	// === Test ===
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("%s\n", string(data))
}
