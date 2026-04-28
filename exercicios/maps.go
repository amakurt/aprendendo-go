package exercicios

import "fmt"

// ExercicioMaps demonstra o uso de maps em Go
func ExercicioMaps() string {
	resultado := ""

	// Agenda de contatos
	agenda := make(map[string]string)
	agenda["Amauri"] = "85988014049"
	agenda["Valdilania"] = "85988687277"

	resultado += "=== Agenda de Contatos ===\n"
	for nome, tel := range agenda {
		resultado += fmt.Sprintf("  Contato: %s - Tel: %s\n", nome, tel)
	}

	// Padrão "Comma ok" - verificar se existe
	telefone, existe := agenda["Pedro"]
	if existe {
		resultado += fmt.Sprintf("  Pedro encontrado: %s\n", telefone)
	} else {
		resultado += "  Pedro não está na agenda.\n"
	}

	return resultado
}
