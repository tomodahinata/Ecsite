package models

import (
	"log"
	"time"
)

type Store struct {
	ID        int
	UUID      string
	StoreName string
	Address   string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

type StoreSession struct {
	ID        int
	UUID      string
	StoreID    int
	Email     string
	CreatedAt time.Time
}



func (s *Store) CreateStore() (err error) {
	cmd := `insert into stores(
		uuid,
		store_name,
		address,
		email,
		password,
		created_at) values (?,?,?,?,?,?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		s.StoreName,
		s.Address,
		s.Email,
		Encrypt(s.PassWord),
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetStore(id int) (store Store, err error) {
	store = Store{}
	cmd := `select id, uuid,store_name , address, email, password, created_at
	from stores where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&store.ID,
		&store.UUID,
		&store.StoreName,
		&store.Address,
		&store.Email,
		&store.PassWord,
		&store.CreatedAt,
	)
	return store, err
}

func (s *Store) UpdateStore() (err error) {
	cmd := `
		update stores set store_name=?,  address=?, email =? 
	where 
		id =?`
	_, err = Db.Exec(cmd, s.StoreName, s.Address,  s.Email, s.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (s *Store) DeleteStore() (err error) {
	cmd := `delte from stores where id =?`
	_, err = Db.Exec(cmd, s.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetStoreByEmail(email string) (store Store, err error) {
	store = Store{}
	cmd := `select id, uuid, store_name,address, email,password,created_at
	from stores where email=?`
	err = Db.QueryRow(cmd, email).Scan(
		&store.ID,
		&store.UUID,
		&store.StoreName,
		&store.Address,
		&store.Email,
		&store.PassWord,
		&store.CreatedAt)

	return store, err
}
func (s *Store) CreateStoreSession() (store_session StoreSession, err error) {
	store_session = StoreSession{}
	cmd1 := `insert into store_sessions(
		uuid,
		email,
		store_id,
		created_at) values (?,?,?,?)`

	_, err = Db.Exec(cmd1, createUUID(), s.Email, s.ID, time.Now())
	if err != nil {
		log.Println(err)
	}
	cmd2 := `select id,uuid,email,store_id,created_at
		from store_sessions where store_id=? and email=?`

	err = Db.QueryRow(cmd2, s.ID, s.Email).Scan(
		&store_session.ID,
		&store_session.UUID,
		&store_session.Email,
		&store_session.StoreID,
		&store_session.CreatedAt)

	return store_session, err
}

func (sess *StoreSession) CheckStoreSession() (valid bool, err error) {
	cmd := `select id,uuid,email,store_id,created_at
	from store_sessions where uuid=?`

	
	log.Println(sess.UUID)
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.StoreID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	log.Println(valid)
	log.Println("＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝")
	log.Println(err)

	return valid, err
}
func (sess *StoreSession) DeleteStoreSessionByUUID() (err error) {
	cmd := `delete from store_sessions where uuid=?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
func (sess *StoreSession) GetStoreBySession() (store Store, err error) {
	store = Store{}
	cmd := `
			select id,uuid,store_name,address,email,created_at 
		FROM 
			stores
		WHERE 
			id =?
	`
	err = Db.QueryRow(cmd, sess.StoreID).Scan(
		&store.ID,
		&store.UUID,
		&store.StoreName,
		&store.Address,
		&store.Email,
		&store.CreatedAt)
	return store, err
}
