package controllers

import (
	"Ecsite/app/models"
	"Ecsite/config"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

func genereateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func BuyerSession(w http.ResponseWriter, r *http.Request) (sess models.BuyerSession, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.BuyerSession{UUID: cookie.Value}
		if ok, _ := sess.CheckBuyerSession(); !ok {
			err = fmt.Errorf("inavalid session")
		}
	}
	return sess, err
}

func StoreSession(w http.ResponseWriter, r *http.Request) (sess models.StoreSession, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.StoreSession{UUID: cookie.Value}
		ok, _ := sess.CheckStoreSession()
		log.Println(ok)
		if !ok {

			err = fmt.Errorf("inavalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/ecsites/()/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

func StarMainSerever() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// バイヤー関連
	http.HandleFunc("/", BuyerSelect)
	http.HandleFunc("/buyer", BuyerTop)
	http.HandleFunc("/buyer_signup", BuyerSignup)
	http.HandleFunc("/buyer_login", BuyerLogin)
	http.HandleFunc("/buyer_authenticate", BuyerAuthenticate)
	http.HandleFunc("/buyer_logout", BuyerLogout)
	http.HandleFunc("/buyers", BuyerIndex)

	// ストア関連
	// http.HandleFunc("/", StoreSelect)
	http.HandleFunc("/store", StoreTop)
	http.HandleFunc("/store_signup", StoreSignup)
	http.HandleFunc("/store_login", StoreLogin)
	http.HandleFunc("/store_authenticate", StoreAuthenticate)
	http.HandleFunc("/store_logout", StoreLogout)
	http.HandleFunc("/stores", StoreIndex)
	http.HandleFunc("/stores/new", StoreNew)
	http.HandleFunc("/stores/save", StoreSave)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
