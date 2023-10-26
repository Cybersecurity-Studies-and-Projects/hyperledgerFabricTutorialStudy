# Blockchain chaincode para desenvolvedores Java - Parte 01

## Pré-requisitos
1. Conhecer as bases da tecnologia blockchain e do framework Hyperledger Fabric.
2. Conhecimento e experiência intermediária em programação Java tanto com a linguagem quanto com a plataforma.
3. Familiaridade com ou (idealmente) proficiente em usar:
    * Eclipse IDE
    * Docker e Docker Compose
    * Gradle
    * Linha de comando do Linux
    * SoapUI ou outro software de cliente HTTP, como o Postman

Existem outras opções para os componentes da pilha. Esse tutorial utiliza o Docker como o ambiente de contêiner de rede, mas outra opção é Vagrant com VirtualBox.

Considerando que Docker é um ambiente de contêiner, Vagrant usa a virtualização. Quando combinado com VirtualBox, um ambiente de virtualização fornece um nível diferente de controle sobre o ambiente computacional que alguns desenvolvedores preferem (o que o torna uma escolha ideal para desenvolvedores do fabric).

## Configurando o ambiente de desenvolvimento

Etapas:
1. Configurar o ambiente de rede – para executar a sua rede blockchain local.
2. Instalar o software de compilação – para fazer o build do  chaincode.
3. Instalar um cliente HTTP – para invocar transações no chaincode.
4. Iniciar a rede blockchain.
5. Construir o cliente JAR Java shim.

### Configurar o ambiente de rede

Usar o Docker junto com as imagens de componentes de rede blockchain pré-construídas do Docker Hub, para executar sua rede blockchain local.

#### Instalar o Docker

https://docs.docker.com/engine/getstarted/step_one 

#### Insalar o software de compilação

Para o sistema de compilação, o Hyperledger Fabric utiliza Gradle. O Gradle é um sistema de automação de compilação que combina sintaxe simples para especificar componentes de compilação, juntamente com os melhores recursos do Apache Ant e do Apache Maven para criar um poderoso sistema de compilação fácil de usar.

```shell
    $ sudo unzip -d /opt/gradle gradle-2.14-bin.zip

    $ echo "export GRADLE_HOME=/opt/gradle/gradle-2.14" >> ~/.bashrc
    echo "export PATH=\$PATH:\$GRADLE_HOME/bin" >> ~/.bashrc
    source ~/.bashrc

    $ gradle -v
```

informar ao Docker Compose como compor e executar a rede de blockchain. A definição da rede está em YAML, e você deve nomeá-la docker-compose.yml. Você pode chamar o arquivo de outra coisa, mas quando você inicia o Docker Compose, você deve especificar o sinalizador -f. Recomenda-se permanecer com o padrão, que é docker-compose.yml.

docker-compose.yml:
```yml
membersrvc:
  image: hyperledger/fabric-membersrvc
  ports:
    - "7054:7054"
  command: membersrvc
vp0:
  image: hyperledger/fabric-peer
  ports:
    - "7050:7050"
    - "7051:7051"
    - "7053:7053"
  environment:
    - CORE_PEER_ADDRESSAUTODETECT=true
    - CORE_VM_ENDPOINT=unix:///var/run/docker.sock
    - CORE_LOGGING_LEVEL=DEBUG
    - CORE_PEER_ID=vp0
    - CORE_PEER_PKI_ECA_PADDR=membersrvc:7054
    - CORE_PEER_PKI_TCA_PADDR=membersrvc:7054
    - CORE_PEER_PKI_TLSCA_PADDR=membersrvc:7054
    - CORE_SECURITY_ENABLED=false
    - CORE_SECURITY_ENROLLID=test_vp0
    - CORE_SECURITY_ENROLLSECRET=MwYpmSRjupbT
  links:
    - membersrvc
  command: sh -c "sleep 5; peer node start --peer-chaincodedev"
```

Esse arquivo diz ao Docker Compose para definir dois serviços:

**membersrvc**: O node de serviços integrantes que fornece serviços de associação, especificamente uma autoridade de certificação/certificate authority (CA), que é responsável pelo manuseio de toda a logística criptográfica (como emissão e revogação de certificados). A imagem Docker pré-construída que você usará para isso é chamada hyperledger/fabric-membersrvc.
**vp0**: O único node de pares de validação na rede. Para fins de desenvolvimento, não precisamos de uma rede de pares de validação extravagante; um único par será suficiente. A imagem Docker pré-construída que você usará para isso é chamada de hyperledger/fabric-peer.

Um grande número de variáveis de ambiente são definidas pelo par vp0. Repare que a variável CORE_LOGGING_LEVEL está definida como DEBUG. Isso produz uma grande quantidade de resultado, que pode ser útil às vezes. No entanto, se desejar menos resultado, altere o nível para INFO. Consulte “Controle de log” nos documentos de Configuração do Hyperledger para obter mais informações sobre os níveis de log.

o valor CORE_SECURITY_ENABLED é false. Isso significa que o fabric não exigirá que você envie qualquer tipo de credenciais de usuário final. A segurança está além do escopo deste tutorial, mas se você estiver interessado em aprender mais, você pode verificar esta nota sobre a funcionalidade de segurança para suas solicitações chaincode.

