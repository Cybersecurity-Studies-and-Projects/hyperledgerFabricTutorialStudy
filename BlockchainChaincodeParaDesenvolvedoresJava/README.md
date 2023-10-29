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

##### Instalar o Gradle

```shell
    $ sudo apt-get install gradle
```

```shell
    $ gradle -v
```

##### Instalar o cliente HTTP

O cliente HTTP, permite que o chaincode se comunique com a interface REST do Hyperledger blockchain fabric. O navegador pode emitir um HTTP GET, mas para interagir com o fabric é necessário fazer POST nas mensagens. Isso significa que é necessário um cliente HTTP.

O cliente HTTP para este tutorial é **SoapUI**, que fornece uma edição comunitária livre que é poderosa, fácil de usar e contém muitos recursos.


##### Iniciar a rede Blockchain

1. Definir a configuração da rede.

Criar um diretório que servirá como a raiz de todo o código fonte que será usado para o desenvolvimento do chaincode.

Defina a variável de ambiente GOPATH para este caminho.

**O GOPATH**

O Hyperledger Fabric é escrito em Go, e o GOPATH é um termo frequentemente nos documentos do Hyperledger.

O GOPATH é a raiz do seu ambiente Go. Código fonte, binários e outros pacotes Golang são todos referenciados em relação a este caminho.

No Linux:

```shell
    $ export GOPATH=~/home/mychaincode
```

Informar ao Docker Compose como compor e executar a rede de blockchain. A definição da rede está em YAML, e você deve nomeá-la docker-compose.yml. Você pode chamar o arquivo de outra coisa, mas quando você inicia o Docker Compose, você deve especificar o sinalizador -f.Recomensa-se permanecer com o padrão, que é docker-compose.yml.