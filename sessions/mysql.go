// Description: a session middleware based on mysql, gorilla-session and gin-gonic
// Author: ZHU HAIHUA
// Since: 2016-04-07 22:34
package sessions

import (
	ginsession "github.com/gin-gonic/contrib/sessions"
	"github.com/gorilla/sessions"
	"github.com/srinathgs/mysqlstore"
)

type MySQLStore interface {
	ginsession.Store
}

type mysqlStore struct {
	*mysqlstore.MySQLStore
}

// NewMySQLStore return a session store based on mongodb
func NewMySQLStore(endpoint, tableName, path string, maxAge int, keyPairs ...[]byte) (MySQLStore, error) {
	store, err := mysqlstore.NewMySQLStore(endpoint, tableName, path, maxAge, keyPairs)
	if err != nil {
		return nil, err
	}
	return &mysqlStore{store}
}

func (c *mysqlStore) Options(options ginsession.Options) {
	c.MySQLStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}