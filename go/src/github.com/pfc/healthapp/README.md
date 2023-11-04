<h1>Healthapp</h1>
<p>Autor: Tauan Coelho Coutinho Borges</p>

<h2>Índice</h2>
<ol>
  <li><a href="#Motivação">Motivação</a></li>
  <li><a href="#Descrição">Descrição</a></li>
  <li><a href="#aplicação">Aplicação</a></li>
  <ol>
    <li><a href="#Pré-requisitos">Pré-requisitos</a></li>
    <li><a href="#clonando-o-repositório">Clonando o repositório</a></li>
    <li><a href="#iniciando-a-aplicação">Iniciando a aplicação</a></li>
    <ol>
      <li><a href="#único-host">Único Host</a></li>
      <li><a href="#multi-host">Multi Host</a></li>
    </ol>
  </ol>
  
</ol>

<h2>Motivação</h2>

<p>Esta aplicação foi desenvolvida como produto para o meu Trabalho de Conclusão de Curso do curso de Ciência da Computação pela Universidade Federal de Goiás.</p>
<p>A pesquisa a qual esta aplicação pertence está relacionada à área de Sistemas Distribuídos, especificamente a tecnologia de Blockchain.</p>
<p>Esta aplicação tem como objetivo demonstrar a utilização de contratos inteligentes para o controle de acesso aos dados relacionados à cuidados em saúde através da plataforma Hyperledger Fabric.</p>

<h2>Descrição</h2>

<p>A aplicação Healthapp consiste de uma interface web para a invocação de contratos inteligentes de forma prática e ágil, através de um servidor http que utiliza a rede Hyperledger Fabric previamente estabelecida para atender tais requisições.</p>
<p>O desenvolvimento desta aplicação foi majoritariamente realizado utilizando recursos da linguagem de programação Go, como seus módulos 'net/http' e 'html/template', e como peça fundamental o fabric-sdk-go (Software Development Kit disponibilizado pela Hyperledger).</p>
<p>Para obter total entendimento sobre a aplicação, espera-se que o leitor tenha conhecimento prévio sobre a plataforma Hyperledger Fabric e seu funcionamento. Caso contrário, recomendo que visite a documentação da própria Hyperledger Fabric como referência, realizando o tutorial "Build Your First Network" como passo principal, onde são abordados os principais artefatos práticos da plataforma.</p>

<h2>Aplicação</h2>

<h3>Pré-requisitos</h3>

<p>Esta aplicação foi desenvolvida utilizando o sistema operacional Linux, especificamente a distribuição Ubuntu 20. Para utilizar a aplicação, utilize o sistema operacional Linux (como esta aplicação não foi testada em outras distribuições, utilize preferencialmente a distribuição Ubuntu).</p>

<p>Caso você já tenha realizado o tutorial "Build Your First Network" em sua máquina, provavelmente já possui os pré-requisitos necessários instalados. Caso contrário, para utilizar esta aplicação é necessário possuir em sua máquina o <strong>ambiente de desenvolvimento Go, Docker e Docker-compose</strong> instalados.

<p>Para garantir o sucesso do próximo passo, recomendo que configure corretamente o workspace da linguagem Go de acordo com <a href="https://www.digitalocean.com/community/tutorials/como-instalar-o-go-e-configurar-um-ambiente-de-programacao-local-no-ubuntu-18-04-pt">este tutorial</a> (especificamente o passo 2 apresentado).</p> 

<h3>Clonando o repositório</h3>

<p>Dentro do seu workspace go (como especificado no tutorial presente no final do passo anterior), siga os seguintes passos:

<ol>
  <li>Entre na pasta 'src'</li>
  <li>Dentro da pasta 'src', entre na pasta 'github.com' (caso não exista a pasta 'github.com', crie uma nova e insira este mesmo nome)</li>
  <li>Dentro da pasta 'github.com', crie uma pasta denominada 'pfc' (a pasta deve ter exatamente este nome sem as aspas simples!)</li>
  <li>Entre na pasta 'pfc'</li>
  <li>Clone este repositório</li>
</ol>

<p>Ao seguir estes passos, seu workspace deve ter a seguinte subestrutura:</p>

```
go:
  scr:
    github.com:
            pfc:
               healthapp
```

<p>Como os materiais criptográficos foram previamente gerados (com o intuito de simplificar a inicialização e a utilização da solução em múltiplos hosts que será apresentada a seguir), é de extrema importância que o diretório tenha a mesma estrutura descrita anteriormente.</p>

<h3>Iniciando a aplicação</h3>

<p>Esta aplicação possui duas arquiteturas: <b>Único Host e Multi Host</b></p>

<p>A arquitetura Único Host foi projetada para ser exetucada em uma única máquina, onde os quatro containers estarão instanciados em uma rede criada pelo docker compose nesta máquina.</p>

<p>A arquitetura Multi Host foi projetada para ser executada em duas máquinas, onde cada host possuirá dois containers. Estes hosts estarão conectados através da formação de um cluster pela ferramenta <b>Docker Swarm</b>.</p>

<p>A execução para cada arquitetura deve ser feita da seguinte forma:</p>

<h4>Único Host</h4>

<p>1 - Dentro do repositório 'healthapp', execute o seguinte comando:</p>

```go build```

<p>Este comando irá realizar o download de todas as dependências necessárias (especificadas no arquivo go.mod) e então irá compilar o arquivo 'healthapp'. Pode ser que este processo demore algum tempo, então aguarde.</p>

<p>Para continuar a execução iremos utilizar dois terminais (Terminal 1 e Terminal 2) para acompanhar a aplicação em detalhes.</p>

