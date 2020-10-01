package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	paramMap := make(map[string]interface{})

	//APP请求
	url := "https://test-t3.ipalfish.com/klian/wechat/wechatcourse/openrank/get"
	paramMap["courseid"] = 258602326571038
	paramMap["uid"] = 2864473666
	paramMap["shareurl"] = "http://test-t3.ipalfish.com/phonics/classroom/learn/readshare.html?courseid=258602326571038&uid=2864473666&test=1"

	//opapi请求
	//url := "https://test.ipalfish.com:30000/opapi/wechat/wechatcourse/yesterday/unfinsh/list"
	//paramMap["courseid"] = 258014983471130
	//paramMap["classid"] = 362666112788480

	//wechatcourse请求
	//url := "https://test.ipalfish.com:30000/wechatcourseopapi/admin/wechat/wechatcourse/debug/lesson/send"
	//paramMap["courseId"] = 279619381258242
	//paramMap["lessonId"] = 362666112788481
	//paramMap["completeType"] = 1
	//paramMap["completeTime"] = "2020-06-29"
	//paramMap["uid"] = 10427975
	//paramMap["uid"] = 7609390
	//SendRequest(url,paramMap, "t3")

	// ==== CRM请求
	//获取用户信息
	//url := "https://sea-test.pri.ibanyu.com/wechatcourse/ops/user/details/get"
	//paramMap["uid"] = 24444
	//SendRequest(url,paramMap)

	//查询员工信息
	//url := "https://sea-test.pri.ibanyu.com/wechatcourse/ops/employee/list/get"
	//paramMap["employee_name"] = "zhao"
	SendRequest(url, paramMap, "t3")

	//Post(url,paramMap,"application/json")
}

func SendRequest(url string, paramMap map[string]interface{}, routerGroup string) {
	jsonStr, _ := json.Marshal(paramMap)
	payload := bytes.NewBuffer(jsonStr)
	req, _ := http.NewRequest("POST", url, payload)
	if len(routerGroup) > 0 {
		req.Header.Add("ipalfish-group", routerGroup)
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Accept", "*/*")
	//req.Header.Add("Cookie", "ipalfish_device_id=65c03b80613fa7b5c1dd31a096ee72cb; utype=op; user=zhaojinxin10242; id=MjYwMzE=; name=6LW16YeR6ZGr; email=; phone=ODYtMTc2NDUwMTM4OTg=; groups=WyIxMzYxOTU0IiwiZHV3byIsInJkIl0=; token=395de2900d737ee90c30685c3b0054e2; logintype=0")
	//req.Header.Add("Cookie", "ipalfish_device_id=8ecaf80d4555c19103f8d7bd8b4b5731; utype=op; user=zhaojinxin10242; id=MjYwMzE=; name=6LW16YeR6ZGr; email=; phone=ODYtMTc2NDUwMTM4OTg=; groups=WyIxMzYxOTU0IiwicmQiLCJkdXdvIl0=; token=53645fe2ad7f5bf429371e6491c443d3; logintype=0")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
