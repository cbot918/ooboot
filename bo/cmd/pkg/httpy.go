package pkg

import "strings"

func GetHttpy() string {

	content := `
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cbot918/liby/httpy"
)

var lg = fmt.Println
var lf = fmt.Printf

const port = ":8080"

type UserRequest struct {
	Password string 'json:"password"'
	Email    string 'json:"email"'
}

func main() {

	r := httpy.New()

	r = SetRouter(r)

	fmt.Println("listening ", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

func SetRouter(r *httpy.Engine) *httpy.Engine {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) error {
		fmt.Println("request get in")

		return nil
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) error {

		user := &UserRequest{}

		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			return err
		}
		lf("%#+v", user)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte('{"hello":"world"}'))

		return nil
	})

	return r
}`

	return strings.ReplaceAll(content, "'", "`")
}
