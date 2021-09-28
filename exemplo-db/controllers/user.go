package controllers

import (
	"encoding/json"
	"exemplo-db/db"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var u user
	if err = json.Unmarshal(body, &u); err != nil {
		w.Write([]byte("Erro ao converter usuário pra struct!"))
		return
	}

	db, err := db.ConnectToDB()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into users (name, email) values ($1, $2)")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(u.Name, u.Email)
	if err != nil {
		w.Write([]byte("Erro ao salvar usuário no banco de dados."))
		return
	}

	id, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter id inserido."))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", id)))
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	db, err := db.ConnectToDB()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Erro ao buscar usuários!"))
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var u user

		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuários!"))
			return
		}

		users = append(users, u)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

func FindById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	ID, err := strconv.ParseUint(param["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, err := db.ConnectToDB()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	row, err := db.Query("select * from users where id=$1", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}
	defer row.Close()

	var user user
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuários!"))
			return
		}
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {}
