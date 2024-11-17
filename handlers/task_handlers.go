package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//Во всех обработчиках тип контента Content-Type — application/json.

func GetTasks(w http.ResponseWriter, _ *http.Request) {
	// здесь нужно вернуть все задачи
	// 200 при успешном выполнении
	// 500 при ошибке
	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// так как все успешно, то статус OK
	w.WriteHeader(http.StatusOK)
	// записываем сериализованные в JSON данные в тело ответа
	w.Write(resp)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	// здесь нужно получить данные из запроса, создать новую задачу и добавить ее в список
	// затем вернуть ответ со статусом 201 Created
	// 400 при ошибке
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	tasks[newTask.ID] = newTask
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// здесь нужно получить ID из URL и вернуть задачу с таким ID
	// 200 при успешном выполнении
	// 400 при ошибке
	id := chi.URLParam(r, "id")
	task, exists := tasks[id]
	if !exists {
		http.Error(w, "Task not found", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	// здесь нужно получить ID из URL и удалить задачу с таким ID
	// 200 при успешном выполнении
	// 400 при ошибке
	id := chi.URLParam(r, "id")
	if _, exists := tasks[id]; !exists {
		http.Error(w, "Task not found", http.StatusBadRequest)
		return
	}
	delete(tasks, id)
	w.WriteHeader(http.StatusOK)
}

// Task ...
type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}
