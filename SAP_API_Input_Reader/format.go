package sap_api_input_reader

import (
	"sap-api-integrations-credit-management-master-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToBusinessPartner() *requests.BusinessPartner {
	data := sdc.CreditManagementMaster
	return &requests.BusinessPartner{
		BusinessPartner:                data.BusinessPartner,
		CrdtMgmtBusinessPartnerGroup:   data.CrdtMgmtBusinessPartnerGroup,
		CreditWorthinessScoreValue:     data.CreditWorthinessScoreValue,
		CrdtWrthnssScoreValdtyEndDate:  data.CrdtWrthnssScoreValdtyEndDate,
		CrdtWorthinessScoreLastChgDate: data.CrdtWorthinessScoreLastChgDate,
		CalcdCrdtWorthinessScoreValue:  data.CalcdCrdtWorthinessScoreValue,
		CreditRiskClass:                data.CreditRiskClass,
		CalculatedCreditRiskClass:      data.CalculatedCreditRiskClass,
		CreditRiskClassLastChangeDate:  data.CreditRiskClassLastChangeDate,
		CreditCheckRule:                data.CreditCheckRule,
		CreditScoreAndLimitCalcRule:    data.CreditScoreAndLimitCalcRule,
	}
}

func (sdc *SDC) ConvertToCreditAccount() *requests.CreditAccount {
	dataCreditManagementMaster := sdc.CreditManagementMaster
	data := sdc.CreditManagementMaster.CreditAccount
	return &requests.CreditAccount{
		BusinessPartner:                dataCreditManagementMaster.BusinessPartner,
		CreditSegment:                  data.CreditSegment,
		BusinessPartnerIsCritical:      data.BusinessPartnerIsCritical,
		CreditAccountIsBlocked:         data.CreditAccountIsBlocked,
		CreditAccountBlockReason:       data.CreditAccountBlockReason,
		CreditAccountResubmissionDate:  data.CreditAccountResubmissionDate,
		CreditLimitAmount:              data.CreditLimitAmount,
		CreditLimitValidityEndDate:     data.CreditLimitValidityEndDate,
		CreditLimitLastChangeDate:      data.CreditLimitLastChangeDate,
		CreditLimitCalculatedAmount:    data.CreditLimitCalculatedAmount,
		CreditLimitIsZero:              data.CreditLimitIsZero,
		CreditLimitRequestedAmount:     data.CreditLimitRequestedAmount,
		CrdtLmtIsReqdFrmAutomCalc:      data.CrdtLmtIsReqdFrmAutomCalc,
		CreditLimitReqdValidityEndDate: data.CreditLimitReqdValidityEndDate,
		CreditLimitRequestDate:         data.CreditLimitRequestDate,
		CreditSegmentCurrency:          data.CreditSegmentCurrency,
	}
}
