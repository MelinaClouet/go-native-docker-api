package models

// Service représente un conteneur Docker à déployer dans un projet.
//
// Chaque service correspond à une instance Docker (ex: nginx, redis, API backend).
// Il contient les informations nécessaires pour créer et configurer un conteneur.
type Service struct {

	// Name est le nom du conteneur Docker.
	// Il sert aussi d'identifiant logique dans l'application.
	Name string `json:"name"`

	// Image est l'image Docker utilisée pour créer le conteneur.
	// Exemple : "nginx:latest", "redis:7"
	Image string `json:"image"`

	// Environment contient les variables d'environnement du conteneur.
	// Format : clé → valeur (ex: {"ENV": "dev"})
	Environment map[string]string `json:"environment"`

	// Volumes liste les volumes Docker à monter dans le conteneur.
	// Exemple : ["/host/path:/container/path"]
	Volumes []string `json:"volumes"`
}

// Project représente un projet complet composé de plusieurs services Docker.
//
// Dans la phase multi-engine, un projet est déployé sur un seul Engine cible.
type Project struct {

	// Name est le nom du projet.
	// Utilisé pour le suivi, la sauvegarde et les logs.
	Name string `json:"name"`

	// Services est la liste des conteneurs à déployer pour ce projet.
	Services []Service `json:"services"`

	// Engine définit le moteur Docker cible (local ou distant).
	// Permet le déploiement sur une machine spécifique dans la phase 2 (multi-engine).
	Engine Engine `json:"engine"`
}
