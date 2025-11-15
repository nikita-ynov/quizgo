package quizzes

func QuizInfo() Quiz {
	return Quiz{
		Title: "Quiz Développement",
		Questions: []Question{
			{"Que signifie HTML ?", []string{"Langage de balisage pour les pages web", "Langage de programmation serveur", "Base de données"}, 0},
			{"Que signifie CSS ?", []string{"Feuilles de style pour les pages web", "Langage de programmation", "Serveur web"}, 0},
			{"JavaScript est utilisé pour… ?", []string{"Rendre les pages web interactives", "Stocker les données", "Gérer les serveurs"}, 0},
			{"Git sert à… ?", []string{"Gérer les versions de code", "Compiler le code", "Déployer des applications"}, 0},
			{"Que signifie API ?", []string{"Interface de programmation d’applications", "Application serveur interne", "Outil de test logiciel"}, 0},
			{"Que fait un framework ?", []string{"Fournit une structure pour le développement", "Exécute le code automatiquement", "Optimise la mémoire"}, 0},
			{"Quel langage est principalement utilisé pour le développement iOS ?", []string{"Swift", "Java", "Python"}, 0},
			{"SQL sert à… ?", []string{"Interroger et gérer des bases de données", "Styler une page web", "Compiler du code"}, 0},
			{"Que fait un debugger ?", []string{"Aide à trouver et corriger les erreurs dans le code", "Compile le code plus rapidement", "Déploie l’application sur le serveur"}, 0},
			{"Que signifie OOP (POO en français) ?", []string{"Programmation orientée objet", "Optimisation opérationnelle", "Open source programming"}, 0},
		},
	}
}
