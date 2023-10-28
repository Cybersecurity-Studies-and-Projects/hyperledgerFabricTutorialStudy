# Tutorial Hyperledger Fabric
## Pré-requisitos
1. Conhecimento prévio sobre a utilização prática da plataforma Hyperledger Fabric e seus componentes
    * Base: documentação oficial da Hyperledger Fabric.
2. Sistema operacional Linux Ubuntu, versão 20.04 LTS.
3. Git instalado.
4. Ambiente de desenvolvimento da linguagem de programação GO - versão 1.15.2 (Workspace configurado na pasta "~/" do Ubuntu).
    * Versão 64 bits - Download
    ```shell 
    $ cd ~

    $ curl -O https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz

    $ sudo tar -xvf go1.12.1.linux-amd64.tar.gz -C /usr/local

    $ sudo chown -R root:root /usr/local/go

    $ mkdir -p $HOME/go/{bin,src}

    ```
    * Definir variáveis de ambiente:
    ```shell
    Abrir arquivo para edição:
    $ nano ~/.profile 

    colar no arquivo:
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin

    Carregar as variáveis globais: 
    $ . ~/.profile
    ```

    go version
    go env
    ```
5. Docker versão 19.03.12
    ```shell
    $ wget https://download.docker.com/linux/static/stable/x86_64/docker-19.03.12.tgz

    $ tar xzvf docker-[19.03.12.tgz

    $ sudo mv docker/* /usr/bin/

    ```
6. Docker Compose versão 1.14.0.

    ```shell
    # Baixe o Docker Compose no diretório /usr/local/bin
    $ sudo curl -L "https://github.com/docker/compose/releases/download/1.14.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

    # Aplique permissões de execução ao binário
    $ sudo chmod +x /usr/local/bin/docker-compose

    # Verifique a instalação
    $ docker-compose --version
    ```
## Clonar o repositório
    * No workspace Go: ~/go
    1. Entrar na pasta src:
      ```shell
      $ cd ~/go/src
      ```

    2. Crie a pasta "github.com" e entre nela:
    ```shell
      $ mkdir github.com

      cd github.com
    ```
    
    3. Crie a pasta "pfc" e entre nela:
    ```shell
      $ mkdir pfc

      cd pfc
    ```

    4. Clone o  repositório:
      ```shell
      $ git clone https://github.com/coutinhotauan/healthapp
      ```
## Editar o arquivo "docker-compose.yaml"
* Comentar o comando "build: .":
```yaml
version: '3'

networks:
  default:

services:

  orderer:
    image: hyperledger/fabric-orderer:x86_64-1.1.0
    # build: .
    container_name: orderer
    environment:
      - ORDERER_GENERAL_LOGLEVEL=debug
      - ORDERER_GENERAL_LISTENADDRESS=0.0.0.0
      - ORDERER_GENERAL_LISTENPORT=7050
      - ORDERER_GENERAL_GENESISPROFILE=Healthapp
      - ORDERER_GENERAL_GENESISMETHOD=file
      - ORDERER_GENERAL_GENESISFILE=/var/hyperledger/orderer/orderer.genesis.block
      - ORDERER_GENERAL_LOCALMSPID=orderer.app.com
      - ORDERER_GENERAL_LOCALMSPDIR=/var/hyperledger/orderer/msp
      - ORDERER_GENERAL_TLS_ENABLED=false
      - ORDERER_GENERAL_TLS_PRIVATEKEY=/var/hyperledger/orderer/tls/server.key
      - ORDERER_GENERAL_TLS_CERTIFICATE=/var/hyperledger/orderer/tls/server.crt
      - ORDERER_GENERAL_TLS_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]

    working_dir: /opt/gopath/src/github.com/hyperledger/fabric
    command: orderer
    volumes:
      - ./artifacts/orderer.genesis.block:/var/hyperledger/orderer/orderer.genesis.block
      - ./crypto-config/ordererOrganizations/orderer.app.com/orderers/orderer.orderer.app.com/msp:/var/hyperledger/orderer/msp
      - ./crypto-config/ordererOrganizations/orderer.app.com/orderers/orderer.orderer.app.com/tls:/var/hyperledger/orderer/tls
    ports:
      - 7050:7050
    networks:
      default:
        aliases:
          - orderer.app.com

  ca:
    image: hyperledger/fabric-ca:x86_64-1.1.0
    # build: .
    container_name: ca
    environment:
      - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
      - FABRIC_CA_SERVER_CA_NAME=ca.org1.app.com
      - FABRIC_CA_SERVER_CA_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.org1.app.com-cert.pem
      - FABRIC_CA_SERVER_CA_KEYFILE=/etc/hyperledger/fabric-ca-server-config/d69797e58d6f371c44ce0ca935aab5666790700a4cff30ce47078d93946c6199_sk
      - FABRIC_CA_SERVER_TLS_ENABLED=false
      - FABRIC_CA_SERVER_TLS_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.org1.app.com-cert.pem
      - FABRIC_CA_SERVER_TLS_KEYFILE=/etc/hyperledger/fabric-ca-server-config/d69797e58d6f371c44ce0ca935aab5666790700a4cff30ce47078d93946c6199_sk
    ports:
      - 7054:7054
    command: sh -c 'fabric-ca-server start -b admin:adminpw -d'
    volumes:
      - ./crypto-config/peerOrganizations/org1.app.com/ca/:/etc/hyperledger/fabric-ca-server-config
    networks:
      default:
        aliases:
          - ca.org1.app.com

  peer0_org1:
    image: hyperledger/fabric-peer:x86_64-1.1.0
    # build: .
    container_name: peer0_org1
    environment:
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      - CORE_VM_DOCKER_ATTACHSTDOUT=true
      - CORE_LOGGING_LEVEL=DEBUG
      - CORE_PEER_NETWORKID=healthapp
      - CORE_PEER_PROFILE_ENABLED=true
      - CORE_PEER_TLS_ENABLED=false
      - CORE_PEER_TLS_CERT_FILE=/var/hyperledger/tls/server.crt
      - CORE_PEER_TLS_KEY_FILE=/var/hyperledger/tls/server.key
      - CORE_PEER_TLS_ROOTCERT_FILE=/var/hyperledger/tls/ca.crt
      - CORE_PEER_ID=peer0.org1.app.com
      - CORE_PEER_ADDRESSAUTODETECT=true
      - CORE_PEER_ADDRESS=peer0.org1.app.com:7051
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.org1.app.com:7051
      - CORE_PEER_GOSSIP_USELEADERELECTION=true
      - CORE_PEER_GOSSIP_ORGLEADER=false
      - CORE_PEER_GOSSIP_SKIPHANDSHAKE=true
      - CORE_PEER_LOCALMSPID=org1.app.com
      - CORE_PEER_MSPCONFIGPATH=/var/hyperledger/msp
      - CORE_PEER_TLS_SERVERHOSTOVERRIDE=peer0.org1.app.com
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: peer node start
    volumes:
      - /var/run/:/host/var/run/
      - ./crypto-config/peerOrganizations/org1.app.com/peers/peer0.org1.app.com/msp:/var/hyperledger/msp
      - ./crypto-config/peerOrganizations/org1.app.com/peers/peer0.org1.app.com/tls:/var/hyperledger/tls
    ports:
      - 7051:7051
      - 7053:7053
    depends_on:
      - orderer
    links:
      - orderer
    networks:
      default:
        aliases:
          - peer0.org1.app.com

  peer1_org1:
    image: hyperledger/fabric-peer:x86_64-1.1.0
    # build: .
    container_name: peer1_org1
    environment:
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      - CORE_VM_DOCKER_ATTACHSTDOUT=true
      - CORE_LOGGING_LEVEL=DEBUG
      - CORE_PEER_NETWORKID=healthapp
      - CORE_PEER_PROFILE_ENABLED=true
      - CORE_PEER_TLS_ENABLED=false
      - CORE_PEER_TLS_CERT_FILE=/var/hyperledger/tls/server.crt
      - CORE_PEER_TLS_KEY_FILE=/var/hyperledger/tls/server.key
      - CORE_PEER_TLS_ROOTCERT_FILE=/var/hyperledger/tls/ca.crt
      - CORE_PEER_ID=peer1.org1.app.com
      - CORE_PEER_ADDRESSAUTODETECT=true
      - CORE_PEER_ADDRESS=peer1.org1.app.com:7051
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer1.org1.app.com:7051
      - CORE_PEER_GOSSIP_USELEADERELECTION=true
      - CORE_PEER_GOSSIP_ORGLEADER=false
      - CORE_PEER_GOSSIP_SKIPHANDSHAKE=true
      - CORE_PEER_LOCALMSPID=org1.app.com
      - CORE_PEER_MSPCONFIGPATH=/var/hyperledger/msp
      - CORE_PEER_TLS_SERVERHOSTOVERRIDE=peer1.org1.app.com
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: peer node start
    volumes:
      - /var/run/:/host/var/run/
      - ./crypto-config/peerOrganizations/org1.app.com/peers/peer1.org1.app.com/msp:/var/hyperledger/msp
      - ./crypto-config/peerOrganizations/org1.app.com/peers/peer1.org1.app.com/tls:/var/hyperledger/tls
    ports:
      - 8051:7051
      - 8053:7053
    depends_on:
      - orderer
    links:
      - orderer
    networks:
      default:
        aliases:
          - peer1.org1.app.com
```
## Iniciar a aplicação
A inicialização da aplicação parte da execução do arquivo Go referente aos mó-
dulos client e server para a utilização da interface de usuário, e também da inicialização da rede de blockchain a ser utilizada.

Esta inicialização apresenta diferentes passos a serem seguidos, conforme o
modelo a ser utilizado para instanciação da rede desenvolvida na plataforma Hyperledger
Fabric.

### Inicializar o modelo Único Host
1. Dentro do repositório clonado:
    ```shell
    cd ~/go/src/github.com/pfc/healthapp
    ```
2. Execute:
    ```shell
    $ go build
    ```
  * Este comando irá realizar o download de todas as dependências necessárias (especificadas no arquivo go.mod) e então irá compilar o arquivo denominado healthapp.

Para continuar a execução iremos utilizar dois terminais (Terminal 1 e Terminal 2)
para acompanhar a aplicação em detalhes.

3. Terminal 1 - execute:
    ```shell
    $ sudo docker-compose up
    ```

4. Você deverá ver o log funcionando, com o final do terminal parecido com:
```shell
...
peer0_org1    | 2023-10-27 16:05:11.416 UTC [flogging] setModuleLevel -> DEBU 1b6 Module 'peer/gossip/mcs' logger enabled for log level 'WARNING'
peer0_org1    | 2023-10-27 16:05:11.417 UTC [nodeCmd] func7 -> INFO 1b7 Starting profiling server with listenAddress = 0.0.0.0:6060

```

Este comando irá criar os devidos containers e inseri-los na rede criada pelo Docker
Compose. A partir do momento da criação, será possível observar os logs gerados
pelos containers. Este terminal será utilizado para a visualização dos logs em tempo
real.

5. No Terminal 2, execute o seguinte comando:
./healthapp
Este comando irá executar o arquivo Go compilado. De acordo com a execução, será
possível acompanhar os processos necessários para efetivação da rede Hyperledger
Fabric através da aplicação. Por fim, será especificado no terminal o endereço em
que a interface de usuário estará disponível.
Acesse a aplicação pelo navegador através do endereço especificado e explore suas
funcionalidades através de sua interface.


ps axf | grep docker | grep -v grep | awk ' {print "kill -9 " $1}' | sudo sh
<hr/>
<hr/>
<hr/>
<hr/>
<hr/>
<hr/>
<hr/>
<hr/>
<hr/>
<hr/>