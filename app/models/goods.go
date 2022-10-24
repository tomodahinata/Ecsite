package models

import (
	"log"
	"time"
)

type Good struct {
	ID        int
	StoreID   int
	GoodName  string
	Category  string
	Price     int
	CreatedAt time.Time
}

func (s *Store) CreateGood(goods_name string, category string, price string) (err error) {
	cmd := `insert into goods(
	goods_name
	category,
	price,
	store_id,
	created_at) values(?,?,?,?)`

	_, err = Db.Exec(cmd, goods_name, category, price, s.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetGood(id int) (good Good, err error) {
	cmd := `
		SELECT 
			id, store_id, goods_name,category,price, created_at 
		FROM 
			goods
		WHERE 
			id=?
			`
	good = Good{}

	err = Db.QueryRow(cmd, id).Scan(
		&good.ID,
		&good.StoreID,
		&good.GoodName,
		&good.Category,
		&good.Price,
		&good.CreatedAt)

	return good, err
}

func GetGoods() (goods []Good, err error) {
	cmd := `SELECT id, store_id,  goods_name,category,price created_at FROM goods`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var good Good
		err = rows.Scan(&good.ID,
			&good.StoreID,
			&good.GoodName,
			&good.Category,
			&good.Price,
			&good.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		goods = append(goods, good)
	}
	rows.Close()

	return goods, err
}

// func GetGoodsByStore() (goods []Good, err error) {
// 	cmd := `
// 		SELECT
// 			id,
// 			store_id,
// 			good_name,
// 			category,
// 			price,
// 			created_at
// 		FROM
// 			goods
// 		WHERE
// 			store_id=?
// 		ORDER BY created_at DESC
// 	`
// 	rows, err := Db.Query(cmd, ID)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	for rows.Next() {
// 		var good Good
// 		err = rows.Scan(
// 			&good.ID,
// 			&good.StoreID,
// 			&good.GoodName,
// 			&good.Category,
// 			&good.Price,
// 			&good.CreatedAt)

// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		goods = append(goods, good)
// 	}
// 	rows.Close()

// 	return goods, err
// }
