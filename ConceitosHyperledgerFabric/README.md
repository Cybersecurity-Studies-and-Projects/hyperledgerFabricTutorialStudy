# Conceitos Hyperledger Fabric

## Hyperledger Fabric

- A plataforma de código aberto Hyperledger Fabric é utilizada para **livro-razão distribuído** em contextos empresariais.
- Permite a execução de **contratos inteligentes** (chaincodes) com **alta modularidade**.
- A plataforma é mantida pela **Linux Foundation** com o apoio de um comitê técnico diversificado.
- A natureza do Hyperledger Fabric é **permissionada**, baseada na **identificação dos participantes** para estabelecer confiança na rede.

### Componentes

* Apresenta a coexistência de cadeias de blocos isoladas e seus papéis.

#### Rede de Peers

- Uma rede de blockchain é composta por um conjunto de **peers**, que hospedam **contratos inteligentes (chaincodes)** e **livros-razão (ledgers)**.
- Os chaincodes e ledgers encapsulam informações e processos compartilhados na rede.
- Uma **aplicação (client app)** se comunica com os peers para acessar ledgers e chaincodes, permitindo a invocação de contratos e a submissão de transações na rede.

#### Organizações

- Diferentes **peers podem se unir para formar uma organização (org)** com um propósito e nível de informações comuns.
- As organizações são essenciais para a rede, pois a construção e gerenciamento da rede dependem do compartilhamento de recursos entre elas.
- As organizações compartilham recursos, principalmente peers, para a coletividade da rede.

#### Canais

- Ao formar um **consórcio entre uma ou mais organizações**, é possível criar um **canal (channel)**.
- Os peers de diferentes organizações podem se comunicar e atender requisições em um mesmo contexto, realizando transações na blockchain.

#### Identidades e Certificados

- Cada peer possui uma **identidade associada através de um certificado digital** fornecido pela sua **autoridade certificadora (CA)**.
- A conexão de um peer em um canal é validada por um componente da plataforma chamado **Membership Service Provider (MSP)**.

#### Orderer

- Além dos peers, a plataforma possui um tipo especial de peer chamado **orderer**.
- O orderer é responsável por manter o estado consistente dos ledgers na rede.
- Ele realiza um processo de três fases, incluindo **endosso, agrupamento e propagação de transações**.
- O orderer age como mediador para alcançar um consenso sobre o conteúdo das transações entre os peers da rede.

### Arquitetura de Transações

#### Visão Geral

A arquitetura de transações do **Hyperledger Fabric** é conhecida como **executar-ordenar-validar**. Essa arquitetura oferece resiliência, flexibilidade, escalabilidade e lida com problemas de confidencialidade. Ela é composta por três etapas essenciais: **execução**, **ordenação** e **validação**.

#### Fase de Execução

Na **fase de execução**, um cliente envia uma proposta de transação para um ou mais peers de endosso. Os endossadores simulam a proposta, produzindo um conjunto de escrita (writeset) e um conjunto de leitura (readset). Cada endossador assina criptograficamente um endosso contendo esses conjuntos e os envia de volta ao cliente. O cliente coleta endossos até que a política de endosso especificada pelo chaincode seja satisfeita.

#### Fase de Ordenação

Na **fase de ordenação**, as transações e endossos são transmitidos e um consenso é estabelecido, mesmo com possíveis orderers defeituosos. As transações são agrupadas em blocos, criando uma sequência encadeada através de hash. Isso melhora a vazão do protocolo de transmissão e garante a total ordenação dos blocos em um canal.

#### Fase de Validação

A **fase de validação** avalia a política de endosso estabelecida pelo sistema de validação de chaincode (VSCC). Realiza a verificação de conflitos nos conjuntos de leitura e escrita e atualiza o livro-razão com as transações válidas.

#### Livro-Razão do Hyperledger Fabric

O **livro-razão do Hyperledger Fabric** contém todas as transações, incluindo as inválidas. Isso diferencia o Fabric de outras blockchains, tornando a plataforma adequada para casos que exigem o rastreamento de transações inválidas.

### Características da Plataforma

