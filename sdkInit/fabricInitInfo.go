package sdkInit

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

type InitInfo struct {
	ChannelID     string //通道名称
	ChannelConfig string //通道交易配置文件路径
	OrgAdmin      string//组织管理员名称
	OrgName       string//组织名称
	OrdererOrgName	string//Orderer名称
	OrgResMgmt *resmgmt.Client//资源管理端实例

	ChaincodeID	string
	ChaincodeGoPath	string
	ChaincodePath	string
	UserName	string
}
