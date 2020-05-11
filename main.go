package main

import (
	"fmt"
	"github.com/orangey.com/education/sdkInit"
	"github.com/orangey.com/education/service"
	"github.com/orangey.com/education/web"
	"github.com/orangey.com/education/web/controller"
	"os"
)

const (
	configFile = "config.yaml"
	initialized = false
	EduCC = "educc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID: "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/orangey.com/education/fixtures/artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID: EduCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "github.com/orangey.com/education/chaincode/",
		UserName:"User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	//程序退出前关闭释放由SDK维护的缓存和连接
	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID:EduCC,
		Client:channelClient,
	}

	edu := service.Education{
		Name: "黄成",
		Gender: "男",
		Nation: "汉",
		EntityID: "412829199705232011",
		Place: "北京",
		BirthDay: "1994-01-01",
		EnrollDate: "2013-09-01",
		GraduationDate: "2016-07-01",
		SchoolName: "北京石油化工学院",
		Major: "计算机科学与技术",
		QuaType: "普招",
		Length: "四年",
		Mode: "全日制",
		Level: "本科",
		Graduation: "毕业",
		CertNo: "111",
		Photo: "/static/photo/黄成.jpg",
	}

	edu2 := service.Education{
		Name: "赵思",
		Gender: "女",
		Nation: "汉族",
		EntityID: "142829199705232018",
		Place: "上海",
		BirthDay: "1994-02-01",
		EnrollDate: "2013-09-01",
		GraduationDate: "2016-07-01",
		SchoolName: "北京石油化工学院",
		Major: "信息与计算科学",
		QuaType: "普招",
		Length: "四年",
		Mode: "全日制",
		Level: "本科",
		Graduation: "毕业",
		CertNo: "222",
		Photo: "/static/photo/赵思.jpg",
	}

	msg, err := serviceSetup.SaveEdu(edu)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	msg, err = serviceSetup.SaveEdu(edu2)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}


	//修改/添加信息
	info := service.Education{
		Name: "黄成",
		Gender: "男",
		Nation: "汉族",
		EntityID: "412829199705232011",
		Place: "北京",
		BirthDay: "1994-01-01",
		EnrollDate: "2016-09-01",
		GraduationDate: "2018-07-01",
		SchoolName: "北京大学",
		Major: "计算机科学与技术",
		QuaType: "普招",
		Length: "两年",
		Mode: "全日制",
		Level: "研究生",
		Graduation: "毕业",
		CertNo: "333",
		Photo: "/static/photo/黄成.jpg",
	}
	msg, err = serviceSetup.ModifyEdu(info)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息操作成功, 交易编号为: " + msg)
	}

	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)

}
