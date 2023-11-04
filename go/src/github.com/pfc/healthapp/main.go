package main

import (
	"fmt"
	"github.com/pfc/healthapp/blockchain"
	"github.com/pfc/healthapp/web"
	"github.com/pfc/healthapp/web/controllers"
	"os"
)

func main() {
	// Definição dos parametros do SDK
	fSetup := blockchain.FabricSetup{
		// parametro da rede
		OrdererID: "orderer.app.com",

		// parametros do canal
		ChannelID:     "healthapp",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/pfc/healthapp/artifacts/healthapp.channel.tx",

		// parametros do chaincode
		ChainCodeID:     "health-app",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/pfc/healthapp/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// parametros do usuario
		UserName: "User1",
	}

	// Inicialização do SDK
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}

	// Fecha o SDK
	defer fSetup.CloseSDK()

	// Instalação e instanciação do Chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}
	
	
	//	Inicia o servidor web
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)

}

