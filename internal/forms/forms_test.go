package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields mission")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it doees")
	}

}

func TestHas(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("no-has", "")
	postedData.Add("has", "something")
	r.PostForm = postedData
	form := New(r.PostForm)

	if form.Has("no-has") {
		t.Error("Got has when entry empty string")
	}
	if !form.Has("has") {
		t.Error("Got no has when entry string")
	}
}

func TestMin_Length(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("no_minlenght", "aa")
	postedData.Add("yes_minlengh", "aaa")

	r.PostForm = postedData
	form := New(r.PostForm)
	if form.MinLenght("no_minlenght", 3) {
		t.Error("Got minleght OK  when does not")
	}
	if !form.MinLenght("yes_minlengh", 3) {
		t.Error("No got minlenght OK when shoud does")
	}

}

func Test_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("invalid_email", "asdas.as")
	postedData.Add("valid_email", "email@domain.com")

	r.PostForm = postedData
	form := New(r.PostForm)

	form.IsEmail("invalid_email")
	if form.Errors.Get("invalid_email") != "Invalid Email Address" {
		t.Error("No got error when passing invalid email address")

	}
	form.IsEmail("valid_email")
	if form.Errors.Get("valid_email") == "Invalid Email Address" {
		t.Error("got error when passing email address")

	}

}
