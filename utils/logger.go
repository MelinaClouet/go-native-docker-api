package utils

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

// InitLogger initialise le logger global de l'application.
//
// Il configure un logger multi-sortie :
//   - stdout (terminal)
//   - fichier logs/app.log
//
// Ce logger est utilisé dans toute l'application pour assurer :
//   - traçabilité des actions
//   - debug en cas d'erreur
//   - audit basique des opérations
func InitLogger() {

	// -------------------------
	// CREATE LOG FILE
	// -------------------------

	// On s'assure que le dossier logs existe
	_ = os.MkdirAll("logs", 0755)

	file, err := os.OpenFile(
		"logs/app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		log.Fatal("[LOGGER] failed to open log file:", err)
	}

	// -------------------------
	// LOGGER CONFIGURATION
	// -------------------------

	Logger = log.New(
		io.MultiWriter(os.Stdout, file),
		"[GO-DOCKER-API] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Logger.Println("[LOGGER] initialized successfully")
}
