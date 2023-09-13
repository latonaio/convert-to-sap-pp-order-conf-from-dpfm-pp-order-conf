package requests

type ConversionProcessingHeader struct {
	ConvertingOrderID     *string `json:"ConvertingOrderID"`
	ConvertedOrderID      *int    `json:"ConvertedOrderID"`
	ConvertingOrderType   *string `json:"ConvertingOrderType"`
	ConvertedOrderType    *string `json:"ConvertedOrderType"`
	ConvertingBuyer       *string `json:"ConvertingBuyer"`
	ConvertedBuyer        *int    `json:"ConvertedBuyer"`
	ConvertingBillToParty *string `json:"ConvertingBillToParty"`
	ConvertedBillToParty  *int    `json:"ConvertedBillToParty"`
	ConvertingPayer       *string `json:"ConvertingPayer"`
	ConvertedPayer        *int    `json:"ConvertedPayer"`
}
