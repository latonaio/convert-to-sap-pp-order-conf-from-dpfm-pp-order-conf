package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf/DPFM_API_Input_Reader"
	"strconv"
)

func OutputFormatter(
	sdc *dpfm_api_input_reader.SDC,
	//psdc *dpfm_api_processing_formatter.ProcessingFormatterSDC,
	osdc *Output,
) error {
	//header := ConvertToHeader(*sdc, *psdc)
	header := ConvertToHeader(*sdc)

	//osdc.DataConcatenation = DataConcatenation{
	//	Header: header,
	//}

	osdc.Header = *header

	osdc.ServiceLabel = "SAPProductionOrderConfCreates"
	osdc.APISchema = "SAPProductionOrderConfCreates"
	osdc.APIProcessingResult = getBoolPtr(true)

	return nil
}

func float32ToString(f *float32) *string {
	if f != nil {
		str := strconv.FormatFloat(float64(*f), 'f', -1, 32)
		return &str
	}
	return nil
}

func ConvertToHeader(
	sdc dpfm_api_input_reader.SDC,
	// psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) *Header {
	orderId := "1000001"
	sequence := "0"
	orderType := "PP01"
	language := ""
	material := "5"
	plant := "0001"
	workCenter := ""
	confirmationGroup := "1"

	header := &Header{
		OrderID:                   &orderId,
		Sequence:                  &sequence,
		OrderType:                 &orderType,
		ConfirmationText:          sdc.Message.Header.ConfirmationText,
		ConfirmationYieldQuantity: float32ToString(sdc.Message.Header.ConfirmedYieldQuantity),
		ConfirmationScrapQuantity: float32ToString(sdc.Message.Header.ConfirmedScrapQuantity),
		ConfirmationGroup:         &confirmationGroup,
		Language:                  &language,
		Material:                  &material,
		Plant:                     &plant,
		WorkCenter:                &workCenter,
	}

	return header
}

func getBoolPtr(b bool) *bool {
	return &b
}
