package dotsql_test

import (
	"dotsql"
	"strings"
	"testing"
)

func failIfError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("err == nil, got '%s'", err)
	}
}

func TestLoad(t *testing.T) {
	_, err := dotsql.Load(strings.NewReader(""))
	failIfError(t, err)
}

func TestLoadFromString(t *testing.T) {
	_, err := dotsql.LoadFromString("")
	failIfError(t, err)
}

func TestRaw(t *testing.T) {
	expectedQuery := "SELECT 1+1"

	dot, err := dotsql.LoadFromString("--name: my-query\n" + expectedQuery)
	failIfError(t, err)

	got, err := dot.Raw("my-query")
	failIfError(t, err)

	got = strings.TrimSpace(got)
	if got != expectedQuery {
		t.Errorf("Raw() == '%s', expected '%s'", got, expectedQuery)
	}
}

func TestQueries(t *testing.T) {
	expectedQueryMap := map[string]string{
		"select": "SELECT * from users",
		"insert": "INSERT INTO users (?, ?, ?)",
	}

	dot, err := dotsql.LoadFromString(`
	-- name: select
	SELECT * from users

	-- name: insert
	INSERT INTO users (?, ?, ?)
	`)
	failIfError(t, err)

	got := dot.QueryMap()

	if len(got) != len(expectedQueryMap) {
		t.Errorf("QueryMap() len (%d) differ from expected (%d)", len(got), len(expectedQueryMap))
	}

	for name, query := range got {
		if query != expectedQueryMap[name] {
			t.Errorf("QueryMap()[%s] == '%s', expected '%s'", name, query, expectedQueryMap[name])
		}
	}
}

func TestMergeHaveBothQueries(t *testing.T) {
	expectedQueryMap := map[string]string{
		"query-a": "SELECT * FROM a",
		"query-b": "SELECT * FROM b",
	}

	a, err := dotsql.LoadFromString("--name: query-a\nSELECT * FROM a")
	failIfError(t, err)

	b, err := dotsql.LoadFromString("--name: query-b\nSELECT * FROM b")
	failIfError(t, err)

	c := dotsql.Merge(a, b)

	got := c.QueryMap()
	if len(got) != len(expectedQueryMap) {
		t.Errorf("QueryMap() len (%d) differ from expected (%d)", len(got), len(expectedQueryMap))
	}
}

func TestMergeTakesPresecendeFromLastArgument(t *testing.T) {
	expectedQuery := "SELECT * FROM c"

	a, err := dotsql.LoadFromString("--name: query\nSELECT * FROM a")
	failIfError(t, err)

	b, err := dotsql.LoadFromString("--name: query\nSELECT * FROM b")
	failIfError(t, err)

	c, err := dotsql.LoadFromString("--name: query\nSELECT * FROM c")
	failIfError(t, err)

	x := dotsql.Merge(a, b, c)

	got := x.QueryMap()["query"]
	if expectedQuery != got {
		t.Errorf("Expected query: '%s', got: '%s'", expectedQuery, got)
	}
}
