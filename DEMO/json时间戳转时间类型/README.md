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

type UnixMilliTime time.Time // 自定义时间类型

// UnmarshalJSON 用于自定义类型反序列化
func (t *UnixMilliTime) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(string(data), 10, 64) // 将传入的data 转数字
	if err != nil {
		return err
	}
	*t = UnixMilliTime(time.UnixMilli(i)) //传入的时间戳转时间类型
	return nil
}

// MarshalJSON 用于自定义类型序列化
func (t *UnixMilliTime) MarshalJSON() ([]byte, error) {
	if t.ToTime().IsZero() { //如果是零值 返回 0
		return []byte{'0'}, nil
	}
	stamp := fmt.Sprintf("%d", t.ToTime().UnixMilli()) //转毫秒级时间戳
	return []byte(stamp), nil
}

// ToTime 自定义类型转原生时间类型
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
	resJson, err := json.Marshal(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println("返回的结构体 res = ", res)
	fmt.Println("返回的结构体 res.TestTimen 原生时间类型", res.TestTime.ToTime())
	fmt.Println("res序列化后返回给前端的body", string(resJson))
}

```