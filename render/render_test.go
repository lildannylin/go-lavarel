package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "no-file", true, "no error rendering non-existent go template, when one is expected"},
	{"jet_page", "jet", "home", false, "error rendering jet template"},
	{"jet_page_no_template", "jet", "no-file", true, "no error rendering non-existent jet template, when one is expected"},
	{"invalid_render_engine", "foo", "home", true, "no error rendering with non-existent template engine"},
}

func TestRender_Page(t *testing.T) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existent template", err)
	}

	testRenderer.Renderer = "jet"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existent template", err)
	}

	testRenderer.Renderer = ""
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err == nil {
		t.Error("No Error return while rendering with invalid renderer specified", err)
	}
}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}
}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "jet"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}
}