#### Visão Geral

O **Hyperledger Fabric** é uma plataforma altamente modular e configurável que oferece uma ampla gama de recursos para atender a diversos casos de uso industriais, incluindo setores como bancos, finanças, seguros, saúde, recursos humanos, cadeias de suprimentos e serviços de música digital.

#### Arquitetura de Transações

- A arquitetura **executar-ordenar-validar** permite uma execução paralela eficiente de transações, eliminando o não-determinismo na execução.
- A eliminação do não-determinismo possibilita o uso de linguagens de programação de propósito geral, como Java, Go e Node.js, tornando o desenvolvimento de contratos inteligentes mais acessível.
- O suporte a **protocolos de consenso plugáveis** torna a plataforma altamente personalizável para atender a casos de uso específicos e diferentes modelos de confiança.
- O Hyperledger Fabric **não depende do uso nativo de criptomoedas** para incentivar o custo de mineração ou a execução de contratos inteligentes, reduzindo riscos de ataques e equilibrando o custo operacional em relação a outros sistemas distribuídos.
- A plataforma oferece flexibilidade na escolha de serviços de ordenação, como **Solo, Raft ou Kafka**, permitindo a adaptação a casos de uso específicos.

#### Ledger

- O **ledger** do Hyperledger Fabric consiste em duas partes: o **Estado Global** e a **Blockchain**.
- O **Estado Global** é um banco de dados que armazena o valor atual do conjunto de estados do ledger.
- A **Blockchain** registra todas as transações que alteram o estado global em blocos anexados e é imutável em relação ao conteúdo das transações.

- O uso do **estado global** simplifica o acesso a informações, permitindo que as aplicações acessem facilmente o valor atual de um estado específico sem a necessidade de analisar toda a blockchain.
- A plataforma oferece as opções de bancos de dados **LevelDB e CouchDB** para representar o estado global do ledger.

Essas características tornam o Hyperledger Fabric uma plataforma altamente flexível e adaptável, pronta para atender a uma ampla variedade de casos de uso e requisitos de desenvolvimento específicos.

### Ferramentas para Desenvolvimento

#### Virtualização Baseada em Containers

- O **Hyperledger Fabric** utiliza a **virtualização baseada em containers** para implementar sua rede e criar e comunicar seus artefatos.
- Um **container** é uma unidade de software que agrupa código e suas dependências, permitindo que as aplicações sejam executadas de forma rápida e confiável em diferentes ambientes computacionais.
- O uso de **imagens Docker** permite a utilização de ferramentas como **Kubernetes** e **Docker Compose** para orquestrar conjuntos de containers.

#### Linguagens de Programação

- O Hyperledger Fabric suporta o uso de **linguagens de programação de propósito geral**, como **Java**, **Node.js** e **Go**, para o desenvolvimento de contratos inteligentes.

#### Linguagem Go

- A **linguagem Go** é amplamente utilizada no desenvolvimento de aplicações no Hyperledger Fabric.
- O Go é um projeto open source que visa tornar os programadores mais produtivos, com mecanismos de concorrência que facilitam a escrita de programas para sistemas multicore e em rede.
- A plataforma disponibiliza **kits de desenvolvimento de software (SDK)** para apoiar o desenvolvimento de aplicações relacionadas ao Hyperledger Fabric. Existem três SDKs relacionados às linguagens mencionadas anteriormente e um SDK para a linguagem Python (ainda não oficialmente liberado).

#### Produtos e Serviços

- Mantenedores da plataforma, como a **IBM**, oferecem produtos para desenvolvimento e gerenciamento de aplicações no Hyperledger Fabric.
- A **IBM Blockchain Platform** disponibiliza interfaces e serviços em nuvem para hospedagem e gerenciamento de componentes do Hyperledger Fabric, juntamente com ferramentas de desenvolvimento.

Essas ferramentas e linguagens tornam o desenvolvimento de aplicações no Hyperledger Fabric mais eficiente e oferecem suporte para diversas linguagens de programação, simplificando o processo de criação e implantação de contratos inteligentes e redes blockchain.