alterar qualquer um desses valores de seus padrões (especialmente os valores de porta) pode fazer com que os exemplos neste tutorial não funcionem. Uma rede de blockchain é um conjunto de componentes de software distribuídos que requerem uma comunicação coordenada precisa. Eu recomendo que você não altere os valores de porta de seus valores padrão até que você entenda como todos os componentes do fabric interoperam.

Agora que a definição de rede blockchain está feita, você está pronto para iniciar sua rede blockchain local. Para fazer isso, execute Docker Compose. Navegue até o seu $GOPATH e execute este comando:

```shell
    $ docker-compose up
```

Este resultado informa que a rede está funcionando e em execução, e pronta para aceitar solicitações de registro de chaincode.

Nota: as linhas realçadas só devem aparecer na primeira vez que você executar sua rede blockchain, porque Docker tem que baixar as imagens do Docker Hub. Uma vez que elas estão baixadas no seu computador, Docker só as puxará se as imagens do Docker Hub forem mais recentes do que as que você tem no seu computador.

Agora você está pronto para construir o cliente JAR Java shim, que permite que o seu chaincode de linguagem Java se comunique com o framework Hyperledger Fabric.

#### Criar o cliente JAR Java shim

Antes de poder executar os exemplos chaincode, você precisa obter o código-fonte mais recente do repositório GitHub da Hyperledger.

Primeiro, você precisará clonar o Hyperledger Fabric em sua máquina local para construir o seu chaincode (Nota: Esta é uma medida temporária; em algum momento, o cliente JAR Java shim deverá ser acessível a partir do repositório Maven central).

Execute este comando para criar a estrutura de diretórios que os scripts de build de fabric esperados:

```shell
    $ export GOPATH=~/home/mychaincode

    $ mkdir -p $GOPATH/src/github.com/hyperledger
```

cd $GOPATH/src/github.com/hyperledger

```shell
    $ cd $GOPATH/src/github.com/hyperledger
```

A partir daqui você precisa recuperar o código-fonte do Hyperledger para que você possa construir o cliente JAR Java shim.

```shell
    $ git clone https://github.com/hyperledger/fabric.git -b v0.6
```

Agora você está pronto para criar o cliente JAR shim do chaincode Java. Navegue até $GOPATH/src/github.com/hyperledger/fabric/core/chaincode/shim/java e execute estes dois comandos:

OBS.: É necessario ter o Gradle e o Java compatíveis com a versão o repositório Hyperledger.
Para este caso, é nessário ter:
1. Gradle 2.14
2. Java 1.8

```shell
    $ cd $GOPATH/src/github.com/hyperledger/fabric/core/chaincode/shim/java

    $ gradle -b build.gradle clean

    $ gradle -b build.gradle build
```

A última coisa que a compilação faz é adicionar o cliente JAR shim ao seu repositório Maven local. Neste ponto, você está pronto para criar o seu chaincode. A menos que você atualize o código fonte do fabric em algum momento no futuro, ou apenas queira reconstruir o cliente JAR shim novamente por algum motivo, você não terá que executar o cliente JAR Java shim compilado novamente.

##### Fazer deploy e executar um exemplo de chaincode Java

Agora que definiu e iniciou a sua rede blockchain local e construiu e instalou o cliente JAR Java shim no seu repositório Maven local, está pronto para criar, registrar e invocar transações num dos exemplos de chaincode Java que é enviado com o Hyperledger Fabric que você baixou anteriormente.

**Passos para seguir:**

1. Crie o exemplo usando Gradle
    * Navegue até o diretório
    ```shell
    $ cd $GOPATH/src/github.com/hyperledger/fabric/examples/chaincode/java/Example
    ```
    * Em seguida, inicie a compilação Gradle através da linha de comando utilizando este comando:
    ```shell
    $ gradle -b build.gradle build
    ```
    * A compilação cria uma distribuição autônoma que está localizada dentro do diretório compilar/distribuições em duas formas: um arquivo TAR e um arquivo ZIP e cada um desses arquivos contém tudo o que você precisa para executar o chaincode, incluindo um script para conduzi-lo chamado Example.

2. Registre o exemplo
    * Certifique-se de que a sua rede blockchain local está em execução. Se não estiver, você precisará iniciá-la. Consulte a seção “Inicie a rede blockchain” se precisar de uma atualização.
    ```shell
    $ cd $GOPATH/src/github.com/hyperledger/fabric/examples/chaincode/java/Example

    $ cd build/distributions/

    $ unzip Example.zip 
    ```
    * A distribuição contém tudo o que você precisa para executar o chaincode autonomamente (em seu próprio processo), juntamente com todos os JARs dependentes.
    * Para registrar o chaincode exemplo, dentro da pasta compilar/distribuições, execute o seguinte script:
    ```shell
    $ docker-compose up -d
    $ ./Example/bin/Example
    ```
    * Isso executa um processo autônomo que registra o chaincode exemplo com a rede local blockchain. Você deve ver o resultado da janela do terminal assim: