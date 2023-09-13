package dpfm_api_processing_formatter

type ProcessingFormatterSDC struct {
	Header *Header `json:"Header"`
}

type ConversionProcessingKey struct {
	SystemConvertTo       string   `json:"SystemConvertTo"`
	SystemConvertFrom     string   `json:"SystemConvertFrom"`
	LabelConvertTo        string   `json:"LabelConvertTo"`
	LabelConvertFrom      string   `json:"LabelConvertFrom"`
	CodeConvertFromInt    *int     `json:"CodeConvertFromInt"`
	CodeConvertFromFloat  *float32 `json:"CodeConvertFromFloat"`
	CodeConvertFromString *string  `json:"CodeConvertFromString"`
	BusinessPartner       int      `json:"BusinessPartner"`
}

type ConversionProcessingCommonQueryGets struct {
	CodeConversionID      int      `json:"CodeConversionID"`
	SystemConvertTo       string   `json:"SystemConvertTo"`
	SystemConvertFrom     string   `json:"SystemConvertFrom"`
	LabelConvertTo        string   `json:"LabelConvertTo"`
	LabelConvertFrom      string   `json:"LabelConvertFrom"`
	CodeConvertFromInt    *int     `json:"CodeConvertFromInt"`
	CodeConvertFromFloat  *float32 `json:"CodeConvertFromFloat"`
	CodeConvertFromString *string  `json:"CodeConvertFromString"`
	CodeConvertToInt      *int     `json:"CodeConvertToInt"`
	CodeConvertToFloat    *float32 `json:"CodeConvertToFloat"`
	CodeConvertToString   *string  `json:"CodeConvertToString"`
	BusinessPartner       int      `json:"BusinessPartner"`
}

type PreparingHeaderPartner struct {
	PartnerFunction           string   `json:"PartnerFunction"`
	Customer                  *string  `json:"Customer"`
	Supplier                  *string  `json:"Supplier"`
	ConvertingBillToParty     *string  `json:"ConvertingBillToParty"`
	ConvertingPayer           *string  `json:"ConvertingPayer"`
	ConvertingDeliverToParty  *string  `json:"ConvertingDeliverToParty"`
	ConvertingPartnerFunction []string `json:"ConvertingPartnerFunction"`
	ConvertingCustomer        *string  `json:"ConvertingCustomer"`
	ConvertingSupplier        *string  `json:"ConvertingSupplier"`
}

type PreparingItemPricingElement struct {
	ConditionType           *string `json:"ConditionType"`
	ConvertingConditionType *string `json:"ConvertingConditionType"`
}

type Header struct {
	ConvertingOrderID                   string  `json:"ConvertingOrderID"`
	Sequence                            *string `json:"Sequence"`
	OrderType                           *string `json:"OrderType"`
	ConfirmationText                    *string `json:"ConfirmationText"`
	ConvertingConfirmationYieldQuantity *string `json:"ConvertingConfirmedYieldQuantity"`
	ConvertingConfirmedScrapQuantity    *string `json:"ConvertingConfirmedScrapQuantity"`
	Language                            *string `json:"Language"`
	ConvertingProduct                   *string `json:"ConvertingProduct"`
	Plant                               *string `json:"Plant"`
	WorkCenter                          *string `json:"WorkCenter"`
}

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

