package router

import (
	"flag"
	"fmt"
	"go-mock/repo"
	"net/http"
	"strconv"
)

type App struct {
	DSN  string
	repo repo.Repo
}

func Router() http.Handler {
	app := App{}
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable timezone=UTC connect_timeout=5", "Posgtres connection")
	db, err := repo.ConnectToDb(app.DSN)
	if err != nil {
		panic(err)
	}
	app.repo, _ = repo.NewRepository(db)
	mux := http.NewServeMux()
	mux.Handle("/get-account/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId, _ := strconv.Atoi(r.FormValue("id"))
		person, err := app.repo.FindById(reqId)
		fmt.Println("1--->", person, reqId)
		if err != nil {
			http.Error(w, "storage error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%v", person)))
	}))
	return mux
}
