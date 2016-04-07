// Description: a session middleware based on mongodb, gorilla-session and gin-gonic
// Author: ZHU HAIHUA
// Since: 2016-04-07 20:20
package sessions

import (
	ginsession "github.com/gin-gonic/contrib/sessions"
	"github.com/gorilla/sessions"
	"github.com/kidstuff/mongostore"
	"gopkg.in/mgo.v2"
)

type MongoStore interface {
	ginsession.Store
}

type mongoStore struct {
	*mongostore.MongoStore
}

// NewMongoStore return a session store based on mongodb
//
func NewMongoStore(s *mgo.Session, dbName, collectionName string, maxAge int, ensureTTL bool, keyPairs ...[]byte) MongoStore {
	store := mongostore.NewMongoStore(s.DB(dbName).C(collectionName), maxAge, ensureTTL, keyPairs...)
	return &mongoStore{store}
}

func (c *mongoStore) Options(options ginsession.Options) {
	c.MongoStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
