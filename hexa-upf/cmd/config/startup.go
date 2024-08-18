package config

import (
	// "log"
	"github.com/kcl17/logger"
)

var Conf UpfConfig

// Init init config for eupf package
func Init() {
	if err := Conf.Unmarshal(); err != nil {
		logger.AppLog.Infof("hello world")
	}
    
	if err := Conf.Validate(); err != nil {
		logger.InitLog.Infof("hwllo eorld")
	}
	logger.AppLog.Infof("+-------------------+------------------+")
	logger.AppLog.Infof("|%-19s|%-18s|","Attach Mode", Conf.XDPAttachMode)
	logger.AppLog.Infof("|%-19s|%-v       |","Interface List", Conf.InterfaceName)
	// for _, inface := range Conf.InterfaceName {
	// }
	logger.AppLog.Infof("+-------------------+------------------+")
	logger.AppLog.Infof("|%-19s|%-18s|","Api address", Conf.ApiAddress)
	logger.AppLog.Infof("|%-19s|%-18s|", "pfcp address",Conf.PfcpAddress)
	logger.AppLog.Infof("|%-19s|%-18s|", "node ID", Conf.PfcpNodeId)
	logger.AppLog.Infof("|%-19s|%-18s|", "N3 address",Conf.N3Address)
	logger.AppLog.Infof("|%-19s|%-18d|", "echo interval",Conf.EchoInterval)
	logger.AppLog.Infof("|%-19s|%-18d|", "qermap size",Conf.QerMapSize)
	logger.AppLog.Infof("|%-19s|%-18d|", "farmap size",Conf.FarMapSize)
	logger.AppLog.Infof("|%-19s|%-18d|", "pdrmap size",Conf.PdrMapSize)
	logger.AppLog.Infof("|%-19s|%-18d|", "heartbeatinterval",Conf.HeartbeatInterval)
	logger.AppLog.Infof("+-------------------+------------------+")

	// logger.AppLog.Infof("|%-19s|%-18s|", "",Conf.UEIPPool)
	// logger.AppLog.Infof("|%-19s|%-18s|", Conf.FTEIDPool)
	// logger.AppLog.Infof("attach mode: %s", Conf.FTEIDPool)
	// log.Printf("Apply eUPF config: %+v", Conf)
}
