package controllers

import (
	"Ecsite/app/models"
	"log"
	"net/http"
)

func BuyerSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := BuyerSession(w, r)
		if err != nil {
			genereateHTML(w, nil, "layout", "public_navbar", "buyer_signup")
		} else {
			http.Redirect(w, r, "/buyers", http.StatusFound)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		

		buyer := models.Buyer{
			BuyerName: r.PostFormValue("buyer_name"),
			Address:   r.PostFormValue("address"),
			Email:     r.PostFormValue("email"),
			PassWord:  r.PostFormValue("password"),
			Wallet:    0,
		}
		if err := buyer.CreateBuyer(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func BuyerLogin(w http.ResponseWriter, r *http.Request) {
	_, err := BuyerSession(w, r)
	if err != nil {
		genereateHTML(w, nil, "layout", "public_navbar", "buyer_login")
	} else {
		http.Redirect(w, r, "/buyers", http.StatusFound)
	}
}

func BuyerAuthenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	buyer, err := models.GetBuyerByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/buyer_login", http.StatusFound)
	}
	if buyer.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := buyer.CreateBuyerSession()
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
		http.Redirect(w, r, "/buyer_login", http.StatusFound)
	}
}

func BuyerLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.BuyerSession{UUID: cookie.Value}
		session.DeleteBuyerSessionByUUID()
	}
	http.Redirect(w, r, "/buyer_login", http.StatusFound)
}
