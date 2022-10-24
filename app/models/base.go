package models

import (
	"Ecsite/config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameBuyer    = "buyers"
	tableBuyerSession = "buyer_sessions"
	tableStoreSession = "store_sessions"
	tableNameStore    = "stores"
	tableNameGood     = "goods"
	tableNameCart     = "carts"
	tableNamePurchase = "purchases"
)

func init() {
	fmt.Println(config.Config.SQLDriver, config.Config.DbName)
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid STRING NOT NULL UNIQUE,
	buyer_name STRING,
	address  STRING,
	email STRING,
	password STRING,
	wallet INTEGER,
	created_at DATETIME)`, tableNameBuyer)
	Db.Exec(cmdU)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		buyer_id INTEGER,
		created_at DATETIME)`, tableBuyerSession)
	Db.Exec(cmdS)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		store_name STRING,
		address  STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameStore)
	Db.Exec(cmdT)

	cmdP := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		store_id INTEGER,
		created_at DATETIME)`, tableStoreSession)
	Db.Exec(cmdP)

	cmdG := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		store_id INTEGER,
		goods_name STRING,
		category INTEGER,
		price INTEGER,
		created_at DATETIME)`, tableNameGood)
	Db.Exec(cmdG)

	cmdC := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		buyer_id INTEGER,
		goods_id INTEGER,
		created_at DATETIME)`, tableNameCart)
	Db.Exec(cmdC)

	cmdB := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		buyer_id INTEGER,
		goods_id INTEGER,
		created_at DATETIME)`, tableNamePurchase)
	Db.Exec(cmdB)

}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
