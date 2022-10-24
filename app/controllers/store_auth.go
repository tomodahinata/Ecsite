package controllers

import (
	"Ecsite/app/models"
	"log"
	"net/http"
)

func StoreSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := StoreSession(w, r)
		if err != nil {
			genereateHTML(w, nil, "layout", "public_navbar", "store_signup")
		} else {
			http.Redirect(w, r, "/stores", http.StatusFound)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		store := models.Store{
			StoreName: r.PostFormValue("store_name"),
			Address:   r.PostFormValue("address"),
			Email:     r.PostFormValue("email"),
			PassWord:  r.PostFormValue("password"),
		}
		if err := store.CreateStore(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func StoreLogin(w http.ResponseWriter, r *http.Request) {
	_, err := StoreSession(w, r)
	if err != nil {
		log.Println(err)
		genereateHTML(w, nil, "layout", "public_navbar", "store_login")
	} else {
		http.Redirect(w, r, "/stores", http.StatusFound)
	}
}

func StoreAuthenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	store, err := models.GetStoreByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/store_login", http.StatusFound)
	}
	if store.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := store.CreateStoreSession()
		if err != nil {
			log.Println(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/store_login", http.StatusFound)
	}
}

func StoreLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.StoreSession{UUID: cookie.Value}
		session.DeleteStoreSessionByUUID()
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
