package dpfm_api_output_formatter

type Output struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	APIStatusCode     int    `json:"api_status_code"`
	RuntimeSessionID  string `json:"runtime_session_id"`
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
	APIType           string `json:"api_type"`
	//DataConcatenation   DataConcatenation `json:"DataConcatenation"`
	Header              Header   `json:"ProductionOrderConfirmation"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

//type DataConcatenation struct {
//	Header *Header `json:"OrdersHeader"`
//}

type Header struct {
	OrderID                   *string `json:"OrderID"`
	Sequence                  *string `json:"Sequence"`
	OrderType                 *string `json:"OrderType"`
	ConfirmationText          *string `json:"ConfirmationText"`
	ConfirmationYieldQuantity *string `json:"ConfirmationYieldQuantity"`
	ConfirmationScrapQuantity *string `json:"ConfirmationScrapQuantity"`
	ConfirmationGroup         *string `json:"ConfirmationGroup"`
	Language                  *string `json:"Language"`
	Material                  *string `json:"Material"`
	Plant                     *string `json:"Plant"`
	WorkCenter                *string `json:"WorkCenter"`
}
