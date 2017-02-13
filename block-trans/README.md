
### 发布sample02

### 循环post1000次

    func main() {
	    sum := 0
    	for i := 0; i < 1000; i++ {
    		sum += i
    		response, _ := PostCC()
    		fmt.Println(i)
    		fmt.Println(response)
    	}
    
    }

    func PostCC() (response string, err error) {
    	finalURL := "http://localhost:7050/chaincode"
    	_json := `{
    		"jsonrpc":	"2.0",
    		"method":	"invoke",
	    	"params":	{
	    					"type":	1,
	    					"chaincodeID":{
	    									"name":"f889cf53dae89ecb8dadbba6d5f578fd0de88b55a8f3bba162d878323b520f16e68c6fe03b5bd9ed7ff8d98e96f030eb1764def3e4fda865c9e8317200508858"
    						},
    						"ctorMsg":	{
    									"function":"invoke",
    									"args":["a",	"b",	"-1"]
    						},
    				"secureContext":	"jim"
    		},
    		"id":	3
    }
    	`
    
    	return http.PostJSONString(finalURL, _json)
    }

### 现象
Fabric会把多个post信息,汇集成一个block里的多个trans

参见目录中的png图

