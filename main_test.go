package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type testState struct {
	originalWD string
	tempdir    string
}

func (state *testState) teardown() {
	err := os.Chdir(state.originalWD)
	if err != nil {
		panic(err.Error())
	}

	err = os.RemoveAll(state.tempdir)
	if err != nil {
		panic(err.Error())
	}
}

func setup() *testState {
	tempdir, err := ioutil.TempDir("", "tmp-gogitit-")
	if err != nil {
		panic(err.Error())
	}

	_, err = gitInvoke("clone", "wikidata", tempdir+"/wikidata")
	if err != nil {
		panic(err.Error())
	}

	originalWD, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	err = os.Chdir(tempdir)
	if err != nil {
		panic(err.Error())
	}

	return &testState{
		originalWD: originalWD,
		tempdir:    tempdir,
	}
}

func TestRoot200(t *testing.T) {
	state := setup()
	defer state.teardown()

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(rr, req)

	res := rr.Result()

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)
		t.Errorf("no dice")
	}
}

func TestMain(m *testing.M) {
	registerHandlers()

	os.Exit(m.Run())
}
