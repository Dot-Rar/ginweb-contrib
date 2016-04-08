// Description: sessions
// Author: ZHU HAIHUA
// Since: 2016-04-07 22:20
package sessions

import (
	. "github.com/gin-gonic/contrib/sessions"
	"testing"
)

//const mysqlTestServer = "testuser:testpw@tcp(localhost:3306)/testdb?parseTime=true&loc=Local"
const mysqlTestServer = "xgsdk2_zhh:xgsdk2_zhh@tcp(10.20.122.235:3306)/xgsdk2_zhh?parseTime=true&loc=Local"

var newMySQLStore = func(_ *testing.T) Store {
	store, err := NewMySQLStore(mysqlTestServer, "sessionstore", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	return store
}

func TestMySQL_SessionGetSet(t *testing.T) {
	sessionGetSet(t, newMySQLStore)
}

func TestMySQL_SessionDeleteKey(t *testing.T) {
	sessionDeleteKey(t, newMySQLStore)
}

func TestMySQL_SessionFlashes(t *testing.T) {
	sessionFlashes(t, newMySQLStore)
}

func TestMySQL_SessionClear(t *testing.T) {
	sessionClear(t, newMySQLStore)
}

func TestMySQL_SessionOptions(t *testing.T) {
	sessionOptions(t, newMySQLStore)
}