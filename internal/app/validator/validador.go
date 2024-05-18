package validator

import (
	"regexp"
	"strconv"
)

func ValidarCPF(cpf string) bool {
	re := regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`)
	if !re.MatchString(cpf) {
		return false
	}
	re = regexp.MustCompile(`[^\d]`)
	cpf = re.ReplaceAllString(cpf, "")
	if valoresIguais(cpf) {
		return false
	}
	primeiroDv := calcularDigitoVerificadorCpf(cpf[:9], 10)
	segundoDv := calcularDigitoVerificadorCpf(cpf[:10], 11)
	return primeiroDv == int(cpf[9])-'0' && segundoDv == int(cpf[10])-'0'
}

func calcularDigitoVerificadorCpf(cpf string, multiplicador int) int {
	var soma int
	for i, s := range cpf {
		numero, _ := strconv.Atoi(string(s))
		soma += numero * (multiplicador - i)
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto

}

func valoresIguais(cpf string) bool {
	primeiroDigito := cpf[0]
	for i := 0; i < len(cpf); i++ {
		if cpf[i] != primeiroDigito {
			return false
		}
	}
	return true
}

func ValidarCNPJ(cnpj string) bool {
	re := regexp.MustCompile(`^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`)
	if !re.MatchString(cnpj) {
		return false
	}
	re = regexp.MustCompile(`[^\d]`)
	cnpj = re.ReplaceAllString(cnpj, "")
	if valoresIguais(cnpj) {
		return false
	}
	primeiroDv := calcularDigitoVerificadorCnpj(cnpj[:12], 5)
	segundoDv := calcularDigitoVerificadorCnpj(cnpj[:13], 6)
	return primeiroDv == int(cnpj[12])-'0' && segundoDv == int(cnpj[13])-'0'
}

func calcularDigitoVerificadorCnpj(cpf string, multiplicador int) int {
	var soma int
	for _, s := range cpf {
		numero, _ := strconv.Atoi(string(s))
		soma += numero * multiplicador
		multiplicador--
		if multiplicador < 2 {
			multiplicador = 9
		}
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto

}

func ValidarEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValidarTelefone(Telefone string) bool {
	re := regexp.MustCompile(`^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`)
	return re.MatchString(Telefone)
}

func ValidarChaveAleatoria(chave string) bool {
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return re.MatchString(chave)
}
