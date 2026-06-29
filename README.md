# 📦 Go Native Docker API

API REST en Go permettant de gérer et déployer des projets Docker sur un ou plusieurs engines (local ou distant).

## 🚀 Fonctionnalités
* Création et stockage de projets (JSON local)
* Déploiement de conteneurs Docker
* Support multi-engine (local / VPS / remote Docker daemon)
* Pull + run automatique des images Docker
* Status des conteneurs
* Logs applicatifs (stdout + fichier)
* API REST avec Gin

## 🧱 Architecture
```bash
.
├── main.go
├── models/
├── services/
├── utils/
├── storage/
│   └── projects.json
├── logs/
│   └── app.log

```

## ⚙️ Prérequis
* Go 1.20+
* Docker installé
* Docker daemon accessible :

  * local : unix socket
  * distant : tcp://IP:2375 (optionnel multi-engine)

## 📥 Installation
```bash
git clone git@github.com:MelinaClouet/go-native-docker-api.git
cd go-native-docker-api
go mod tidy
```


## ▶️ Lancer l’API
1. Exécute:
```bash
go run main.go
```


2. L'API sera disponible sur : http://localhost:8080

## 🧪 Tester l’API (Postman)
📌 1. Obtenir les informations Docker
* Méthode : GET
* Endpoint : /docker-info

📌 2. Créer un projet
* Méthode : POST
* Endpoint : /projects
* Body :
```bash
{
  "name": "test-project",
  "engine": {
    "name": "local",
    "host": "unix:///var/run/docker.sock"
  },
  "services": [
    {
      "name": "nginx",
      "image": "nginx:latest",
      "environment": {
        "ENV": "dev"
      },
      "volumes": []
    }
  ]
}
```

📌 3. Lister les projets
* Méthode : GET
* Endpoint : /projects


📌 4. Déployer un projet
* Méthode : POST
* Endpoint : /deploy
* Body :

```bash
{
  "name": "test-project",
  "engine": {
    "name": "local",
    "host": "unix:///var/run/docker.sock"
  },
  "services": [
    {
      "name": "nginx",
      "image": "nginx:latest",
      "environment": {},
      "volumes": []
    }
  ]
}

```

📌 5. Status container
* Méthode : GET
* Endpoint : /status/nginx

🌐 Multi-Engine (Phase 2)

Tu peux déployer sur une machine distante Docker :

```bash 
"engine": {
   "name": "remote-server",
   "host": "tcp://192.168.1.50:2375"
}
```

⚠️ Exigences pour le serveur distant :

Docker doit être configuré pour le TCP listener :
```bash
sudo dockerd -H unix:///var/run/docker.sock -H tcp://0.0.0.0:2375
```

## 🪵 Logs

Les logs sont disponibles dans :

1. Terminal (stdout lorsque l'API est exécutée avec go run main.go).
2. Fichier : logs/app.log.


📁 Storage (mini base de données)

Les projets sont stockés ici :

storage/projects.json


## 🧠 Fonctionnement global
1. L'API reçoit un projet via un endpoint.
2. Sauvegarde du projet au format JSON.
3. Connexion au Docker engine (local ou distant).
4. Pull de l’image.
5. Création du conteneur.
6. Démarrage du conteneur.
7. Logs de toutes les actions.

## 🔥 Exemple workflow complet
1. Créer un projet :
POST /projects
2. Déployer le projet :
POST /deploy
3. Vérifier le status du conteneur :
GET /status/nginx


## 🧪 Exemple rapide test local
1. Lancer le conteneur hello-world avec Docker natif:
```bash 
docker run hello-world
```
2. Lancer l'API
```bash 
go run main.go
```
## 🚀 Améliorations possibles
* Retry automatique des pulls
* Rollback si container fail
* Dashboard web (React / Angular)
* Stockage PostgreSQL
* Logs JSON + Grafana Loki
* Authentification API
