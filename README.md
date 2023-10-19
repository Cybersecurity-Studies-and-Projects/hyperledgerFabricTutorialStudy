# Tutorial Hyperledger Fabric

Este repositório contém tutoriais sobre o [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/release-2.5/tutorials.html). Os tutoriais do Fabric podem ser usados por desenvolvedores para começar a criar suas próprias soluções.

## Etapas

1. **Trabalhar com o Fabric implantando a rede de teste em uma máquina local.**
    - Usar os passos fornecidos pelo tutorial de implantação de um contrato inteligente em um  channel.

2. **Executando uma Aplicação Fabric**
    - Fornece uma introdução sobre como usar as APIs fornecidas pelas SDKs do Fabric para invocar contratos inteligentes a partir de aplicações cliente.

3. **Implantando um contrato inteligente em um  channel e Criando um  channel**
    - Podem ser usados para aprender aspectos importantes da administração de uma rede em execução.

4. **Dados privados e CouchDB**
    - Explorar recursos importantes do Fabric.

5. **Implantação de uma rede de produção**
    - Implantar o Hyperledger Fabric em produção.

6. **Atualizar um  channel**
    - Atualização da configuração de um  channel e Atualização do nível de capacidade de um  channel.

7. **Atualizar componentes como nós pares, nós de ordenação, SDKs e outros**
    - Atualização de seus componentes.

8. **Introdução sobre como escrever um contrato inteligente básico**
    - Escrevendo seu Primeiro Chaincode.

# Using the Fabric test network (Usando a rede de teste do Fabric)

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

   - JQ:
     ```shell 
       $ sudo apt-get install jq
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
   - JQ - Opcional: Instale a versão mais recente do jq, se ainda não estiver instalada (apenas necessário para os tutoriais relacionados a transações de configuração de  channel).

   1.2. **Instalar o Fabric e o Fabric Samples**

      1.2.1. **Instalar o Fabric:**
      - Crie um diretório de trabalho
        ```shell
        mkdir <nome_pasta>
        cd <nome_pasta>
        ```
      - Baixe o script de instalação:
        ```shell
        curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh
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
   - O tutorial "Implantando um contrato inteligente em um  channel" ajuda no desenvolvimento e implantação de contratos inteligentes.
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

6. Para encerrar a rede:
    ```shell
    ./network.sh down
    ```

<hr/>
<hr/>
<hr/>

# Implantação de um Contrato Inteligente em um  channel no Hyperledger Fabric usando Java

Os usuários finais interagem com o ledger blockchain, invocando contratos inteligentes. No Hyperledger Fabric, contratos inteligentes são implantados em pacotes chamados chaincode. Organizações que desejam validar transações ou consultar o ledger precisam instalar um chaincode em seus peers. Depois que um chaincode foi instalado nos peers que fazem parte de um  channel, os membros do  channel podem implantar o chaincode no  channel e usar os contratos inteligentes nele para criar ou atualizar assets no ledger do  channel.

Um chaincode é implantado em um  channel usando um processo conhecido como o ciclo de vida do chaincode do Fabric. O ciclo de vida do chaincode do Fabric permite que várias organizações concordem com a forma como um chaincode será operado antes que ele possa ser usado para criar transações. Por exemplo, enquanto uma política de endosso especifica quais organizações precisam executar um chaincode para validar uma transação, os membros do  channel precisam usar o ciclo de vida do chaincode do Fabric para concordar com a política de endosso do chaincode. Para obter uma visão mais detalhada de como implantar e gerenciar um chaincode em um  channel, consulte o [ciclo de vida do chaincode do Fabric](link-para-ciclo-de-vida).

## Etapa 1: Iniciar a Rede

Começaremos implantando uma instância da rede de teste do Fabric. Antes de começar, certifique-se de que você instalou o software necessário seguindo as instruções em [Getting Started with Fabric](link-para-getting-started). Use o seguinte comando para navegar até o diretório da rede de teste dentro do seu clone local do repositório fabric-samples:

