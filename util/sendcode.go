package util

import (
	"encoding/json"
	"fmt"
	"github.com/SSunSShine/QAsystem/database"
	"log"
	"math/rand"
	"time"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

)

func  SendCode(phone string) (ok bool) {
	//1.产生验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	//code :="5201314"
	//2.调用阿里云sdk 完成发送

	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4G9NPb36CHogBuhRzpah", "Fno40bapVuwpdHrJiS6eBX2BSUMdhl")
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = "QA问答"
	request.TemplateCode = "SMS_206535804"
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err1 := client.SendSms(request)
	if err1 != nil {
		log.Fatal(err1.Error())
	}
	fmt.Printf("response is %#v\n", response)

	//3.接受返回结果，判断发送状体
	if response.Code == "OK" {
		//讲验证码保存到数据库
		//smscode := model.SmsCode{Phone: phone, Code: code, BizId: response.BizId, CreatTime: time.Now().Unix()}
		//memberdao := dao.MemberDao{tool.DbEngine}
		//result := memberdao.InsertCode(smscode)
		CodeSaveRedis(phone,code)

		return true
	}
	return false
}

func CodeSaveRedis(phone,code string){
	//fmt.Println("setcode_begin")
	//defer fmt.Println("setcode_end")

	//kv读写
	err :=database.RDB.Set(phone, code, 5*time.Minute).Err()
	if err!=nil {
		log.Fatal(err.Error())
	}
	//获取过期时间
	//tm, err := RDB.TTL(phone).Result()
	//log.Println(tm)

	//获取值
	//val, err := database.RDB.Get(phone).Result()
	//log.Println(val)
}
