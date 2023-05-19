> json序列化和反序列化,时间类型和时间戳互转

自定义时间类型,并实现UnmarshalJSON和MarshalJSON方法,用于处理该类型JSON序列化与反序列化
```go
package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type UnixMilliTime time.Time

func (t *UnixMilliTime) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = UnixMilliTime(time.UnixMilli(i))
	return nil
}

func (t *UnixMilliTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Time(*t).UnixMilli())
	return []byte(stamp), nil
}
func (t *UnixMilliTime) ToTime() time.Time {
	return time.Time(*t)
}

func main() {
	type Params struct {
		TestTime UnixMilliTime `json:"testTime"`
	}
	testJSON := `{"testTime":1683693520064}`
	req := new(Params)

	err := json.Unmarshal([]byte(testJSON), &req)
	if err != nil {
		panic(err)
	}
	//
	fmt.Println("前端传入Body", testJSON)
	fmt.Println("反序列化后", req)
	fmt.Println("反序列化后,TestTimen 转原生时间类型", req.TestTime.ToTime())

	resTestTime := time.Now()
	res := Params{
		UnixMilliTime(resTestTime),
	}
	t2json, err := json.Marshal(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println("返回的结构体 res = ", res)
	fmt.Println("返回的结构体 res.TestTimen 原生时间类型", res.TestTime.ToTime())
	fmt.Println("res序列化后返回给前端的body", string(t2json))
}

```