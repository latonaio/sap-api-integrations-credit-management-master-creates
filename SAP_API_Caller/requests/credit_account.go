package requests

type CreditAccount struct {
	BusinessPartner                string  `json:"BusinessPartner"`
	CreditSegment                  string  `json:"CreditSegment"`
	BusinessPartnerIsCritical      *bool   `json:"BusinessPartnerIsCritical"`
	CreditAccountIsBlocked         *bool   `json:"CreditAccountIsBlocked"`
	CreditAccountBlockReason       *string `json:"CreditAccountBlockReason"`
	CreditAccountResubmissionDate  *string `json:"CreditAccountResubmissionDate"`
	CreditLimitAmount              *string `json:"CreditLimitAmount"`
	CreditLimitValidityEndDate     *string `json:"CreditLimitValidityEndDate"`
	CreditLimitLastChangeDate      *string `json:"CreditLimitLastChangeDate"`
	CreditLimitCalculatedAmount    *string `json:"CreditLimitCalculatedAmount"`
	CreditLimitIsZero              *bool   `json:"CreditLimitIsZero"`
	CreditLimitRequestedAmount     *string `json:"CreditLimitRequestedAmount"`
	CrdtLmtIsReqdFrmAutomCalc      *bool   `json:"CrdtLmtIsReqdFrmAutomCalc"`
	CreditLimitReqdValidityEndDate *string `json:"CreditLimitReqdValidityEndDate"`
	CreditLimitRequestDate         *string `json:"CreditLimitRequestDate"`
	CreditSegmentCurrency          *string `json:"CreditSegmentCurrency"`
}
