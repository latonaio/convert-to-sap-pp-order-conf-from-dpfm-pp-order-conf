# convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf

convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf は、周辺業務システム　を データ連携基盤 と統合することを目的に、SAP 受注データを データ連携基盤 オーダーデータに変換するマイクロサービスです。  
https://xxx.xxx.io/api/FUNCTION_ORDERS_SRV/creates/

## 動作環境

convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/FUNCTION_ORDERS_SRV/creates/

## 本レポジトリ に 含まれる API名
convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf には、次の API をコールするためのリソースが含まれています。  

* A_Header（オーダー - ヘッダデータ）
* A_Item（オーダー - 明細データ）
* A_ItemPricingElement（オーダー - 明細取引先プラントデータ）
* A_ItemScheduleLine（オーダー - 明細納入日程行データ）
* A_Address（オーダー - 住所データ）
* A_Partner（オーダー - 取引先データ）

## API への 値入力条件 の 初期値
convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "DPFMOrdersCreates",
	"accepter": ["Header"],
	"order_id": null,
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMOrdersCreates",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncOrderCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,

	log *logger.Logger,
	// msg rabbitmq.RabbitmqMessage,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	convertFin := make(chan error)

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, convertFin, log, &errs, input, output)
		case "Item":
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}
    
	ticker := time.NewTicker(10 * time.Second)
	select {
	case e := <-convertFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		mtx.Lock()
		errs = append(errs, xerrors.Errorf("time out"))
		return errs
	}

	return nil
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は SAP 受注データが データ連携基盤 オーダーデータ に変換された結果の JSON の例です。  
以下の項目のうち、"OrderID" ～ "HeaderText" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。  

```
{
  "connection_key": "response",
  "result": true,
  "redis_key": "abcdefg",
  "filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
  "api_status_code": 200,
  "runtime_session_id": "",
  "business_partner": 201,
  "service_label": "FUNCTION_ORDERS",
  "message": {
    "Header": {
      "OrderID": null,
      "OrderDate": "2022-09-17",
      "OrderType": "",
      "Buyer": null,
      "Seller": null,
      "CreationDate": null,
      "LastChangeDate": null,
      "ContractType": "",
      "ValidityStartDate": null,
      "ValidityEndDate": null,
      "InvoiceScheduleStartDate": null,
      "InvoiceScheduleEndDate": null,
      "TotalNetAmount": null,
      "TotalTaxAmount": null,
      "TotalGrossAmount": null,
      "OverallDeliveryStatus": "",
      "TotalBlockStatus": null,
      "OverallOrdReltdBillgStatus": "",
      "OverallDocReferenceStatus": "",
      "TransactionCurrency": "JPY",
      "PricingDate": "2022-11-26",
      "PriceDetnExchangeRate": 0.01,
      "RequestedDeliveryDate": "2022-12-12",
      "HeaderCompleteDeliveryIsDefined": null,
      "HeaderBillingBlockReason": null,
      "DeliveryBlockReason": null,
      "Incoterms": "CIF",
      "PaymentTerms": "0001",
      "PaymentMethod": "T",
      "ReferenceDocument": null,
      "ReferenceDocumentItem": null,
      "BPAccountAssignmentGroup": "",
      "AccountingExchangeRate": 0.01,
      "BillingDocumentDate": "2022-12-31",
      "IsExportImportDelivery": null,
      "HeaderText": ""
    }
  },
  "api_schema": "SAPSalesOrderCreates",
  "accepter": [
    "All"
  ],
  "order_id": null,
  "deleted": false
}
```