// Description: sessions
// Author: ZHU HAIHUA
// Since: 2016-04-07 22:20
package sessions

import (
	ginsession "github.com/gin-gonic/contrib/sessions"
	"gopkg.in/mgo.v2"
	"testing"
)

const mongoTestServer = "localhost:6379"

var newRedisStore = func(_ *testing.T) ginsession.Store {
	dbsess, err := mgo.Dial(mongoTestServer)
	store := NewMongoStore(dbsess, "session", "session", 3600, true, []byte("secret"))
	return store
}