<p>2 - No Terminal 1, execute o seguinte comando:</p>

```sudo docker-compose up```

<p>Este comando irá criar os devidos docker containers (orderer, ca, peer0_org1, peer1_org1) e inseri-los na rede default criada pela ferramenta. A partir do momento da criação, será possível observar os logs gerados pelo containers. Este container será utilizado para a visualização dos logs (não será possível executar nada através deste terminal).</p>

<p>3 - No Terminal 2, execute o seguinte comando: </p>

```./healthapp```

<p>Este comando irá executar o arquivo go compilado. De acordo com a execução, será possível acompanhar os processos necessários para criação da rede Hyperledger Fabric atráves da aplicação (como criação do canal, instanciação do chaincode, etc). Por fim, o servidor irá especificar no terminal o endereço em que a interface web estará disponível (http://localhost:3000/).</p>

<p>Acesse a aplicação através do endereço especificado pelo navegador e explore suas funcionalidades através de sua interface.</p>

<p>Para <b>finalizar</b> a aplicação, execute os seguintes comando:</p>

<p>1 - No terminal 2, pressione as teclas 'ctrl+c' para finalizar o servidor.</p>

<p>2 - No terminal 1, pressione as teclas 'ctrl+c' para parar os containers.</p>

<p>3 - No terminal 1, com o fim do processo de parada dos containers, utilize o seguinte comando:</p>

```sudo docker-compose down```

<p>Este comando irá excluir os containers criados e a rede, encerrando a aplicação.</p>

<h4>Multi Host</h4>

<p>Para utilizar esta arquitetura, serão necessarios dois hosts. Estes hosts deverão possuir todos os pré requisitos instalados. Recomendo executar a arquitetura de único host para verificar a exatidão da instalação anterior nos determinados hosts.</p>

<p>Para facilitar a compreensão, iremos denominar as duas máquinas em Host1 e Host2.</p>

<p>A ferramenta Docker Swarm será utilizada para criarmos um cluster formado pelos Hosts.</p>

<p>1 - No Host1, execute o seguinte comando:</p>

```sudo docker swarm init --advertise-addr <ENDERECO-IP>```

<p>Substitua o campo <ENDERECO-IP> pelo endereço IP do Host1, este comando irá iniciar o swarm.</p>

<p>Este comando irá gerar um output para inserir um novo host ao cluster como <i>worker</i>, porém iremos utilizar o seguinte comando para inclusão de um novo host:</p>

```sudo docker swarm join-token manager```

<p>Este comando irá possibilitar inserir um novo host como <i>manager</i> (possibilitando o mesmo nível de permissões). Copie o output deste comando e execute no Host2.</p>

<p>Ao executar o output no Host2, execute o seguinte comando:</p>

```sudo docker node ls```

<p>Caso estejam presentes os dois hosts na listagem, a criação do cluster teve sucesso.</p>

<p>2 - No Host1, execute o seguinte comando:</p>

```sudo docker network create -d overlay --attachable healthnet```

<p>Este comando irá criar a rede com o driver <i>overlay</i> denominada healthnet a qual os containers irão se conectar para realizar a comunicação.<p>

<p>3 - No Host1, execute o seguinte comando:</p>

```sudo docker-compose -f pc1-compose.yaml up```

<p>Este comando irá iniciar os containers 'orderer' e 'ca' e conecta-los à rede healthnet, de acordo com a especificação do arquivo pc1-compose.yaml.</p>

<p>4 - No Host2, execute o seguinte comando:</p>

```sudo docker-compose -f pc2-compose.yaml up```

<p>Este comando irá iniciar os containers 'peer0_org1' e 'peer1_org1' e conecta-los à rede healthnet, de acordo com a especificação do arquivo pc2-compose.yaml.</p>

<p>5 - De volta ao Host1, no diretório clonado, abra o arquivo 'config.yaml.</p>

<p>Ao abrir este arquivo, note que nas linhas 208, 214, 238, 244 290, 295, 304 e 309 seus campos estão preenchidos com 'localhost:<numero-porta>'. Estes valores devem ser substituídos de acordo com as instruções das duas linhas superiores de cada uma dessas linhas, substituindo 'localhost' pelo endereço IP do Host2 (mantendo o numero da porta especificada).</p>

<p>6 - Após realizar as alterações do passo anterior, execute o seguinte comando:</p>

```go build```

<p>Este comando irá compilar o arquivo go, baixando as dependências caso necessário. Logo após, execute em um novo terminal o arquivo com o comando:</p>

```./healthapp```

<p>Ao concluir o processo, a aplicação estará disponível através do endereço (http://localhost:3000/).</p>

<p>Esta arquitetura tem como objetivo promover a descentralização dos artefatos da rede Hyperledger Fabric através dos recursos do Docker.</p>

<p>Para <b>finalizar</b> a aplicação, execute os seguintes comando:</p>

<p>1 - No terminal do Host2, pressione as teclas 'ctrl+c' para parar os containers</p>

<p>2 - Após o processo de parada, execute o seguinte comando:</p>

```sudo docker-compose -f pc2-compose.yaml down```

<p>Este comando irá excluir os containers presentes no Host2.</p>

<p>3 - No terminal do Host1 correspondente a execução da aplicação, pressione as teclas 'ctrl+c' para encerrar sua execução.</p> 

<p>4 - No terminal do Host1 correspondente a execução dos containers, pressione as teclas 'ctrl+c' para parar a execução.</p>

<p>5 - Após o processo de parada dos containers, execute o seguinte comando:</p>

```sudo docker-compose -f pc1-compose.yaml down```

<p>Este comando irá excluir os containers presentes no Host1.</p>

