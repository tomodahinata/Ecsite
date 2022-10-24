package models

import (
	"log"
	"time"
)

type Buyer struct {
	ID        int
	UUID      string
	BuyerName string
	Address   string
	Wallet    int
	Email     string
	PassWord  string
	CreatedAt time.Time
}

type BuyerSession struct {
	ID        int
	UUID      string
	BuyerID   int
	Email     string
	CreatedAt time.Time
}

func (b *Buyer) CreateBuyer() (err error) {
	cmd := `insert into buyers(
		uuid,
		buyer_name,
		address,
		wallet,
		email,
		password,
		created_at) values (?,?,?,?,?,?,?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		b.BuyerName,
		b.Address,
		b.Wallet,
		b.Email,
		Encrypt(b.PassWord),
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
func GetBuyer(id int) (buyer Buyer, err error) {
	buyer = Buyer{}
	cmd := `select id, uuid, buyer_name,  address,wallet, email, password, created_at
	from buyers where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&buyer.ID,
		&buyer.UUID,
		&buyer.BuyerName,
		&buyer.Address,
		&buyer.Wallet,
		&buyer.Email,
		&buyer.PassWord,
		&buyer.CreatedAt,
	)
	return buyer, err
}

func (b *Buyer) UpdateBuyer() (err error) {
	cmd := `
		update buyers set buyer_name =?,  address=?, wallet=?,email =? 
	where 
		id =?`
	_, err = Db.Exec(cmd, b.BuyerName, b.Address, b.Wallet, b.Email, b.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (b *Buyer) DeleteBuyer() (err error) {
	cmd := `delte from buyers where id =?`
	_, err = Db.Exec(cmd, b.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetBuyerByEmail(email string) (buyer Buyer, err error) {
	buyer = Buyer{}
	cmd := `select id, uuid, buyer_name,address, wallet,email,password,created_at
	from buyers where email=?`
	err = Db.QueryRow(cmd, email).Scan(
		&buyer.ID,
		&buyer.UUID,
		&buyer.BuyerName,
		&buyer.Address,
		&buyer.Wallet,
		&buyer.Email,
		&buyer.PassWord,
		&buyer.CreatedAt)

	return buyer, err
}
func (b *Buyer) CreateBuyerSession() (buyer_session BuyerSession, err error) {
	buyer_session = BuyerSession{}
	cmd1 := `insert into buyer_sessions(
		uuid,
		email,
		buyer_id,
		created_at) values (?,?,?,?)`

	_, err = Db.Exec(cmd1, createUUID(), b.Email, b.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id,uuid,email,buyer_id,created_at
		from buyer_sessions where buyer_id=? and email=?`

	err = Db.QueryRow(cmd2, b.ID, b.Email).Scan(
		&buyer_session.ID,
		&buyer_session.UUID,
		&buyer_session.Email,
		&buyer_session.BuyerID,
		&buyer_session.CreatedAt)

	return buyer_session, err
}

func (sess *BuyerSession) CheckBuyerSession() (valid bool, err error) {
	cmd := `select id,uuid,email,buyer_id,created_at
	from buyer_sessions where uuid=?`
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.BuyerID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}
func (sess *BuyerSession) DeleteBuyerSessionByUUID() (err error) {
	cmd := `delete from buyer_sessions where uuid=?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
func (sess *BuyerSession) GetBuyerBySession() (buyer Buyer, err error) {
	buyer = Buyer{}
	cmd := `
			select id,uuid,buyer_name,address,wallet,email,created_at 
		FROM 
			buyers
		WHERE 
			id =?
	`
	err = Db.QueryRow(cmd, sess.BuyerID).Scan(
		&buyer.ID,
		&buyer.UUID,
		&buyer.BuyerName,
		&buyer.Address,
		&buyer.Wallet,
		&buyer.Email,
		&buyer.CreatedAt)
	return buyer, err
}
