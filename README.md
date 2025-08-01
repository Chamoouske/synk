# Synk Synk
Sincronize seus arquivos de forma segura e privada. Sem nuvem, sem servidores, sem concessões.

Synk é uma ferramenta de sincronização de arquivos contínua e peer-to-peer. Ele substitui serviços de nuvem proprietários por uma rede direta e criptografada entre os seus próprios dispositivos. Seus dados são seus, e somente seus.

## ✨ Recursos Principais
- 🔒 Privado por Design: Sem servidor central. Os arquivos são transferidos diretamente entre seus dispositivos com criptografia de ponta a ponta (TLS).

- ⚡ Rápido e Eficiente: Utiliza sincronização a nível de bloco. Apenas as partes modificadas de um arquivo são transferidas, economizando tempo e largura de banda. Na sua rede local, a velocidade é máxima.

- 🌐 Descoberta Automática: Encontra outros dispositivos na sua rede local automaticamente usando mDNS. Sem necessidade de configurar endereços IP.

- 💻 Multiplataforma: Escrito em Go, o Synk roda em Linux, macOS, Windows e, claro, no Termux em Android.

- 📦 Código Aberto: Totalmente transparente e aberto à comunidade.

## ⚙️ Como Funciona
O Synk cria uma malha de conexões diretas entre seus dispositivos. Cada dispositivo roda o Synk, que executa quatro tarefas principais:

- Descobre outros dispositivos na rede.

- Monitora os diretórios compartilhados em busca de alterações.

- Compara o estado dos arquivos com outros peers.

- Transfere apenas os blocos necessários de forma segura para manter tudo sincronizado.

```
+----------------+                       +----------------+
|                |                       |                |
|   Laptop       |<---- Conexão TLS ---->|   Smartphone   |
| (Linux/macOS)  |     (Direta, P2P)     |    (Termux)    |
|                |                       |                |
+----------------+                       +----------------+
```

## 🚀 Começando
Pré-requisitos:

- Ter o Go (versão 1.18+) instalado.

Instalação:
```bash
# Clone o repositório
git clone https://github.com/Chamoouske/synk.git
cd synk

# Compile o projeto
go build .
```

## Uso Básico:

### No primeiro dispositivo (ex: Laptop):

```bash
# Inicialize o Synk e crie a identidade do dispositivo
./synk init

# Anote o Device ID que será exibido. Ex: "k3j4h5g6..."

# Comece a monitorar uma pasta
./synk watch /caminho/para/sua/pasta
```

### No segundo dispositivo (ex: Termux):

```bash
# Inicialize o Synk
./synk init

# Adicione o primeiro dispositivo como um peer confiável
./synk add-peer k3j4h5g6...

# Comece a monitorar a mesma pasta (o caminho pode ser diferente)
./synk watch /data/data/com.termux/files/home/storage/shared/pasta
```

Pronto! Os dois dispositivos agora se encontrarão na rede local e manterão as pastas idênticas.

## 🤝 Contribuindo
Contribuições são muito bem-vindas! Se você tem ideias para novos recursos, melhorias ou encontrou um bug, por favor:

Abra uma Issue para discutir a mudança.

Faça um Fork do projeto e crie uma branch para sua feature (git checkout -b feature/minha-feature).

Faça o commit de suas alterações (git commit -am 'Adiciona nova feature').

Envie para a branch (git push origin feature/minha-feature).

Abra um Pull Request.

## 📜 Licença
Este projeto está sob a licença GNU GENERAL PUBLIC LICENSE. Veja o arquivo [LICENSE](./LICENSE) para mais detalhes.