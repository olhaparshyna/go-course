package main

//1. Створити вебсервер для перегляду інформації щодо класу школи.
//Користувач повинен мати можливість отримувати загальну статистику про клас.
//Додаткові вимоги:
//• інформація про учнів має зберігатися в оперативній пам'яті та бути доступною під час кожного запиту;
//• отримання інформації про учня має здійснюватись методом GET на адресі "/student/{id}",
//де {id} — унікальний ідентифікатор учня;
//дані можна отримати, лише якщо користувач є вчителем у цьому класі.

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	studentRes := &studentRes{
		storage: NewStorage(),
	}

	r.Handle("/student/{id}", auth(http.HandlerFunc(studentRes.getStudent), studentRes)).
		Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}

func auth(next http.Handler, res *studentRes) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		teacher, ok := res.storage.GetTeacherByUsername(username)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if teacher.Password != password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		studentId := mux.Vars(r)["id"]

		class, ok := res.storage.GetClassByStudentId(studentId)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if class.Id != teacher.ClassId {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})

}

type studentRes struct {
	storage *Storage
}

func (sr *studentRes) getStudent(w http.ResponseWriter, r *http.Request) {
	studentId := mux.Vars(r)["id"]

	student, ok := sr.storage.GetStudentById(studentId)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewEncoder(w).Encode(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