```bash 
cd fabric-samples/test-network
```

Para este tutorial, queremos partir de um estado inicial conhecido. O comando a seguir encerrará qualquer contêiner Docker asset ou obsoleto e removerá artefatos gerados anteriormente:

```bash
./network.sh down
```

Você pode então usar o seguinte comando para iniciar a rede de teste:

```bash
./network.sh up createChannel
```

O comando createChannel cria um  channel chamado mychannel com dois membros do  channel, Org1 e Org2. O comando também junta um peer que pertence a cada organização ao  channel. Se a rede e o  channel forem criados com sucesso, você verá a seguinte mensagem nos logs:

```diff
========= Channel successfully joined ===========
```

Agora podemos usar o Peer CLI para implantar o chaincode de transferência de assets (básico) no  channel seguindo as etapas a seguir.

## Etapa 2: Empacotar o Contrato Inteligente
Primeiro, você precisa empacotar seu contrato inteligente em Java. Isso envolve a criação de um arquivo JAR que contém todas as classes e dependências necessárias para seu contrato inteligente. Certifique-se de que seu código-fonte do contrato inteligente esteja corretamente estruturado e inclua todas as classes e recursos necessários.

Para empacotar seu contrato inteligente em Java, você pode usar uma ferramenta de criação de pacotes como o Apache Maven ou o Gradle. Ambos os sistemas de construção são amplamente utilizados na comunidade Java e podem ajudá-lo a gerar um arquivo JAR a partir do seu código-fonte.

### Usando o Apache Maven:
Certifique-se de que o Maven esteja instalado e configurado em seu ambiente.

Navegue até o diretório raiz do seu projeto Java que contém seu contrato inteligente. Suponha que seu projeto esteja organizado da seguinte forma:

```css
meu-projeto-java/
    └── src/
        └── main/
            └── java/
            └── resources/
```
Certifique-se de que seu código-fonte do contrato inteligente esteja corretamente estruturado na pasta src/main/java e que inclua todas as classes e recursos necessários. Por exemplo, você pode ter uma estrutura de diretório semelhante a esta:
```css
meu-projeto-java/
    └── src/
        └── main/
            └── java/
                └── meu/
                    └── contrato/
                        └── ContratoInteligente.java
            └── resources/
                └── META-INF/
                    └── minha-configuracao.yaml
```

Abra um terminal ou prompt de comando e navegue até o diretório raiz do seu projeto Java. Use o seguinte comando para criar um arquivo JAR do seu contrato inteligente:
```bash
mvn clean package
```

Este comando irá compilar seu código-fonte, resolver as dependências definidas em seu arquivo pom.xml (arquivo de configuração do Maven), e gerar um arquivo JAR na pasta target do seu projeto.

Após a conclusão, você encontrará o arquivo JAR do contrato inteligente na pasta target. Este arquivo JAR contém seu contrato inteligente e todas as classes e recursos necessários para implantá-lo no Hyperledger Fabric.

Lembre-se de que o exemplo acima pressupõe que você tenha configurado corretamente seu projeto Java e seu arquivo pom.xml para incluir as dependências necessárias. Certifique-se de adaptar a estrutura do seu projeto e as configurações do Maven de acordo com suas necessidades específicas.

Depois de empacotar seu contrato inteligente em um arquivo JAR, você pode prosseguir com as etapas subsequentes para instalar e implantar o chaincode no Hyperledger Fabric.

### Usando o Gradle:

1. Certifique-se de que o Gradle esteja instalado e configurado em seu ambiente.
2. Navegue até o diretório raiz do seu projeto Java que contém seu contrato inteligente.
3. Execute o comando a seguir para criar um arquivo JAR do seu contrato inteligente:
```bash
gradle build
```

Isso irá compilar seu código-fonte, resolver as dependências e gerar um arquivo JAR na pasta build/libs do seu projeto.

Depois de empacotar seu contrato inteligente em Java, você terá um arquivo JAR que pode ser implantado no Hyperledger Fabric.

