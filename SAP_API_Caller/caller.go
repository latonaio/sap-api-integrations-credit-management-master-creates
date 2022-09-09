package sap_api_caller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-credit-management-master-creates/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-credit-management-master-creates/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostCreditManagementMaster(
	businessPartner *requests.BusinessPartner,
	creditAccount *requests.CreditAccount,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "BusinessPartner":
			func() {
				c.BusinessPartner(businessPartner)
				wg.Done()
			}()
		case "CreditAccount":
			func() {
				c.CreditAccount(creditAccount)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) BusinessPartner(businessPartner *requests.BusinessPartner) {
	outputDataBusinessPartner, err := c.callCreditManagementMasterSrvAPIRequirementBusinessPartner("CreditMgmtBusinessPartner", businessPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataBusinessPartner)
}

func (c *SAPAPICaller) callCreditManagementMasterSrvAPIRequirementBusinessPartner(api string, businessPartner *requests.BusinessPartner) (*sap_api_output_formatter.BusinessPartner, error) {
	body, err := json.Marshal(businessPartner)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_CRDTMBUSINESSPARTNER", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToBusinessPartner(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) CreditAccount(creditAccount *requests.CreditAccount) {
	outputDataCreditAccount, err := c.callCreditManagementMasterSrvAPIRequirementCreditAccount("CreditManagementAccount", creditAccount)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataCreditAccount)
}

func (c *SAPAPICaller) callCreditManagementMasterSrvAPIRequirementCreditAccount(api string, creditAccount *requests.CreditAccount) (*sap_api_output_formatter.CreditAccount, error) {
	body, err := json.Marshal(creditAccount)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_CRDTMBUSINESSPARTNER", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToCreditAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
