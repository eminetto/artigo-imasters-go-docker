package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
)

func hello() string {
	return "World"
}

type Produto struct {
	Codigo         string
	Uf             string
	Ex             int
	Descricao      string
	Nacional       float64
	Estadual       float64
	Importado      float64
	Municipal      float64
	Tipo           string
	VigenciaInicio string
	VigenciaFim    string
	Chave          string
	Versao         string
	Fonte          string
}

func getProduto(codigo string, uf string, ex string) (produto Produto, err error) {
	var result = Produto{}
	db, dbError := sql.Open("sqlite3", "./artigo.db")
	defer db.Close()
	checkErr(dbError)
	stmt, stmtError := db.Prepare("SELECT * FROM produto where codigo = ? and uf = ? and ex = ?")
	checkErr(stmtError)
	sqlError := stmt.QueryRow(codigo, uf, ex).Scan(&result.Codigo,
		&result.Uf, &result.Ex, &result.Descricao, &result.Nacional,
		&result.Estadual, &result.Importado, &result.Municipal,
		&result.Tipo, &result.VigenciaInicio, &result.VigenciaFim,
		&result.Chave, &result.Versao, &result.Fonte)
	switch {
	case sqlError == sql.ErrNoRows:
		return produto, errors.New("Produto not found")
	default:
		checkErr(sqlError)
	}

	return result, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	codigo := r.FormValue("codigo")
	uf := r.FormValue("uf")
	ex := r.FormValue("ex")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	produto, err := getProduto(codigo, uf, ex)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		returnError := map[string]string{"error": err.Error()}
		errorMessage, errJson := json.Marshal(returnError)
		checkErr(errJson)
		io.WriteString(w, string(errorMessage))
		return
	}
	result, err := json.Marshal(produto)
	checkErr(err)

	io.WriteString(w, string(result))
}
func main() {
	http.HandleFunc("/", HandleIndex)
	http.ListenAndServe(":8082", nil)
}
