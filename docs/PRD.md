**Documento de Requisitos do Produto (PRD):** Synk

**Autor:** Chamoouske

**Versão:** 1.0

**Data:** 31 de Julho de 2025

# 1. Visão e Oportunidade
## 1.1. Problema
No cenário digital atual, usuários dependem de serviços centralizados (como Google Drive, Dropbox) para sincronizar arquivos entre dispositivos. Este modelo apresenta três desvantagens principais:

- Privacidade: Os arquivos são armazenados em servidores de terceiros, criando um ponto central de falha e vulnerabilidade a acessos não autorizados.

- Custo e Limites: Serviços gratuitos possuem limites de armazenamento, e planos pagos representam um custo recorrente.

- Eficiência: A sincronização entre dispositivos na mesma rede local é limitada pela velocidade da conexão com a internet, em vez de usar a velocidade máxima da LAN.

## 1.2. Solução Proposta: Synk
Synk é uma ferramenta de sincronização de arquivos de código aberto, peer-to-peer (P2P) e segura. Ela cria uma rede privada e descentralizada entre os dispositivos de um usuário, permitindo que diretórios sejam sincronizados diretamente, sem intermediários. Os dados trafegam apenas entre os dispositivos autorizados pelo usuário, criptografados de ponta a ponta.

## 1.3. Público-Alvo

- Usuários Preocupados com a Privacidade: Indivíduos que desejam controle total sobre seus dados e não confiam em soluções de nuvem de terceiros.

- Desenvolvedores e Entusiastas de Tecnologia: Pessoas confortáveis com a linha de comando (especialmente no Termux) que procuram uma solução de sincronização eficiente e personalizável.

- Equipes Pequenas: Grupos que trabalham na mesma rede local e precisam de uma forma rápida e privada de compartilhar e sincronizar arquivos de projeto.

# 2. Recursos e Requisitos
## 2.1. Requisitos do Produto (MVP - Produto Mínimo Viável)

| ID | Recurso | Descrição Detalhada | Prioridade |
|---|---|---|---|
| MVP-01 | Identidade Segura do Dispositivo | Na primeira execução, o Synk gera um ID de dispositivo único baseado em um par de chaves criptográficas (pública/privada). Este ID é usado para autenticação na rede. | P1 |
| MVP-02 | Descoberta de Peers em Rede Local (LAN) | Dispositivos na mesma rede Wi-Fi devem se descobrir automaticamente usando o protocolo mDNS (Zeroconf), sem necessidade de configuração manual de IPs. | P1 |
| MVP-03 | Pareamento de Dispositivos | Um novo dispositivo só pode entrar na rede de sincronização se for explicitamente autorizado por um dispositivo existente, através da troca de seus IDs. | P1 |
| MVP-04 | Monitoramento de Diretórios em Tempo Real | O Synk deve monitorar os diretórios configurados e detectar criações, modificações e exclusões de arquivos instantaneamente, usando eventos do sistema de arquivos (fsnotify). | P1 |
| MVP-05 | Sincronização por Blocos | Quando um arquivo é modificado, apenas os blocos alterados (partes do arquivo) são transferidos pela rede, não o arquivo inteiro. Isso otimiza a largura de banda e a velocidade. | P1 |
| MVP-06 | Comunicação Criptografada (TLS) | Toda a comunicação entre os peers (metadados e blocos de arquivos) deve ser obrigatoriamente criptografada usando TLS, com cada peer validando o certificado do outro. | P1 |
| MVP-07 | Interface de Linha de Comando (CLI) | A interação inicial será via CLI, com comandos para inicializar, adicionar peers, e especificar diretórios para sincronização. | P2 |

## 2.2. Recursos Futuros (Pós-MVP)

- Resolução de Conflitos: Implementar uma estratégia para lidar com arquivos modificados em múltiplos locais simultaneamente (ex: renomear uma versão como arquivo.conflict.txt).

- Descoberta via WAN: Introduzir um servidor de descoberta opcional para conectar peers em redes diferentes.

- Interface Gráfica (GUI): Criar uma interface gráfica simples para facilitar a configuração e o monitoramento.

- Otimização de Bateria: Modos de baixo consumo para dispositivos móveis (ex: pausar a sincronização quando em redes móveis ou com bateria baixa).

# 3. Métricas de Sucesso
- Adoção: Número de downloads e estrelas no repositório do projeto.

- Confiabilidade: A sincronização ocorre com sucesso em 99.9% dos casos, sem corrupção de dados.

- Eficiência: O uso de CPU e memória em estado ocioso (apenas monitorando) é inferior a 5%.

- Performance: A velocidade de sincronização em uma rede local gigabit supera significativamente a de serviços baseados em nuvem para arquivos grandes.