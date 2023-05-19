package main

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

type JsonDecimal decimal.Decimal // 自定义Decimal类型

// UnmarshalJSON 用于自定义类型反序列化
func (j *JsonDecimal) UnmarshalJSON(data []byte) error {
	dataS := string(data)
	dataS = strings.TrimPrefix(dataS, `"`) //传入的字符串 前后会有 \" \" 需要去掉
	dataS = strings.TrimSuffix(dataS, `"`)
	if dataS == "" { //字符串的零值 为 "" 需要转0
		dataS = "0"
	}
	decimalData, err := decimal.NewFromString(dataS)
	if err != nil {
		return err
	}
	*j = JsonDecimal(decimalData)
	return nil
}

// MarshalJSON 用于自定义类型序列化
func (j *JsonDecimal) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, j.ToDecimal().String())), nil
}

// ToDecimal 转原生Decimal
func (j *JsonDecimal) ToDecimal() decimal.Decimal {
	return decimal.Decimal(*j)
}

func main() {
	type Params struct {
		Amount JsonDecimal `json:"amount"`
	}
	testJSON := `{"amount":"157.45"}`
	req := new(Params)

	err := json.Unmarshal([]byte(testJSON), &req)
	if err != nil {
		panic(err)
	}
	//
	fmt.Println("前端传Body", testJSON)
	fmt.Println("反序列化后", req)
	fmt.Println("反序列化后,AmountDecimal 转decimal.Decimal", req.Amount.ToDecimal())

	resAmount := decimal.NewFromFloat(1.4)
	res := Params{
		JsonDecimal(resAmount),
	}
	resJson, err := json.Marshal(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println("返回的结构体 res = ", res)
	fmt.Println("返回的结构体 AmountDecimal 转decimal.Decimal", res.Amount.ToDecimal())
	fmt.Println("res序列化后返回给前端的body", string(resJson))
}
