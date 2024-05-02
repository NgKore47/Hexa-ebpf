package config

import (
	// "log"
	"github.com/kcl17/logger"
)

var Conf UpfConfig

// Init init config for eupf package
func Init() {
	if err := Conf.Unmarshal(); err != nil {
		logger.AppLog.Infof("Config verification failed ")
	}
    
	if err := Conf.Validate(); err != nil {
		// logger.InitLog.Infof("hwllo eorld")
	}
	logger.AppLog.Infof("+-------------------+--------------------+")
	logger.AppLog.Infof("|%-19s|%-20s|","Attach Mode", Conf.XDPAttachMode)
	logger.AppLog.Infof("|%-19s|%-18v|","Interface List", Conf.InterfaceName)
	// for _, inface := range Conf.InterfaceName {
	// }
	logger.AppLog.Infof("+-------------------+--------------------+")
	logger.AppLog.Infof("|%-19s|%-20s|","Api address", Conf.ApiAddress)
	logger.AppLog.Infof("|%-19s|%-20s|", "pfcp address",Conf.PfcpAddress)
	logger.AppLog.Infof("|%-19s|%-20s|", "node ID", Conf.PfcpNodeId)
	logger.AppLog.Infof("|%-19s|%-20s|", "N3 address",Conf.N3Address)
	logger.AppLog.Infof("|%-19s|%-20d|", "echo interval",Conf.EchoInterval)
	logger.AppLog.Infof("|%-19s|%-20d|", "qermap size",Conf.QerMapSize)
	logger.AppLog.Infof("|%-19s|%-20d|", "farmap size",Conf.FarMapSize)
	logger.AppLog.Infof("|%-19s|%-20d|", "pdrmap size",Conf.PdrMapSize)
	logger.AppLog.Infof("|%-19s|%-20d|", "heartbeatinterval",Conf.HeartbeatInterval)
	logger.AppLog.Infof("+-------------------+--------------------+")

	// logger.AppLog.Infof("|%-19s|%-20s|", "",Conf.UEIPPool)
	// logger.AppLog.Infof("|%-19s|%-20s|", Conf.FTEIDPool)
	// logger.AppLog.Infof("attach mode: %s", Conf.FTEIDPool)
	// log.Printf("Apply eUPF config: %+v", Conf)
}
