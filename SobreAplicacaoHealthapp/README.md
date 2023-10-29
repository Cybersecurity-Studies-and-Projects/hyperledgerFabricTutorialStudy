
# A aplicação Healthapp

## Aplicação de Assistência à Saúde com Hyperledger Fabric - Ùnico host

### Aplicação

A aplicação final deste projeto (Healthapp) consiste no desenvolvimento de uma aplicação web que opera em conjunto com uma rede instanciada na plataforma Hyperledger Fabric para o armazenamento e acesso a dados de saúde. A aplicação utiliza contratos inteligentes para a concessão de acesso a dados de saúde por meio da blockchain permissionada, proporcionando praticidade aos usuários por meio de uma interface web.

### Visão Geral

A aplicação desenvolvida utiliza contratos inteligentes para a concessão de acesso a dados de saúde por meio de uma interface de usuário que interage com uma rede baseada no Hyperledger Fabric. Os contratos inteligentes visam dar maior controle ao paciente sobre o acesso aos seus próprios registros de saúde, permitindo que compartilhe esses dados com profissionais de saúde ou outros interessados.

A aplicação é composta por módulos cliente e servidor. O módulo cliente é responsável por renderizar a interface da aplicação, enquanto o módulo servidor comunica-se com a rede de blockchain e invoca os contratos inteligentes por meio do fabric-sdk-go da Hyperledger.

### Arquitetura

A arquitetura da aplicação permite que os usuários utilizem contratos inteligentes para gerenciar o acesso aos dados de saúde de forma mais prática por meio de uma interface web. O fluxo de operações é o seguinte:

1. O módulo cliente envia uma requisição de invocação de chaincode gerada pelo usuário.
2. O módulo servidor recebe a requisição e envia uma proposta de transação à rede Hyperledger Fabric.
3. A rede Hyperledger Fabric processa a transação e envia uma resposta ao módulo servidor.
4. O módulo servidor envia a resposta ao módulo cliente, que a apresenta ao usuário.

As funcionalidades da aplicação incluem o cadastro de pacientes, identificação de pacientes, cadastro de interessados, remoção de interessados e verificação de dados, cada uma correspondendo a funções específicas nos chaincodes da Hyperledger Fabric.

### Rede Hyperledger Fabric

A rede estabelecida na plataforma Hyperledger Fabric utilizada na aplicação é composta por quatro entidades: dois peers, um orderer e uma autoridade certificadora (CA). Essa rede fornece as propriedades de uma blockchain permissionada e permite a execução das transações.

### Desenvolvimento

O módulo servidor da aplicação utiliza o fabric-sdk-go para a comunicação com a rede de blockchain Hyperledger Fabric e a invocação dos contratos inteligentes. A aplicação é representada como uma aplicação cliente na rede, permitindo operações como a criação e inserção de entidades no canal.

A aplicação oferece uma interface de usuário baseada em HTML, CSS (com o uso do framework Bootstrap) e recursos da linguagem de programação Go. Cada funcionalidade é representada por uma função correspondente nos chaincodes da Hyperledger Fabric.

Através do uso do fabric-sdk-go, o módulo servidor envia propostas de transação à rede para processamento, e a resposta é apresentada ao usuário por meio do módulo cliente. A aplicação utiliza a arquitetura de transações executar-ordenar-validar para garantir a consistência e confiabilidade das operações.

### Implementação

A implementação dos módulos cliente (client) e servidor (server) desta aplicação foi realizada usando a linguagem de programação Go. O Go proporcionou praticidade ao desenvolvimento, pois permitiu a renderização de templates HTML, a construção da interface da aplicação e o atendimento a requisições HTTP por meio do pacote "net/http".

A interface da aplicação foi desenvolvida com HTML e CSS, com a ajuda do framework Bootstrap. O Go também é utilizado para o processamento de requisições HTTP no módulo servidor.

A comunicação com a rede Hyperledger Fabric e a invocação dos contratos inteligentes são realizadas por meio do fabric-sdk-go, um SDK desenvolvido pela Hyperledger para a linguagem Go.

A infraestrutura da rede de blockchain foi implementada com a ajuda das ferramentas cryptogen e configtxgen fornecidas pela Hyperledger. Essas ferramentas geraram o material criptográfico necessário para cada componente da rede, incluindo certificados de entidades, o bloco gênesis da blockchain e a definição de âncoras.

Nesta implementação, foi utilizado o serviço de ordenação Solo, que é ideal para aplicações de prova de conceito. Os peers da rede utilizam o banco de dados LevelDB para armazenar o estado global da blockchain.

A rede de blockchain é composta por quatro entidades: dois peers (peer0 e peer1), um orderer e uma autoridade certificadora (CA). Todos esses componentes são instanciados como containers Docker.

#### Modelo Único Host

Um dos modelos de implementação desta rede é o chamado "Único Host", onde toda a rede é instanciada em uma única máquina. Nesse modelo, as entidades que compõem a rede se comunicam através da rede padrão do Docker. O Docker Compose é usado para orquestrar os containers da rede.

Este modelo permite a instância rápida e prática da rede em um único host, simplificando a complexidade da virtualização baseada em containers. O componente "Fabric Client App" representa a aplicação neste modelo e é responsável por interagir com a rede Hyperledger Fabric por meio do fabric-sdk-go.

### Diagrama de Componentes UML para o Modelo Único Host

O Diagrama de Componentes UML apresenta os artefatos utilizados na implementação do modelo Único Host. O arquivo "docker-compose.yaml" especifica os parâmetros necessários para criar os containers que representam as entidades da rede. A aplicação, representada pelo componente "Fabric Client App," utiliza o arquivo "config.yaml" para obter informações criptográficas e de endereçamento das entidades da rede.