## Etapa 3: Instalar o Chaincode nos Peers
Agora que você tem seu contrato inteligente empacotado em um arquivo JAR, é hora de instalá-lo nos peers da rede. Use o seguinte comando para instalar o chaincode nos peers:

```bash
peer lifecycle chaincode install basic_1.0.jar
```

Substitua basic_1.0.jar pelo nome do arquivo JAR gerado na Etapa 2.

Este comando instalará o chaincode nos peers, tornando-o disponível para uso.

## Etapa 4: Aprovar a Definição do Chaincode
Antes de poder implantar o chaincode em um  channel, ele deve ser aprovado por uma maioria das organizações no  channel. Use o seguinte comando para aprovar a definição do chaincode:

```bash
peer lifecycle chaincode approveformyorg -o orderer.example.com:7050 --ordererTLSHostname orderer.example.com --channelID mychannel --name basic --version 1.0 --package-id basic_1.0:42a29e5f26d9f51e9b1a2f32d2c411e1a06a2e120869a9f8c9900da0e30e992b --sequence 1
```
Certifique-se de substituir os valores apropriados:

* -o: O endereço do nó orderer.
* --ordererTLSHostname: O nome de host TLS do nó orderer.
* --channelID: O nome do  channel (neste exemplo, é mychannel).
* --name: O nome do chaincode (neste exemplo, é basic).
* --version: A versão do chaincode (neste exemplo, é 1.0).
* --package-id: O ID do pacote gerado durante a instalação do chaincode. Você pode obtê-lo no resultado da etapa anterior.
- Este comando aprovará a definição do chaincode para a organização do usuário. Repita o processo de aprovação para todas as organizações que fazem parte do  channel.

## Etapa 5: Comprometer a Definição do Chaincode no  channel
Depois que a definição do chaincode é aprovada por todas as organizações necessárias, você pode comprometê-la no  channel com o seguinte comando:

```bash
peer lifecycle chaincode commit -o orderer.example.com:7050 --ordererTLSHostname orderer.example.com --channelID mychannel --name basic --version 1.0 --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
```

Certifique-se de substituir os valores apropriados:

* -o: O endereço do nó orderer.
* --ordererTLSHostname: O nome de host TLS do nó orderer.
* --channelID: O nome do  channel (neste exemplo, é mychannel).
* --name: O nome do chaincode (neste exemplo, é basic).
* --version: A versão do chaincode (neste exemplo, é 1.0).
* --sequence: A sequência de aprovação.
* --tls: Indica que as conexões são feitas usando TLS.
* --cafile: O caminho para o arquivo de certificado da autoridade de certificação (CA) do orderer.
* --peerAddresses: Os endereços dos peers que pertencem à organização.
* --tlsRootCertFiles: Os caminhos para os arquivos de certificado raiz TLS dos peers.
- Esta etapa compromete o chaincode no  channel e o torna disponível para uso por seus membros.

## Conclusão
Neste tutorial, você aprendeu a implantar um contrato inteligente Java em um  channel no Hyperledger Fabric. Começamos iniciando uma rede de teste do Fabric, empacotamos o contrato inteligente em um arquivo JAR, instalamos o chaincode nos peers, aprovamos a definição do chaincode e o comprometemos no  channel. Este é um processo fundamental para criar um ambiente blockchain no Fabric, onde os contratos inteligentes podem ser executados e usados para criar e atualizar assets no ledger do  channel.

Lembre-se de que este tutorial se concentra na implantação de contratos inteligentes em Java. Se você estiver usando outra linguagem, como Go, JavaScript ou TypeScript, o processo será diferente, mas as etapas gerais de implantação do chaincode em um  channel no Hyperledger Fabric serão semelhantes.

Agora que você implantou com sucesso seu contrato inteligente, você pode começar a interagir com ele, criando assets, lendo do ledger e realizando outras operações definidas no contrato. Este é apenas o começo de seu projeto de blockchain e há muito mais a explorar no Hyperledger Fabric e nas aplicações de blockchain.