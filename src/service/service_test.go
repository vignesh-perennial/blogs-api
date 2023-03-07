package service_test

import (
	"blogs_api/src/controller"
	"net/http"
	"net/http/httptest"
	"testing"
)

var blogcontroller = new(controller.BlogController)

func TestGetBlogByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/articals/:id", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(blogcontroller.GetBlogByID)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":"6405af0349229973f3c53e50","title":"Book 1","author":"Author 1","content":"content 1"},{"id":"64061d7c057174edb680c5fb","title":"Book 2","author":"Author 2", ,"content":"content 1"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
