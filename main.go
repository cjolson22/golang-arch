package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/", bar)
	http.ListenAndServe(":8080", nil)
}

func getCode(msg string) string{
	
}
func bar (w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Redirect(w,r, "/", http.StatusSeeOther)
	}

	email := r.FormValue("email")
	if email == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}


	c := http.Cookie{
		Name: "session",
		Value: email
	}
}
func foo(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>HMAC Example</title>
	</head>
	<body>
		<form action="/submit" method="post">
			<input type="email" name="email"/>
			<input type="submit" />

		</form>
	</body>
	</html>`

	io.WriteString(w, html)
}
