// Description: sessions
// Author: ZHU HAIHUA
// Since: 2016-04-07 22:20
package sessions

import (
	. "github.com/gin-gonic/contrib/sessions"
	"gopkg.in/mgo.v2"
	"testing"
)

const mongoTestServer = "localhost:6379"

var newMongoStore = func(_ *testing.T) Store {
	dbsess, err := mgo.Dial(mongoTestServer)
	store := NewMongoStore(dbsess, "session", "session", 3600, true, []byte("secret"))
	return store
}

func TestMongo_SessionGetSet(t *testing.T) {
	sessionGetSet(t, newMongoStore)
}

func TestMongo_SessionDeleteKey(t *testing.T) {
	sessionDeleteKey(t, newMongoStore)
}

func TestMongo_SessionFlashes(t *testing.T) {
	sessionFlashes(t, newMongoStore)
}

func TestMongo_SessionClear(t *testing.T) {
	sessionClear(t, newMongoStore)
}

func TestMongo_SessionOptions(t *testing.T) {
	sessionOptions(t, newMongoStore)
}