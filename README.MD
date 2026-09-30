# Sistema CRUD em Go com Avaliação de Desempenho

Aplicação desenvolvida em **Go (Golang)** estruturada em camadas (Clean Architecture / Clean Code), projetada para fins de avaliação de desempenho de operações de leitura e escrita sob diferentes cargas de clientes simultâneos.

## Pré-requisitos
- Go 1.22 ou superior instalado.
- PostgreSQL rodando localmente ou via container Docker.
- Apache JMeter (para execução dos cenários de carga).

## Como Executar

1. **Configurar o Banco de Dados:**
   Crie um banco de dados PostgreSQL chamado `trabalho_db`.

2. **Popular o Dataset (50k registros):**
   ```bash
   go run cmd/seeder/main.go

   Endpoints da API

    POST /produtos — Cria um novo produto (Escrita).

    GET /produtos/{id} — Consulta um produto pelo ID (Leitura).


---

## 4. Respostas Detalhadas das Questões Analíticas

### 1. O que aconteceria com o desempenho da aplicação se o tamanho do dataset fosse reduzido para 20k registros?
* **Comportamento Esperado:** Observar-se-ia um **aumento no *throughput* (requisições por segundo)** e uma **redução na latência média** tanto para leituras quanto para escritas.
* **Justificativa Técnica:** Com um volume menor de dados (20 mil em vez de 50 mil linhas), o volume físico ocupado pelas tabelas e índices no disco diminui proporcionalmente. Isso faz com que o **Buffer Pool (memória RAM) do banco de dados** consiga armazenar uma porcentagem muito maior (ou a totalidade) dos dados mais acessados (*working set*). Consequentemente, evitam-se operações custosas de I/O em disco (*disk reads*), agilizando a recuperação e persistência dos registros.

---

### 2. O que aconteceria se ocorresse a adição de índices no BD? Qual índice você escolheria? O desempenho de consulta melhora?
* **Índice Escolhido:** Para o cenário proposto, o ideal seria a criação de um **índice B-Tree** na chave primária (já gerada nativamente pelo `SERIAL PRIMARY KEY`) ou em colunas de busca frequente (ex: `CREATE INDEX idx_produtos_nome ON produtos(nome);`).
* **Impacto nas Consultas (Leituras):** O desempenho melhora drasticamente. Em vez de o banco de dados realizar uma varredura sequencial completa na tabela (*Sequential Scan* de complexidade $O(N)$), ele passa a navegar pela árvore balanceada do índice (*Index Scan*), reduzindo a complexidade de busca para $O(\log N)$.
* **Trade-off (Escritas):** As operações de escrita (`INSERT`, `UPDATE`) sofrem uma leve sobrecarga (*overhead*). Cada inserção exige que o banco atualize não apenas os blocos de dados da tabela, mas também reorganize e insira chaves na estrutura da árvore do índice. Contudo, para cargas equilibradas ou orientadas à leitura, o ganho nas consultas compensa largamente o custo de escrita.

---

### 3. Como a escolha do BD (SQL vs. NoSQL) impacta no meu sistema?
A escolha entre sistemas relacionais (SQL) e não relacionais (NoSQL) molda diretamente a arquitetura de persistência e os gargalos de desempenho da aplicação:

* **Bancos de Dados Relacionais (SQL - ex: PostgreSQL):**
  * **Impacto:** Fortemente voltados à integridade estrutural e garantias **ACID** (Atomicidade, Consistência, Isolamento e Durabilidade). São ideais para sistemas onde transações financeiras e relacionamentos complexos entre entidades (`JOINs`) são essenciais.
  * **Gargalos de Desempenho:** Sob altíssima concorrência de escritas, o controle de bloqueios (*locks*) e controle de concorrência multiversão (MVCC) pode gerar contenção de recursos, limitando a escalabilidade vertical se o hardware não for dimensionado adequadamente.

* **Bancos de Dados Não Relacionais (NoSQL - ex: MongoDB / Redis):**
  * **Impacto:** Focados em alta escalabilidade horizontal e flexibilidade de esquema (*schema-less*). O armazenamento orientado a documentos (JSON/BSON) ou pares de chave-valor elimina a necessidade de junções complexas.
  * **Gargalos de Desempenho:** Operações de escrita costumam ser extremamente performáticas e escaláveis devido à ausência de restrições relacionais complexas. No entanto, se o sistema exigir buscas com múltiplos filtros cruzados ou consistência estrita imediata em múltiplos nós (*Distributed Transactions*), a complexidade de modelagem e consistência eventual pode impactar o desenvolvimento e a previsibilidade das consultas.