package main

import "github.com/vmihailenco/msgpack/v5"

func main() {
	var requestMap = make(map[string]interface{})
	requestMap["query"] = "1234"

	var responseMap = make(map[string]interface{})
	//responseMap = requestMap
	CopyMap(requestMap, &responseMap)

	responseMap["query"] = "3456"

	fmt.Printf("request:\n%+v\n\n", requestMap)
	fmt.Printf("response:\n%+v\n\n", responseMap)


}

func CopyMap(in, out interface{}) {
	b, err := msgpack.Marshal(in)
	if err != nil {
		fmt.Println(err)
	}

	err = msgpack.Unmarshal(b, &out)
	if err != nil {
		fmt.Println(err)
	}
}
