# Tutorial Hyperledger Fabric

Este repositório contém tutoriais sobre o [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/release-2.5/tutorials.html). Os tutoriais do Fabric podem ser usados por desenvolvedores para começar a criar suas próprias soluções.

## Etapas

1. **Trabalhar com o Fabric implantando a rede de teste em uma máquina local.**
    - Usar os passos fornecidos pelo tutorial de implantação de um contrato inteligente em um canal.

2. **Executando uma Aplicação Fabric**
    - Fornece uma introdução sobre como usar as APIs fornecidas pelas SDKs do Fabric para invocar contratos inteligentes a partir de aplicações cliente.

3. **Implantando um contrato inteligente em um canal e Criando um canal**
    - Podem ser usados para aprender aspectos importantes da administração de uma rede em execução.

4. **Dados privados e CouchDB**
    - Explorar recursos importantes do Fabric.

5. **Implantação de uma rede de produção**
    - Implantar o Hyperledger Fabric em produção.

6. **Atualizar um canal**
    - Atualização da configuração de um canal e Atualização do nível de capacidade de um canal.

7. **Atualizar componentes como nós pares, nós de ordenação, SDKs e outros**
    - Atualização de seus componentes.

8. **Introdução sobre como escrever um contrato inteligente básico**
    - Escrevendo seu Primeiro Chaincode.

## Using the Fabric test network (Usando a rede de teste do Fabric)

Após baixar as imagens e os exemplos do Docker do Hyperledger Fabric, é possível implantar uma rede de teste usando scripts fornecidos no repositório. A rede de teste é destinada ao aprendizado do Fabric, executando nós em sua máquina local. Os desenvolvedores podem usar a rede para testar seus contratos inteligentes e aplicativos. A rede destina-se a ser usada apenas como uma ferramenta para educação e teste e não como um modelo para configurar uma rede de produção. Modificações nos scripts são desencorajadas e podem interromper a rede.

### Características

- Inclui duas organizações de pares e uma organização de ordenação.
- Para simplificação, é configurado um serviço de ordenação Raft de um único nó.
- Para reduzir a complexidade, uma Autoridade de Certificação (CA) TLS não é implantada. Todos os certificados são emitidos pelas ACs raiz.
- A rede de exemplo implanta uma rede Fabric com o Docker Compose. Como os nós estão isolados em uma rede Docker Compose, a rede de teste não está configurada para se conectar a outros nós do Fabric em execução.

### Passos principais do tutorial

