package utils

import "fmt"

// MapToSlice convertit une map de variables d'environnement en slice de strings.
//
// Docker attend les variables d'environnement sous le format :
//
//	["KEY=value", "KEY2=value2"]
//
// Cette fonction est utilisée lors de la création de conteneurs
// pour injecter les variables dans le runtime Docker.
func MapToSlice(env map[string]string) []string {

	var result []string

	// Parcours des variables d'environnement
	for k, v := range env {

		// On ignore les clés ou valeurs vides pour éviter des erreurs Docker
		if k != "" && v != "" {
			result = append(result, fmt.Sprintf("%s=%s", k, v))
		}
	}

	return result
}
