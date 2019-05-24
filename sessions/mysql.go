// Description: a session middleware based on mysql, gorilla-session and gin-gonic
// Author: ZHU HAIHUA
// Since: 2016-04-07 22:34
package sessions

import (
	"github.com/TicketsBot/mysqlstore"
	ginsession "github.com/gin-gonic/contrib/sessions"
)

type MySQLStore interface {
	ginsession.Store
}

type mysqlStore struct {
	*mysqlstore.MySQLStore
}

// NewMySQLStore return a session store based on mysql
/* by default the maxAge is 0, you can change the value by
   setting session.Options:

	    session.Options(Options{
	        MaxAge: 3600,
	        Domain: localhost,
			Path: "/foo/bar/bat",
		})
*/
func NewMySQLStore(endpoint, tableName string, keyPairs ...[]byte) (MySQLStore, error) {
	store, err := mysqlstore.NewMySQLStore(endpoint, tableName, "/", 0, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &mysqlStore{store}, nil
}

func (c *mysqlStore) Options(options ginsession.Options) {
	c.Options(options)
}