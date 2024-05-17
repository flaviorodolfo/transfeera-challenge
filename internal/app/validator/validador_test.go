package validator

import "testing"

func TestValidarCpf(t *testing.T) {
	tests := map[string]struct {
		input  string
		result bool
	}{"vazio": {
		input:  "",
		result: false,
	},
		"cpf válido": {
			input:  "615.546.455-30",
			result: true,
		},
		"cpf inválido": {
			input:  "189.307.395-56",
			result: false,
		},
		"cpf não formatado": {
			input:  "18930739555",
			result: true,
		},
		"cnpj válido": {
			input:  "12.198.283/0001-07",
			result: false,
		},
		"cpf tipo especial invalido": {
			input:  "11111111111",
			result: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := ValidarCPF(tt.input)
			if result != tt.result {
				t.Errorf("esperado %v, obtido %v", tt.result, result)
			}
		})
	}
}

func TestValidarCnpj(t *testing.T) {
	tests := map[string]struct {
		input  string
		result bool
	}{"vazio": {
		input:  "",
		result: false,
	}, "cnpj válido": {
		input:  "41.916.896/0001-30",
		result: true,
	},
		"cnpj válido(sem formatacao)": {
			input:  "12198283000107",
			result: true,
		},
		"cnpj inválido ": {
			input:  "13.198.283/0001-07",
			result: false,
		},
		"cnpj inválido(sem formatacao)": {
			input:  "13198283000107",
			result: false,
		},
		"cnpj tipo especial invalido": {
			input:  "11111111111111",
			result: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := ValidarCNPJ(tt.input)
			if result != tt.result {
				t.Errorf("esperado %v, obtido %v", tt.result, result)
			}
		})
	}
}

func TestValidarChaveAleatoria(t *testing.T) {
	tests := map[string]struct {
		input  string
		result bool
	}{"vazio": {
		input:  "",
		result: false,
	},
		"chave válida": {
			input:  "46892703-d647-4a2c-a6be-a6e0f1488da7",
			result: true,
		},
		"valor aleatorio": {
			input:  "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			result: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := ValidarChaveAleatoria(tt.input)
			if result != tt.result {
				t.Errorf("esperado %v, obtido %v", tt.result, result)
			}
		})
	}
}

func TestValidarTelefone(t *testing.T) {
	tests := map[string]struct {
		input  string
		result bool
	}{"vazio": {
		input:  "",
		result: false,
	},
		"valor aleatorio": {
			input:  "46892703-d647-4a2c-a6be-a6e0f1488da7",
			result: false,
		},
		"numero válido(sem codigo postal)": {
			input:  "79992433805",
			result: true,
		},
		"numero válido(com codigo postal)": {
			input:  "5511998765432",
			result: true,
		},
		"numero válido(com + codigo postal)": {
			input:  "+5511998765432",
			result: true,
		},
		"numero inválido(sem ddd)": {
			input:  "998765432",
			result: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := ValidarTelefone(tt.input)
			if result != tt.result {
				t.Errorf("esperado %v, obtido %v", tt.result, result)
			}
		})
	}
}

func TestValidarEmail(t *testing.T) {
	tests := map[string]struct {
		input  string
		result bool
	}{"vazio": {
		input:  "",
		result: false,
	},
		"email válido": {
			input:  "flaviorodolfo@transfeera.com",
			result: true,
		},
		"email válido com tag": {
			input:  "flavio.rodolfo+transfeera@example.co.uk",
			result: true,
		},
		"email válido(domain-domain.com)": {
			input:  "user_name@example-domain.org",
			result: true,
		},
		"email sem dominio completo": {
			input:  "user@example",
			result: false,
		},
		"email formato errado": {
			input:  "user@.com",
			result: false,
		},
		"email incompleto": {
			input:  "user@com.",
			result: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := ValidarEmail(tt.input)
			if result != tt.result {
				t.Errorf("esperado %v, obtido %v", tt.result, result)
			}
		})
	}
}
