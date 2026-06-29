package models

// Engine représente une cible Docker sur laquelle les conteneurs peuvent être déployés.
//
// Dans le cadre de l'architecture multi-engine (Phase 2),
// un Engine correspond à une machine Docker locale ou distante.
//
// Exemple :
//   - host local : unix:///var/run/docker.sock
//   - host distant : tcp://192.168.1.10:2375
type Engine struct {

	// Name est un identifiant lisible du moteur Docker.
	// Il permet de différencier plusieurs environnements (local, staging, prod, etc.).
	Name string `json:"name"`

	// Host est l'URL de connexion au daemon Docker.
	// Il peut être :
	//   - un socket local (unix:///var/run/docker.sock)
	//   - une connexion TCP distante (tcp://IP:2375)
	Host string `json:"host"`
}
