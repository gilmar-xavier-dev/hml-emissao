/*
 * API Emissão server
 *
 * API para integração da paySmart com emissor para consulta e consumo de saldo e/ou limite para cartões. Esta especificação descreve os serviços que serão consumidos pela paySmart em uma API cuja implementação é de responsabilidade do emissor. As funções serão chamadas na sequência da interação entre a bandeira e paySmart no fluxo de pagamentos, após serem realizadas todas as validações de segurança por parte da bandeira e da processadora respectivamente.<p><strong>Tempos de resposta:</strong></p><ul><li><strong>< 2 segundos: Ideal</strong></li><li><strong>2 a 4 segundos: Começa a gerar alarmes</strong></li><li><strong>4 a 7 segundos: Baixa probabilidade de autorização</strong></li><li><strong>> 7 segundos: Timeout</strong></li></ul>
 *
 * API version: v1-oas3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type Success struct {
	Message string `json:"message"`
	// Deve-se retornar sempre código \"0\" em caso de sucesso.
	Code int32 `json:"code"`
	// Código de autorização de 6 dígitos gerado pelo emissor.
	AuthorizationId int32 `json:"authorization_id"`

	Balance *Amount `json:"balance,omitempty"`
}
