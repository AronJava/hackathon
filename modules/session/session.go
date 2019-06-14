package session

import(
	"github.com/gorilla/sessions"
	"os"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
