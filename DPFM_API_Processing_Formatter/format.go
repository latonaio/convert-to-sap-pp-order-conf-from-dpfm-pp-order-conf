package dpfm_api_processing_formatter

import (
	"context"
	dpfm_api_input_reader "convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf/DPFM_API_Input_Reader"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	//"golang.org/x/xerrors"
)

type ProcessingFormatter struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewProcessingFormatter(ctx context.Context, db *database.Mysql, l *logger.Logger) *ProcessingFormatter {
	return &ProcessingFormatter{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (p *ProcessingFormatter) ProcessingFormatter(
	sdc *dpfm_api_input_reader.SDC,
	psdc *ProcessingFormatterSDC,
) error {
	//var err error
	//var e error

	//if bpIDIsNull(sdc) {
	//	return xerrors.New("business_partner is null")
	//}
	//
	//wg := sync.WaitGroup{}
	//
	//// Ref: PreparingHeaderPartner
	//psdc.Header, err = p.Header(sdc, psdc)
	//if err != nil {
	//	return err
	//}
	//
	//wg.Add(1)
	//go func(wg *sync.WaitGroup) {
	//	defer wg.Done()
	//	// Ref: Header
	//	//psdc.ConversionProcessingHeader, e = p.ConversionProcessingHeader(sdc, psdc)
	//	if e != nil {
	//		err = e
	//		return
	//	}
	//}(&wg)
	//
	//wg.Wait()
	//if err != nil {
	//	return err
	//}
	//
	//p.l.Info(psdc)

	return nil
}

//func (p *ProcessingFormatter) Header(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) (*Header, error) {
//	data := sdc.Header
//	dataPreparingHeaderPartner := psdc.PreparingHeaderPartner
//
//	systemDate := getSystemDatePtr()
//
//	res := Header{
//		ConvertingOrderID:     data.SalesOrder,
//		OrderDate:             data.SalesOrderDate,
//		ConvertingOrderType:   data.SalesOrderType,
//		ConvertingBuyer:       data.SoldToParty,
//		Seller:                sdc.BusinessPartnerID,
//		ConvertingBillToParty: dataPreparingHeaderPartner.ConvertingBillToParty,
//		BillFromParty:         sdc.BusinessPartnerID,
//		ConvertingPayer:       dataPreparingHeaderPartner.ConvertingPayer,
//		Payee:                 sdc.BusinessPartnerID,
//		CreationDate:          systemDate,
//		LastChangeDate:        systemDate,
//	}
//
//	return &res, nil
//}
//
//func (p *ProcessingFormatter) ConversionProcessingHeader(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) (*ConversionProcessingHeader, error) {
//	dataKey := make([]*ConversionProcessingKey, 0)
//
//	p.appendDataKey(
//		&dataKey, sdc,
//		"OrderID", "OrderID",
//		psdc.Header.ConvertingOrderID,
//	)
//	p.appendDataKey(
//		&dataKey, sdc,
//		"SalesOrderType", "ConfirmedYieldQuantity",
//		psdc.Header.ConvertingConfirmationYieldQuantity,
//	)
//	p.appendDataKey(&dataKey, sdc, "SoldToParty", "Buyer", psdc.Header.ConvertingBuyer)
//	p.appendDataKey(&dataKey, sdc, "Customer", "BillToParty", psdc.Header.ConvertingBillToParty)
//	p.appendDataKey(&dataKey, sdc, "Customer", "Payer", psdc.Header.ConvertingPayer)
//
//	dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
//	if err != nil {
//		return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
//	}
//
//	data, err := p.ConvertToConversionProcessingHeader(dataKey, dataQueryGets)
//	if err != nil {
//		return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
//	}
//
//	return data, nil
//}
//
//func (psdc *ProcessingFormatter) ConvertToConversionProcessingHeader(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingHeader, error) {
//	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
//	for _, v := range conversionProcessingCommonQueryGets {
//		data[v.LabelConvertTo] = v
//	}
//
//	for _, v := range conversionProcessingKey {
//		if _, ok := data[v.LabelConvertTo]; !ok {
//			return nil, xerrors.Errorf("Value of %s is not in the database", v.LabelConvertTo)
//		}
//	}
//
//	res := &ConversionProcessingHeader{}
//
//	if _, ok := data["OrderID"]; ok {
//		res.ConvertingOrderID = data["OrderID"].CodeConvertFromString
//		res.ConvertedOrderID = data["OrderID"].CodeConvertToInt
//	}
//	if _, ok := data["OrderType"]; ok {
//		res.ConvertingOrderType = data["OrderType"].CodeConvertFromString
//		res.ConvertedOrderType = data["OrderType"].CodeConvertToString
//	}
//
//	return res, nil
//}
//
//func (p *ProcessingFormatter) appendDataKey(
//	dataKey *[]*ConversionProcessingKey,
//	sdc *dpfm_api_input_reader.SDC,
//	labelConvertFrom string,
//	labelConvertTo string,
//	codeConvertFrom any,
//) {
//	switch v := codeConvertFrom.(type) {
//	case int, float32:
//		if v == 0 {
//			return
//		}
//	case string:
//		if v == "" {
//			return
//		}
//	case *int, *float32:
//		// if v == nil || *v == 0 {
//		if v == nil {
//			return
//		}
//	case *string:
//		if v == nil || *v == "" {
//			return
//		}
//	default:
//		return
//	}
//	*dataKey = append(*dataKey, p.ConversionProcessingKey(sdc, labelConvertFrom, labelConvertTo, codeConvertFrom))
//}
//
//func bpIDIsNull(sdc *dpfm_api_input_reader.SDC) bool {
//	return sdc.BusinessPartnerID == nil
//}
