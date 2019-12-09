package db

import (
	"github.com/globalsign/mgo"
)

var MgoSession Session

type Session struct {
	session *mgo.Session
}

func NewSession(url string) (*Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}
	return &Session{session}, err
}

func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

func (s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

func (s *Session) DropDatabase(db string) error {
	if s.session != nil {
		return s.session.DB(db).DropDatabase()
	}
	return nil
}
