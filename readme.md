# GoPay

GoPay é um sistema de pagamento simplificado desenvolvido em Go, que permite transferências de dinheiro entre usuários e lojistas.

## Funcionalidades:

- Registro de usuários
- Transferência de dinheiro
- Notificação de pagamento

## Configuração do Banco de Dados:

O GoPay requer um banco de dados MySQL para armazenar informações de usuários e transações. Você pode configurar um banco de dados MySQL localmente ou usar um contêiner Docker com uma imagem pré-configurada.



### Configuração com Docker:

1. Certifique-se de ter o Docker instalado em seu sistema. Você pode baixá-lo em [docker.com](https://www.docker.com/get-started).

2. Clone este repositório para o seu ambiente local.

3. Navegue até o diretório clonado e execute o seguinte comando para construir a imagem Docker com o banco de dados MySQL modelado:

**docker build -t gopay-mysql .**

4. Após a conclusão da construção da imagem, execute o seguinte comando para iniciar um contêiner MySQL:

**docker run -d --name gopay-db -e MYSQL_ROOT_PASSWORD=sua_senha_secreta -p 3306:3306 gopay-mysql**

## Instalação e Execução:

1. Certifique-se de ter Go instalado em seu sistema. Você pode baixá-lo [aqui](https://golang.org/dl/).
2. Clone este repositório para o seu ambiente local.
3. Navegue até o diretório clonado e execute o seguinte comando para iniciar o servidor:

**go run main.go**

4. O servidor será iniciado e estará ouvindo em http://localhost:8080 por padrão.

## API Endpoints:

### Registro de Usuário:

**POST /register**


Payload:

```json
{
  "name": "Nome do Usuário",
  "cpf": "12345678900",
  "email": "usuario@example.com",
  "password": "senha123",
  "is_merchant": false
}
```

### Transferência de Dinheiro:

**POST /transfer**
```json
{
  "value": 100.0,
  "payer": 1,
  "payee": 2
}
```

TODO: fix test