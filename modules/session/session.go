package session

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