1. **Instalar as ferramentas necessárias, como Docker, Docker Compose, Go, Node.js e Python.**

   1.1. **Pré Requisitos**
   - Linux
   - Git: Instale a versão mais recente do git, se ainda não estiver instalada.
     ```shell
     $ sudo apt-get install git
     ```
   - cURL: Instale a versão mais recente do cURL, se ainda não estiver instalada.
     ```shell
     $ sudo apt-get install curl
     ```
   - Docker: Instale a versão mais recente do Docker, se ainda não estiver instalada.
     ```shell
     $ sudo apt-get -y install docker-compose
     ```
   - Após a instalação, confirme que as versões mais recentes dos executáveis do Docker e do Docker Compose foram instaladas.
     ```shell
     $ docker --version
     Docker version 19.03.12, build 48a66213fe
     $ docker-compose --version
     docker-compose version 1.27.2, build 18f557f9
     ```
   - Certifique-se de que o daemon do Docker está em execução.
     ```shell
     $ sudo systemctl start docker
     ```
   - Opcional: Se você deseja que o daemon do Docker seja iniciado quando o sistema for inicializado, use o seguinte comando:
     ```shell
     $ sudo systemctl enable docker
     ```
   - Adicione seu usuário ao grupo Docker.
     ```shell
     $ sudo usermod -a -G docker <nome_de_usuário>
     ```
   - Go - Opcional: Instale a versão mais recente do Go, se ainda não estiver instalada (somente necessário se você for escrever chaincode em Go ou aplicativos SDK em Go).
   - JQ - Opcional: Instale a versão mais recente do jq, se ainda não estiver instalada (apenas necessário para os tutoriais relacionados a transações de configuração de canal).

   1.2. **Instalar o Fabric e o Fabric Samples**

      1.2.1. **Instalar o Fabric:**
      - Crie um diretório de trabalho
        ```shell
        mkdir <nome_pasta>
        cd <nome_pasta>
        ```
      - Baixe o script de instalação:
        ```shell
        curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh &e chmod +x install-fabric.sh
        ```
      - Execute o script para instalar os componentes necessários, como imagens Docker e binários:
        ```shell
        ./install-fabric.sh docker samples binary
        ```
      - Verifique se a instalação foi concluída.

   1.3. **APIs de Contrato do Fabric e APIs de Aplicativos**

      1.3.1. **APIs de Contrato do Fabric**
      - O Hyperledger Fabric oferece várias APIs para o desenvolvimento de smart contracts (chaincode) em várias linguagens de programação. As APIs de smart contract estão disponíveis para Go, Node.js e Java.
      - [Go contract API and documentation](https://hyperledger-fabric.readthedocs.io/en/release-2.5/developapps/smartcontract.html).
      - [Node.js contract API and documentation](https://hyperledger-fabric.readthedocs.io/en/release-2.5/developapps/smartcontract.html).
      - [Java contract API and documentation](https://hyperledger-fabric.readthedocs.io/en/release-2.5/developapps/smartcontract.html).

      1.3.2. **APIs de Aplicações do Fabric**
      - O Hyperledger Fabric oferece uma API de cliente do Fabric Gateway para o desenvolvimento de aplicações em Go, Node.js e Java. Esta API utiliza a capacidade de peer do Gateway introduzida no Fabric v2.4 para interagir com a rede do Fabric e é uma evolução do novo modelo de programação de aplicações introduzido no Fabric v1.4. A API de cliente do Fabric Gateway é a API preferida para o desenvolvimento de aplicações para o Fabric a partir do v2.4.
      - [Node.js SDK and documentation](https://hyperledger-fabric.readthedocs.io/en/release-2.5/developapps/application.html).
      - [Java SDK and documentation](https://hyperledger-fabric.readthedocs.io/en/release-2.5/developapps/application.html).
      - [Go SDK and documentation](https://hyperledger-fabric.readthedocs.io/en/release-2.5/developapps/application.html).

      **Java SDK**
      - Observação: Esta API está obsoleta a partir do Fabric v2.5. Ao desenvolver aplicativos para o Hyperledger Fabric v2.4 e posterior, você deve usar a Fabric Gateway client API.

1.4. **Run Fabric**
   - Baixado o Fabric e os exemplos, é possível começar a executar o Fabric.
   - O tutorial "Executando uma rede de teste" ajuda a entender como as redes do Fabric funcionam.
   - O tutorial "Implantando um contrato inteligente em um canal" ajuda no desenvolvimento e implantação de contratos inteligentes.
   - O tutorial "Executando um aplicativo Fabric" ajuda a desenvolver aplicativos de blockchain em cima de uma rede Fabric.
   - Ambos os tutoriais têm links para explicações mais detalhadas em Conceitos Chave.

## Iniciar a rede de teste

Para iniciar a rede de teste do Hyperledger Fabric, siga as etapas abaixo:

1. **Navegue até o diretório da rede de teste com o seguinte comando:**

    ```shell
    cd fabric-samples/test-network
    ```

2. Nesse diretório, encontra-se um script que cria uma rede Fabric usando as imagens do Docker na máquina local. É possível executar o seguinte comando para ver o texto de ajuda do script:

    ```shell
    ./network.sh -h
    ```

3. Para iniciar a rede, execute o seguinte comando:

    ```shell
    ./network.sh up
    ```

    Este comando cria uma rede Fabric com dois nós pares e um nó de ordenação. Executando o comando, os logs dos nós criados serão exibidos. É necessário executar o script no diretório da rede de teste.

4. Após implantar a rede de teste, é possível examinar os componentes. O seguinte comando lista todos os containers Docker em execução na máquina:

    ```shell
    docker ps -a
    ```

    Serão vistos os três nós criados pelo script.

5. Com a rede de teste implantada com sucesso, é possível criar canais e iniciar a implantação de contratos inteligentes para interagir com a rede do Fabric.

6. Se foi utilizado o chaincode "asset-transfer (basic)" para instalar e iniciar, é possível invocar a função do chaincode (Go) para colocar uma lista inicial de ativos no ledger, executando:

    ```shell
    ./network.sh deployCC -ccl goInitLedger
    ```

    Ou, se preferir usar JavaScript, execute:

    ```shell
    ./network.sh deployCC -ccl javascriptInitLedger
    ```

7. Em seguida, é possível inicializar o ledger com “assets” usando o seguinte comando:

    ```shell
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'
    ```

8. Para consultar o ledger e ver os “assets” adicionados:

    ```shell
    peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllAssets"]}'
    ```

9. Para transferir a propriedade de um “assets” no ledger:

    ```shell
    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"TransferAsset","Args":["asset6","Christopher"]}'
    ```

    A configuração exige que a transação seja assinada por Org1 e Org2.

10. Depois de invocar o chaincode, é possível usar outra consulta para ver como a transação alterou os “assets” no ledger.

11. Para encerrar a rede:

    ```shell
    ./network.sh down
    ```

Isso conclui as etapas para iniciar e interagir com a rede de teste do Hyperledger Fabric.
