package validator

import (
	"regexp"
	"strconv"
)

func ValidarCPF(cpf string) bool {
	//verifica se o CPF está no formato padrão ###.###.###-## ou
	//########### onde # é um digito
	re := regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`)
	if !re.MatchString(cpf) {
		return false
	}
	//remove todos caracteres não digitos
	re = regexp.MustCompile(`[^\d]`)
	cpf = re.ReplaceAllString(cpf, "")
	if valoresIguais(cpf) {
		return false
	}

	primeiroDv := calcularDigitoVerificadorCpf(cpf[:9], 10)
	segundoDv := calcularDigitoVerificadorCpf(cpf[:10], 11)
	return primeiroDv == int(cpf[9])-'0' && segundoDv == int(cpf[10])-'0'
}

// Regra para calculo de digito de CPF
// 1. Multiplicar cada um dos 9 primeiros dígitos por uma sequência decrescente de pesos (10 a 2 para o primeiro dígito, 11 a 2 para o segundo).
// 2. Somar os resultados dessas multiplicações.
// 3. Calcular o módulo 11 da soma obtida.
//  4. Se o passo 3 for menor que 11 o digito verificador é 0.
//  5. Senão o resultado é o valor do passo 3 subtraido por 11.
//
// veja mais em: https://www.sefaz.pe.gov.br/Servicos/sintegra/Paginas/calculo-do-digito-verificador.aspx
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

// Recebe uma string x e retorna true se todos os caracteres de uma string são iguais
func valoresIguais(x string) bool {
	primeiroDigito := x[0]
	for i := 0; i < len(x); i++ {
		if x[i] != primeiroDigito {
			return false
		}
	}
	return true
}

func ValidarCNPJ(cnpj string) bool {
	//verifica se o CNPJ está no formato padrão ##.###.###/####-## ou
	//############## onde # é um digito
	re := regexp.MustCompile(`^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`)
	if !re.MatchString(cnpj) {
		return false
	}
	//remove todos não digitos
	re = regexp.MustCompile(`[^\d]`)
	cnpj = re.ReplaceAllString(cnpj, "")

	// cnpjs que possuem o mesmo digito durante sua estrutura são válidos de acorcdo com o calculo de digito mas
	// não devem ser considerados válidos.
	if valoresIguais(cnpj) {
		return false
	}
	primeiroDv := calcularDigitoVerificadorCnpj(cnpj[:12], 5)
	segundoDv := calcularDigitoVerificadorCnpj(cnpj[:13], 6)
	return primeiroDv == int(cnpj[12])-'0' && segundoDv == int(cnpj[13])-'0'
}

// calculo digito CNPJ:
//  1. Multiplicar cada um dos 12 primeiros dígitos por uma sequência de pesos (5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2 para o primeiro dígito,
//     e 6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2 para o segundo dígito).
//  2. Somar os resultados dessas multiplicações.
//  3. Calcular o módulo 11 da soma obtida.
//  4. Se o passo 3 for menor que 11 o digito verificador é 0.
//  5. Senão o resultado é o valor do passo 3 subtraido por 11.
//
// ver mais em : https://pt.wikipedia.org/wiki/D%C3%ADgito_verificador
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

// verifica se o email está em um formato valido
// ex user@example.com ou flaviorodolfo@transfeera.com.br
func ValidarEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Retorna true se o valor informado é um número de telefone válido nos formatos
// 55XXXXXXXXXXX ou XXXXXXXXXXX
func ValidarTelefone(Telefone string) bool {
	re := regexp.MustCompile(`^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`)
	return re.MatchString(Telefone)
}

// Retorna true se o valor está no formato UUID(Universally Unique Identifier)
// veja mais em https://pt.wikipedia.org/wiki/Identificador_%C3%BAnico_universal
func ValidarChaveAleatoria(chave string) bool {
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return re.MatchString(chave)
}
