
### prepare
- sudo apt install aufs-tools
- change max open files

### docker
- sudo service docker stop
- sudo docker daemon --api-cors-header="*" -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock


### derict run hyperledger/fabric

- peer  
peer node start --peer-chaincodedev

### chaincode
- CORE_CHAINCODE_ID_NAME=mycc CORE_PEER_ADDRESS=0.0.0.0:7051 ./chaincode_example02



### start
sudo docker-compose -f docker-compose-with-membersrvc.yml up

### 用户登陆
- POST		HOST:7050/registrar

    {
	   	"enrollId":	"jim",
	    "enrollSecret":	"6avZQLwcUe9b"
    }


### chaincode	部署
peer	chaincode	deploy	-p	github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02	-c

- POST		HOST:7050/chaincode

    {
		"jsonrpc":	"2.0",
		"method":	"deploy",
		"params":	{
				"type":	1,
				"chaincodeID":{
								"path":"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02"
				},
				"ctorMsg":	{
								"function":"init",
								"args":["a",	"1000",	"b",	"2000"]
				},
				"secureContext":	"jim"
		},
		"id":	1
    }

- Response

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "04233c6dd8364b9f0749882eb6d1b50992b942aa0a664182946f411ab46802a88574932ccd75f8c75e780036e363d52dd56ccadc2bfde95709fc39148d76f050"
      },
      "id": 1
    } 

### chaincode	调用  
- POST		HOST:7050/chaincode 

    {
		"jsonrpc":	"2.0",
		"method":	"invoke",
		"params":	{
						"type":	1,
						"chaincodeID":{
										"name":"04233c6dd8364b9f0749882eb6d1b50992b942aa0a664182946f411ab46802a88574932ccd75f8c75e780036e363d52dd56ccadc2bfde95709fc39148d76f050"
						},
						"ctorMsg":	{
									"function":"invoke",
									"args":["a",	"b",	"100"]
						},
				"secureContext":	"jim"
		},
		"id":	3
    }

- Response

    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "36691956-8b8a-4c29-b891-9adc35a14c1b"
      },
      "id": 3
    }    

### chaincode	查询
- POST		HOST:7050/chaincode

{
		"jsonrpc":	"2.0",
		"method":	"query",
		"params":	{
						"type":	1,
						"chaincodeID":{
										"name":"04233c6dd8364b9f0749882eb6d1b50992b942aa0a664182946f411ab46802a88574932ccd75f8c75e780036e363d52dd56ccadc2bfde95709fc39148d76f050"
						},
						"ctorMsg":	{
									"function":"query",
									"args":["a"]
						},
						"secureContext":	"jim"
		},
		"id":	5
}

### 区块信息查询
- GET		HOST:7050/chain/blocks/2

