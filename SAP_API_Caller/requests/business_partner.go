package requests

type BusinessPartner struct {
	BusinessPartner                string  `json:"BusinessPartner"`
	CrdtMgmtBusinessPartnerGroup   *string  `json:"CrdtMgmtBusinessPartnerGroup"`
	CreditWorthinessScoreValue     *string `json:"CreditWorthinessScoreValue"`
	CrdtWrthnssScoreValdtyEndDate  *string `json:"CrdtWrthnssScoreValdtyEndDate"`
	CrdtWorthinessScoreLastChgDate *string `json:"CrdtWorthinessScoreLastChgDate"`
	CalcdCrdtWorthinessScoreValue  *string `json:"CalcdCrdtWorthinessScoreValue"`
	CreditRiskClass                *string `json:"CreditRiskClass"`
	CalculatedCreditRiskClass      *string `json:"CalculatedCreditRiskClass"`
	CreditRiskClassLastChangeDate  *string `json:"CreditRiskClassLastChangeDate"`
	CreditCheckRule                *string `json:"CreditCheckRule"`
	CreditScoreAndLimitCalcRule    *string `json:"CreditScoreAndLimitCalcRule"`
}
