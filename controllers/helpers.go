package controllers

import (
	"github.com/gorilla/schema"
	"net/http"
)

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
