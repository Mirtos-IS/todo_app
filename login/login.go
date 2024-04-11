package login

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))

func GetUserId(request *http.Request) (userId int) {
    if cookie, err := request.Cookie("session"); err == nil {
        cookieValue := make(map[string]int)
        if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
            userId = cookieValue["id"]
        }
    }
    return userId
}

func SetSession(id int, response http.ResponseWriter) {
    value := map[string]int{
        "id": id,
    }
    if encoded, err := cookieHandler.Encode("session", value); err == nil {
        cookie := &http.Cookie{
            Name: "session",
            Value: encoded,
            Path: "/",
        }
        http.SetCookie(response, cookie)
    }
}

func ClearSession(response http.ResponseWriter) {
    cookie := &http.Cookie{
        Name: "session",
        Value: "",
        Path: "/",
        MaxAge: -1,
    }
    http.SetCookie(response, cookie)
}

