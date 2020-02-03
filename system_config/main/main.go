package main

import (
	"encoding/json"
	"fmt"
	"go_system/system_net"
)

func main() {
	option := system_net.NetOptions{}
	option.InsecureSkipVerify = true
	option.Cookies = map[string]string{"ff": "ck_god", "xx": "god_girl"}
	res, err := system_net.Get("http://106.54.93.177:8081/v1/user/getall", option)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	type item struct {
		Auto_id    int
		FriendlyId string `json:"friendly_id"`
		Id         string
		Status     int
		Type       int
		UserId     int    `json:"user_id"`
		PoolTypeId string `json:"pool_type_id"`
		Name       string
	}
	type commonRes struct {
		Code uint16
		Data [] item
	}
	js := res
	//js := `{"code":201,"data":[{"auto_id":1,"id":"011ca600-88ab-11e8-beba-3129d7a30c16","status":1,"type":2,"user_id":200173,"pool_type_id":"322e73b0-8408-11e8-965f-ef46ea4ca8e7","friendly_id":"C201807160001","name":"C360","class":4,"abbreviation":"C","miner_count":100,"miner_amount":100000,"amount_coin":2,"miner_coin":100000,"basic_yield":33,"presented_yield":87,"annual_yield":43.2,"sold":13,"expires":1562860800,"duration":31104000,"mining_date":"2018-10-12","mining_left":1,"mining_count":1,"created_at":"2018-7-16","updated_at":"2018-10-11"}]}`
	var xm commonRes
	err = json.Unmarshal([]byte(js), &xm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(xm) //输出{xiaoming 18 {Hunan ChangSha}}
}
