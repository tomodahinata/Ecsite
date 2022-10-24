package controllers

import (
	"Ecsite/app/models"
	"log"
	"net/http"
	// "sort"
	// "strconv"
)

func BuyerSelect(w http.ResponseWriter, r *http.Request) {
	_, err := BuyerSession(w, r)
	if err != nil {
		genereateHTML(w, "Hello", "layout", "public_navbar", "select")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func BuyerTop(w http.ResponseWriter, r *http.Request) {
	_, err := BuyerSession(w, r)
	if err != nil {
		genereateHTML(w, "Hello", "layout", "public_navbar", "buyer_top")
	} else {
		http.Redirect(w, r, "/buyers", http.StatusFound)
	}
}

func BuyerIndex(w http.ResponseWriter, r *http.Request) {
	sess, err := BuyerSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		buyer, err := sess.GetBuyerBySession()
		if err != nil {
			log.Println(err)
		}

		// 商品情報の取得

		goods, err := models.GetGoods()
		if err != nil {
			log.Println(err)
		}

		// 商品名・カテゴリー・値段の取得
		for i := 0; i < len(goods); i++ {
			buyer, err := models.GetGood(goods[i].ID)
			if err != nil {
				log.Println(err)
			}
			goods[i].GoodName = buyer.GoodName
			goods[i].Category = buyer.Category
			goods[i].Price = buyer.Price
		}

		genereateHTML(w, buyer, "layout", "buyer_navbar", "buyer_index")
	}

}
