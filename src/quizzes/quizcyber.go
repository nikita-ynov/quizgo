package quizzes

func QuizCyber() Quiz {
	return Quiz{
		Title: "Quiz Cyber",
		Questions: []Question{
			{"Que signifie CPU ?", []string{"Un processeur", "Une carte graphique", "Un disque dur"}, 0},
			{"Linux est… ?", []string{"Un OS", "Un langage", "Un jeu vidéo"}, 0},
			{"Que signifie RAM ?", []string{"Mémoire vive", "Disque dur", "Carte réseau"}, 0},
			{"Le phishing consiste à… ?", []string{"Voler des informations", "Programmer un virus", "Installer un OS"}, 0},
			{"HTTPS sert à… ?", []string{"Sécuriser une connexion web", "Augmenter la vitesse d’un site", "Bloquer les publicités"}, 0},
			{"Un pare-feu (firewall) sert à… ?", []string{"Protéger un réseau", "Augmenter la mémoire", "Optimiser un processeur"}, 0},
			{"Quel protocole est utilisé pour envoyer des emails ?", []string{"SMTP", "HTTP", "FTP"}, 0},
			{"Quel est l’objectif d’un antivirus ?", []string{"Détecter et supprimer les logiciels malveillants", "Augmenter la vitesse", "Optimiser le GPU"}, 0},
			{"Que signifie VPN ?", []string{"Réseau privé virtuel", "Virus personnel", "Processeur graphique virtuel"}, 0},
			{"Quel est le rôle d’un mot de passe fort ?", []string{"Sécuriser un compte", "Accélérer un programme", "Nettoyer un disque dur"}, 0},
		},
	}
}
