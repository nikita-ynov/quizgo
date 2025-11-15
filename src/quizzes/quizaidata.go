package quizzes

func QuizIAData() Quiz {
	return Quiz{
		Title: "Quiz IA & Data",
		Questions: []Question{
			{
				"Que signifie IA (AI en anglais) ?",
				[]string{"Intelligence Artificielle", "Interface Automatisée", "Instruction Algorithmique"},
				0,
			},
			{
				"Quel langage est le plus utilisé en Data Science ?",
				[]string{"Python", "Java", "C++"},
				0,
			},
			{
				"Que signifie le terme 'dataset' ?",
				[]string{"Un ensemble de données", "Un logiciel d'analyse", "Un serveur de stockage"},
				0,
			},
			{
				"Quel est l’outil principal pour entraîner un modèle de Machine Learning ?",
				[]string{"Des données", "Un GPU obligatoire", "Une base SQL"},
				0,
			},
			{
				"Quel type d’IA reconnaît des images ?",
				[]string{"Vision par ordinateur", "Traitement du langage", "Réseaux sociaux"},
				0,
			},
			{
				"Que signifie NLP dans l’IA ?",
				[]string{"Natural Language Processing", "Neural Logic Program", "Numeric Learning Protocol"},
				0,
			},
			{
				"Quel algorithme est utilisé pour la classification ?",
				[]string{"KNN", "Dijkstra", "AES"},
				0,
			},
			{
				"Quel est le rôle d’un Data Scientist ?",
				[]string{"Analyser et modéliser des données", "Réparer des serveurs", "Coder des sites web"},
				0,
			},
			{
				"Quel type de modèle est ChatGPT ?",
				[]string{"Un modèle de langage", "Un antivirus", "Une base SQL"},
				0,
			},
			{
				"Que signifie 'Big Data' ?",
				[]string{"Un très grand volume de données", "Des données très lourdes physiquement", "Un gros ordinateur dédié"},
				0,
			},
		},
	}
}
