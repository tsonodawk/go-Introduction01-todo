package controllers

import (
	"introduction/todoApp/app/models"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	sees, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		user, err := sees.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos

		generateHTML(w, user, "layout", "private_navbar", "index")

	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", http.StatusFound)
	}

}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		todo, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}

		generateHTML(w, todo, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}

		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		t := &models.Todo{ID: id}

		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}
