package quizzes

func QuizCyber() Quiz {
	return Quiz{
		Title: "Quiz Cyber",
		Questions: []Question{
			{"Que signifie CPU ?", []string{"Un processeur", "Une carte graphique", "Un disque dur"}, 0},
			{"Linux est… ?", []string{"Un OS", "Un langage", "Un jeu vidéo"}, 0},
		},
	}
}
