package controllers

import (
	// "Ecsite/app/models"
	"log"
	"net/http"
	// "sort"
	// "strconv"
)

// func StoreSelect(w http.ResponseWriter, r *http.Request) {
// 	_, err := StoreSession(w, r)
// 	if err != nil {
// 		genereateHTML(w, "Hello", "layout", "public_navbar", "select")
// 	} else {
// 		http.Redirect(w, r, "/", http.StatusFound)
// 	}
// }


func StoreTop(w http.ResponseWriter, r *http.Request) {
	_, err := StoreSession(w, r)
	if err != nil {
		genereateHTML(w, "Hello", "layout", "public_navbar", "store_top")
	} else {
		http.Redirect(w, r, "/stores", http.StatusFound)
	}
}

func StoreIndex(w http.ResponseWriter, r *http.Request) {
	sess, err := StoreSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		store, err := sess.GetStoreBySession()
		if err != nil {
			log.Println(err)
		}

		genereateHTML(w, store, "layout", "store_navbar", "store_index")
	}

}

func StoreNew(w http.ResponseWriter, r *http.Request) {
	_, err := StoreSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		genereateHTML(w, nil, "layout", "store_navbar", "store_new")
	}
}

func StoreSave(w http.ResponseWriter, r *http.Request) {
	sess, err := StoreSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/store_login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		store, err := sess.GetStoreBySession()
		if err != nil {
			log.Println(err)
		}
		goods_name := r.PostFormValue("good_name")
		category := r.PostFormValue("category")
		price := r.PostFormValue("price")
		if err := store.CreateGood(goods_name,  category,price); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/stores", http.StatusFound)

	}
}