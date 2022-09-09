package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-credit-management-master-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToBusinessPartner(raw []byte, l *logger.Logger) (*BusinessPartner, error) {
	pm := &responses.BusinessPartner{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to BusinessPartner. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	businessPartner := &BusinessPartner{
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

	return businessPartner, nil
}

func ConvertToCreditAccount(raw []byte, l *logger.Logger) (*CreditAccount, error) {
	pm := &responses.CreditAccount{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CreditAccount. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	creditAccount := &CreditAccount{
		BusinessPartner:                data.BusinessPartner,
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

	return creditAccount, nil
}
