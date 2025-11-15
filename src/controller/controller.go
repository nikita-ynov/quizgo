package controller

import (
	"goquiz/pages"
	"goquiz/quizzes"
	"net/http"
	"strconv"
)

func renderPage(w http.ResponseWriter, filename string, data any) {
	err := pages.Temp.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, "Erreur rendu template : "+err.Error(), http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Message": "Controller Home !",
	}
	renderPage(w, "index.html", data)
}

func Quiz(w http.ResponseWriter, r *http.Request) {
	quizType := r.URL.Query().Get("type")
	stepStr := r.URL.Query().Get("step")
	answerStr := r.URL.Query().Get("answer")
	scoreStr := r.URL.Query().Get("score")

	step, _ := strconv.Atoi(stepStr)
	score, _ := strconv.Atoi(scoreStr)

	var q quizzes.Quiz

	switch quizType {
	case "info":
		q = quizzes.QuizInfo()
	case "cyber":
		q = quizzes.QuizCyber()
	case "ia":
		q = quizzes.QuizIAData()
	default:
		http.Error(w, "Quiz inconnu", http.StatusBadRequest)
		return
	}

	if step > 0 && answerStr != "" {
		answer, _ := strconv.Atoi(answerStr)
		if answer == q.Questions[step-1].Answer {
			score++
		}
	}

	if step >= len(q.Questions) {
		http.Redirect(w, r,
			"/score?score="+strconv.Itoa(score)+
				"&total="+strconv.Itoa(len(q.Questions)),
			http.StatusSeeOther,
		)
		return
	}

	data := map[string]any{
		"Quiz":     q,
		"Question": q.Questions[step],
		"Step":     step,
		"NextStep": step + 1,
		"Score":    score,
		"QuizType": quizType,
	}

	renderPage(w, "quiz.html", data)
}

func Score(w http.ResponseWriter, r *http.Request) {

	scoreStr := r.URL.Query().Get("score")
	totalStr := r.URL.Query().Get("total")

	score, _ := strconv.Atoi(scoreStr)
	total, _ := strconv.Atoi(totalStr)

	message := "Bien joué !"

	if score == total {
		message = "🔥 Score parfait ! Tu es un(e) crack !"
	} else if score >= total/2 {
		message = "😎 Pas mal du tout, continue comme ça !"
	} else {
		message = "😂 Aïe... On va dire que c'était l'échauffement !"
	}

	data := map[string]any{
		"Score":   score,
		"Total":   total,
		"Message": message,
	}

	renderPage(w, "score.html", data)
}
