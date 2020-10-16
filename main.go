package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/submit", bar)
	http.ListenAndServe(":8080", nil)
}

func getJWT(msg string) (string, error) {
	myKey := "I love thursdays when it rains"

	type myClaims struct {
		jwt.StandardClaims
		Email string
	}

	claims := myClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	ss, err := token.SignedString([]byte(myKey))
	if err != nil {
		return "", fmt.Errorf("couldn't SignedString in NewWithClaims %w", err)
	}
	return ss, nil
	// 	h.Write([]byte(msg))
	// 	return fmt.Sprintf("%x", h.Sum(nil))
}
func bar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	email := r.FormValue("emailThing")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	ss, err := getJWT(email)
	if err != nil {
		http.Error(w, "Couldn get JWT", http.StatusInternalServerError)
	}
	c := http.Cookie{
		Name:  "session",
		Value: ss + "|" + email,
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{}
	}

	//isEqual := true

	// message := "Not logged in"
	// if {
	// 	message = "Logged in"
	// }

	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>HMAC Example</title>
	</head>
	<body>
		<p>Cookie value: ` + c.Value + `</p>
		<p>` + "message" + `</p>
		<form action="/submit" method="post">
			<input type="emailThing" name="email"/>
			<input type="submit" />

		</form>
	</body>
	</html>`

	io.WriteString(w, html)
}
