# data-interface-for-salesforce-price-type-list
data-interface-for-salesforce-price-type-list は、salesforce の価格表オブジェクトに対する各種アクションに必要なデータの整形、および salesforce から受け取った response の MySQL への格納を行うマイクロサービスです。

## 動作環境  
data-interface-for-salesforce-price-type-listは、aion-coreのプラットフォーム上での動作を前提としています。  
使用する際は、事前に下記の通りAIONの動作環境を用意してください。     
  
* OS: Linux OS     
* CPU: ARM/AMD/Intel    
* Kubernetes     
* [AION](https://github.com/latonaio/aion-core)のリソース      

## セットアップ
1. 以下のコマンドを実行して、docker imageを作成してください。
```
$ cd /path/to/data-interface-for-salesforce-price-type-list
$ make docker-build
```

## 起動方法
以下のコマンドを実行して、podを立ち上げてください。
```
$ cd /path/to/data-interface-for-salesforce-price-type-list
$ kubectl apply -f data-interface-for-salesforce-customer.yaml
```

## kanban との通信
### kanban から受信するデータ
kanban から受信する metadata に下記の情報を含む必要があります。

| key | value |
| --- | --- |
| method | get |
| object | Customer |
| id | 顧客 ID |
| account_id | 顧客 ID |

具体例 1: 
```example
# metadata (map[string]interface{}) の中身

"method": "get"
"object": "Account"
"id": "xxxx"
```

具体例 2: 
```example
# metadata (map[string]interface{}) の中身

"method": "get"
"object": "Account"
"account_id": "xxxx"
```

### kanban に送信するデータ
kanban に送信する metadata は下記の情報を含みます。

| key | type | description |
| --- | --- | --- |
| method | string | 文字列 "get" を指定 |
| object | string | 文字列 "Account" を指定 |
| path_param | string | 顧客 ID を指定 |
| connection_key | string | 文字列 "customer" を指定 |

具体例: 
```example
# metadata (map[string]interface{}) の中身

"method": "get"
"object": "Account"
"path_param": "xxxx"
"connection_key": "customer"
```

## kanban(salesforce-api-kube) から受信するデータ
kanban からの受信可能データは下記の形式です

| key | value |
| --- | --- |
| key | 文字列 "Account" |
| content | Contract の詳細情報を含む JSON 配列 |
| connection_type | 文字列 "response" |

具体例:
```example
# metadata (map[string]interface{}) の中身

"key": "Account"
"content": "[{xxxxxxxxxxx}]"
"connection_type": "response"
```