type Item struct {
	ConvertingOrderID                       string   `json:"ConvertingOrderID"`
	ConvertingOrderItem                     string   `json:"ConvertingOrderItem"`
	ConvertingOrderItemCategory             *string  `json:"ConvertingOrderItemCategory"`
	OrderItemTextBySeller                   *string  `json:"OrderItemTextBySeller"`
	ConvertingProduct                       *string  `json:"ConvertingProduct"`
	ConvertingProductGroup                  *string  `json:"ConvertingProductGroup"`
	BaseUnit                                *string  `json:"BaseUnit"`
	PricingDate                             *string  `json:"PricingDate"`
	PriceDetnExchangeRate                   *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                   *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                   *string  `json:"RequestedDeliveryTime"`
	ConvertingDeliverToParty                *string  `json:"ConvertingDeliverToParty"`
	DeliverFromParty                        *int     `json:"DeliverFromParty"`
	CreationDate                            *string  `json:"CreationDate"`
	LastChangeDate                          *string  `json:"LastChangeDate"`
	DeliverFromPlant                        *string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation         *string  `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                            *string  `json:"DeliveryUnit"`
	StockConfirmationBusinessPartner        *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch             *string  `json:"StockConfirmationPlantBatch"`
	OrderQuantityInBaseUnit                 *float32 `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit             *float32 `json:"OrderQuantityInDeliveryUnit"`
	ConfirmedOrderQuantityInBaseUnit        *float32 `json:"ConfirmedOrderQuantityInBaseUnit"`
	ItemWeightUnit                          *string  `json:"ItemWeightUnit"`
	ItemGrossWeight                         *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                           *float32 `json:"ItemNetWeight"`
	QuantityPerPackage                      *float32 `json:"QuantityPerPackage"`
	NetAmount                               *float32 `json:"NetAmount"`
	TaxAmount                               *float32 `json:"TaxAmount"`
	GrossAmount                             *float32 `json:"GrossAmount"`
	InvoiceDocumentDate                     *string  `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner          *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                         *string  `json:"ProductionPlant"`
	Incoterms                               *string  `json:"Incoterms"`
	TransactionTaxClassification            *string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   *string  `json:"ProductTaxClassificationBillToCountyr"`
	ProductTaxClassificationBillFromCountry *string  `json:"ProductTaxClassificationBillFromCountry"`
	AccountAssignmentGroup                  *string  `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup           *string  `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                            *string  `json:"PaymentTerms"`
	TaxCode                                 *string  `json:"TaxCode"`
	ItemBlockStatus                         *bool    `json:"ItemBlockStatus"`
	ItemBillingBlockStatus                  *bool    `json:"ItemBillingBlockStatus"`
	ItemDeliveryBlockStatus                 *bool    `json:"ItemDeliveryBlockStatus"`
	IsCancelled                             *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                     *bool    `json:"IsMarkedForDeletion"`
}

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

type ItemPricingElement struct {
	ConvertingOrderID                   string   `json:"ConvertingOrderID"`
	ConvertingOrderItem                 string   `json:"ConvertingOrderItem"`
	ConvertingBuyer                     *string  `json:"ConvertingBuyer"`
	Seller                              *int     `json:"Seller"`
	ConvertingConditionRecord           *string  `json:"ConvertingConditionRecord"`
	ConvertingConditionSequentialNumber *string  `json:"ConvertingConditionSequentialNumber"`
	ConditionType                       *string  `json:"ConditionType"`
	PricingDate                         *string  `json:"PricingDate"`
	ConditionRateValue                  *float32 `json:"ConditionRateValue"`
	ConditionCurrency                   *string  `json:"ConditionCurrency"`
	ConditionQuantity                   *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit               *string  `json:"ConditionQuantityUnit"`
	TaxCode                             *string  `json:"TaxCode"`
	ConditionAmount                     *float32 `json:"ConditionAmount"`
	TransactionCurrency                 *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged          *bool    `json:"ConditionIsManuallyChanged"`
}

type ConversionProcessingItemPricingElement struct {
	ConvertingConditionRecord           *string `json:"ConvertingConditionRecord"`
	ConvertedConditionRecord            *int    `json:"ConvertedConditionRecord"`
	ConvertingConditionSequentialNumber *string `json:"ConvertingConditionSequentialNumber"`
	ConvertedConditionSequentialNumber  *int    `json:"ConvertedConditionSequentialNumber"`
}

type ItemScheduleLine struct {
	ConvertingOrderID                     string   `json:"ConvertingOrderID"`
	ConvertingOrderItem                   string   `json:"ConvertingOrderItem"`
	ConvertingScheduleLine                string   `json:"ConvertingScheduleLine"`
	Product                               *string  `json:"Product"`
	StockConfirmationBussinessPartner     *int     `json:"StockConfirmationBussinessPartner"`
	StockConfirmationPlant                *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch           *string  `json:"StockConfirmationPlantBatch"`
	RequestedDeliveryDate                 *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                 *string  `json:"RequestedDeliveryTime"`
	ConfirmedDeliveryDate                 *string  `json:"ConfirmedDeliveryDate"`
	OriginalOrderQuantityInBaseUnit       *float32 `json:"OriginalOrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheck *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheck"`
	DeliveredQuantityInBaseUnit           *float32 `json:"DeliveredQuantityInBaseUnit"`
	OpenConfirmedQuantityInBaseUnit       *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
}

type ConversionProcessingItemScheduleLine struct {
	ConvertingScheduleLine *string `json:"ConvertingScheduleLine"`
	ConvertedScheduleLine  *int    `json:"ConvertedScheduleLine"`
}

type Address struct {
	ConvertingOrderID string `json:"ConvertingOrderID"`
}

type Partner struct {
	ConvertingOrderID         string  `json:"ConvertingOrderID"`
	ConvertingPartnerFunction string  `json:"ConvertingPartnerFunction"`
	ConvertingBusinessPartner *string `json:"ConvertingBusinessPartner"`
	ExternalDocumentID        *string `json:"ExternalDocumentID"`
}

type ConversionProcessingPartner struct {
	ConvertingPartnerFunction *string `json:"ConvertingPartnerFunction"`
	ConvertedPartnerFunction  *string `json:"ConvertedPartnerFunction"`
	ConvertingBusinessPartner *string `json:"ConvertingBusinessPartner"`
	ConvertedBusinessPartner  *int    `json:"ConvertedBusinessPartner"`
}
