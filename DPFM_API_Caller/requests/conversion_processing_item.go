package requests

type ConversionProcessingItem struct {
	ConvertingOrderItem         *string `json:"ConvertingOrderItem"`
	ConvertedOrderItem          *int    `json:"ConvertedOrderItem"`
	ConvertingOrderItemCategory *string `json:"ConvertingOrderItemCategory"`
	ConvertedOrderItemCategory  *string `json:"ConvertedOrderItemCategory"`
	ConvertingProduct           *string `json:"ConvertingProduct"`
	ConvertedProduct            *string `json:"ConvertedProduct"`
	ConvertingProductGroup      *string `json:"ConvertingProductGroup"`
	ConvertedProductGroup       *string `json:"ConvertedProductGroup"`
	ConvertingDeliverToParty    *string `json:"ConvertingDeliverToParty"`
	ConvertedDeliverToParty     *int    `json:"ConvertedDeliverToParty"`
}
