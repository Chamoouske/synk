**Documento de Arquitetura e Design (ADD):** Synk

**Versão:** 1.0

**Data:** 31 de Julho de 2025

# 1. Introdução
Este documento detalha a arquitetura técnica do Synk, um sistema de sincronização de arquivos peer-to-peer. Ele serve como um guia para os desenvolvedores, descrevendo os componentes principais, protocolos, estruturas de dados e o modelo de concorrência.

# 2. Componentes Principais
O sistema é composto por quatro módulos concorrentes principais, implementados como goroutines em Go.

## 2.1. Módulo de Descoberta (discovery)

Tecnologia: mDNS (Multicast DNS) via biblioteca grandcat/zeroconf.

Operação:

Registro: Ao iniciar, o Synk registra um serviço na LAN.

Nome do Serviço: _synk._tcp

Domínio: local.

Porta: A porta TCP na qual o módulo de comunicação está ouvindo.

Metadados (TXT Record): Contém o DeviceID único do peer.

Busca: Simultaneamente, o módulo procura por outros serviços _synk._tcp na rede.

Saída: Quando um novo peer é encontrado, seu DeviceID, endereço IP e porta são passados para o módulo de comunicação para estabelecer uma conexão.

## 2.2. Módulo de Monitoramento de Arquivos (watcher)

Tecnologia: Biblioteca fsnotify/fsnotify.

Operação:

O módulo estabelece um "watch" nos diretórios configurados pelo usuário.

Ele ouve eventos do sistema operacional: CREATE, WRITE, REMOVE, RENAME.

Para evitar picos de processamento (ex: salvar um arquivo grande gera múltiplos eventos WRITE), os eventos são agrupados (debounced) em um curto intervalo de tempo (ex: 500ms).

Após o debounce, uma única notificação de "mudança de arquivo" é enviada para o Módulo de Sincronização.

## 2.3. Módulo de Sincronização (synchronizer)
Este é o núcleo lógico do Synk. Ele mantém o estado do diretório e executa o protocolo de sincronização.

Estrutura de Dados: O Índice (Index)

Cada peer mantém um índice em memória do estado de cada diretório compartilhado.

type Index map[string]FileInfo (mapeia o caminho relativo do arquivo para suas informações).

Estrutura de Dados: FileInfo
```go
type FileInfo struct {
    ModTime int64       // Timestamp da última modificação
    Size    int64       // Tamanho total do arquivo
    Blocks  []BlockInfo // Lista de hashes dos blocos
}
```

Estrutura de Dados: BlockInfo
```go
type BlockInfo struct {
    Hash string // Hash SHA-256 do bloco
}
```

O tamanho do bloco é fixo (ex: 1MB). O último bloco pode ser menor.

Fluxo do Protocolo de Sincronização:

Mudança Local: O watcher notifica o synchronizer sobre uma mudança no arquivo.txt.

Atualização do Índice Local: O synchronizer lê o arquivo.txt, divide-o em blocos, calcula o hash SHA-256 de cada bloco e atualiza sua própria estrutura Index.

Anúncio: O peer A envia uma mensagem AnnounceUpdate para todos os seus peers conectados.

Requisição de Índice: O peer B, ao receber o anúncio, envia uma mensagem RequestIndex para o peer A.

Comparação: O peer A responde com seu Index atualizado. O peer B o compara com seu próprio índice para o arquivo.txt. Ele identifica quais hashes de bloco são diferentes ou novos.

Requisição de Blocos: O peer B envia uma mensagem RequestBlocks para o peer A, especificando os hashes dos blocos que ele precisa.

Transferência de Blocos: O peer A lê apenas os blocos correspondentes do arquivo.txt no disco e os envia para o peer B.

Reconstrução: O peer B recebe os blocos, escreve-os na posição correta em sua cópia local do arquivo.txt, e finalmente atualiza seu próprio Index. A sincronização está completa.

## 2.4. Módulo de Comunicação (transport)

Tecnologia: Conexões TCP sobre TLS. Uso de crypto/tls e crypto/x509 para gerar certificados autoassinados.

Identidade: Na primeira inicialização, um par de chaves RSA e um certificado x509 autoassinado são gerados. O DeviceID é um hash da chave pública para facilitar a leitura.

Handshake TLS:

Quando o Peer A se conecta ao Peer B, eles realizam um handshake TLS padrão.

Crucialmente, cada peer verifica se o DeviceID do certificado apresentado pelo outro corresponde a um DeviceID na sua lista de peers confiáveis. Se não corresponder, a conexão é encerrada imediatamente. Isso previne ataques man-in-the-middle.

Protocolo de Mensagens:

Para garantir uma comunicação estruturada, as mensagens trocadas serão definidas usando Protocol Buffers (Protobuf) ou gRPC. Isso cria um contrato de API forte entre os peers.

Mensagens Principais:

AnnounceUpdate{}

RequestIndex{}

IndexMessage{...}

RequestBlocks{Hashes: []string}

BlockData{Hash: string, Data: []byte}

# 3. Modelo de Concorrência
O programa principal inicia os quatro módulos como goroutines separadas.

Canais (Channels) são usados para a comunicação segura entre as goroutines (ex: o watcher envia eventos para o synchronizer através de um canal).

O acesso a estruturas de dados compartilhadas, como o Index de arquivos e a lista de Peers conectados, deve ser protegido por Mutexes (sync.Mutex) para evitar race conditions.

Cada conexão de peer aceita será gerenciada em sua própria goroutine para lidar com múltiplos peers simultaneamente.