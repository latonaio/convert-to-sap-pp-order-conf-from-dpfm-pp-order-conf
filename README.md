# convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf

convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf は、周辺業務システム　を データ連携基盤 と統合することを目的に、データ連携基盤 製造指図確認データ を SAP 製造指図確認データ に変換するマイクロサービスです。  
https://xxx.xxx.io/api/FUNCTION_CONVERT_TO_SAP_PRODUCTION_ORDER_CONF_FROM_DPFM_PRODUCTION_ORDER_CONF/function/

## 動作環境

convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  

## 本レポジトリ が 対応する API サービス
convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/FUNCTION_CONVERT_TO_SAP_PRODUCTION_ORDER_CONF_FROM_DPFM_PRODUCTION_ORDER_CONF/function/

## 本レポジトリ に 含まれる API名
convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf には、次の API をコールするためのリソースが含まれています。  

* A_Header（製造指図確認 - ヘッダデータ）

## API への 値入力条件 の 初期値
convert-to-sap-pp-order-conf-from-dpfm-pp-order-conf において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "FunctionConvertToSAPProductionOrderConfFromDPFMProductionOrderConf",
	"accepter": ["Header"],
	"order_id": null,
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "FunctionConvertToSAPProductionOrderConfFromDPFMProductionOrderConf",
	"accepter": ["All"],
	"order_id": null,
	"deleted": false
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は データ連携基盤 製造指図 ヘッダ データが sap 製造指図 ヘッダ データ に変換された結果の JSON の例です。  
以下の項目のうち、"OrderID" ～ "IsCancelled" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。  

```
{
  "connection_key": "",
  "result": false,
  "redis_key": "",
  "filepath": "",
  "api_status_code": 0,
  "runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
  "business_partner": null,
  "service_label": "SAPProductionOrderConfCreates",
  "api_type": "",
  "ProductionOrderConfirmation": {
    "OrderID": "1000001",
    "Sequence": "0",
    "OrderType": "PP01",
    "ConfirmationText": null,
    "ConfirmationYieldQuantity": null,
    "ConfirmationScrapQuantity": null,
    "ConfirmationGroup": "1",
    "Language": "",
    "Material": "5",
    "Plant": "0001",
    "WorkCenter": ""
  },
  "api_schema": "SAPProductionOrderConfCreates",
  "accepter": null,
  "deleted": false,
  "api_processing_result": true,
  "api_processing_error": ""
}

```
