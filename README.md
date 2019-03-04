| Data | Autor | Comentarios | Versao |
| --- | --- | --- | --- |
| 10/01/2019 | Wagner Alves (aka Barão) | Initial Version | 0.0.3 |

# GO Fast AWS Connections API

É uma API de adapters desenvolvida para integração com alguns serviços da AWS [Amazon AWS Cloud](https://aws.amazon.com) (DynamoDB, SQS, SNS, S3, Lambda).

## O que a API fornece:

> Fornece conexões e abstrações para os serviços citados acima, centralizando todo codigo de infraestrutura e facilitando a interação com os recursos e/ou serviços.

## IMPORTANTE:

> As conexões estão configuradas com um provider para pegar o valor diretamente das variáveis de ambiente OU do arquivos 'credentials' e 'config' que fica localizado na pasta: usuario/.aws/ (http://docs.aws.amazon.com/pt_br/cli/latest/userguide/cli-chap-getting-started.html) configurado na máquina).

## Exemplo do arquivo 'credentials' em: usuario/.aws/credentials

```
[company]
aws_access_key_id = AKIAJ2TPF...
aws_secret_access_key = fjGtNv3Jidr9dYzS7...
region = us-east-1

```

## Exemplo do arquivo 'config' em: usuario/.aws/config

```
[profile company-dev]
output = json
role_arn = arn:aws:iam::ACCOUNT_ID:role/IamRoleAdmin
source_profile = company
region = us-east-1

[profile company-hml]
output = json
role_arn = arn:aws:iam::ACCOUNT_ID:role/IamRoleAdmin
source_profile = company
region = us-east-1

[profile company-prd]
output = json
role_arn = arn:aws:iam::ACCOUNT_ID:role/IamRoleAdmin
source_profile = company
region = us-east-1

```
## Ordem de obtenção de credenciais AWS pelo Auth do GO Fast AWS Connections:

 O Modulo 'fac_credentials' do GO Fast Connections obtem credencial AWS na seguinte ordem:

```
 1-Metadados de instacias de EC2.
 2-Propriedades passadas pelo ClientConfiguration 'application.properties'
 3-Variáveis de Ambiente
 4-Roles configuradas por profile nos arquivos .aws/credentials e .aws/config
 5-Static Profile, por profile do arquivo .aws/credentials

```

## Configurar variáveis de ambiente para DEV (OPCIONAL):

Linux, macOS, or Unix

> $ export AWS_ACCESS_KEY_ID=AKIAJ2TPF...

> $ export AWS_SECRET_ACCESS_KEY=fjGtNv3Jidr9dYzS7...

> $ export AWS_REGION=us-west-1

Windows

> set AWS_ACCESS_KEY_ID=AKIAJ2TPF...

> set AWS_SECRET_ACCESS_KEY=fjGtNv3Jidr9dYzS7...

> set AWS_REGION=us-west-1

Links úteis:

[Variáveis de ambiente](http://docs.aws.amazon.com/pt_br/cli/latest/userguide/cli-environment.html)

[Arquivos de Configuração de Credencial](http://docs.aws.amazon.com/pt_br/cli/latest/userguide/cli-config-files.html)


## Exemplos de utilização:

Exemplo AWS SQS:

Inserindo dependencia:
 
``` 
    package awssqs

    import (
        "go-fast-aws-connections/fac_sqs"
    )

```

# SQS

Inicializando o client:

```
    func init() {
        profile := "company-dev"
        facsqs.Start(profile)
    }

```

Method to send messages:

```
    func SendMessage(queueName string, message string) {

        //DelaySeconds: aws.Int64(10),
        result, err := facsqs.SendMessage(queueName, message)

        if err != nil {
            fmt.Printf("Error sending message to queue: %s , %s ", queueName, err)
            return
        }

        fmt.Println("Success", *result.MessageId)
    }

```

# DynamoDB
```

```
# S3
```

```

## Change log:

| Data | Autor | Descrição | Versao |
| --- | --- | --- | --- |
| 10/02/2017 | Wagner Alves (aka Barão) | Initial Version | 1.0.0 | 

## Restrições:

- Testes de integração não são habilitados por default, pois necessitam de criação prévia de recursos na AWS.

## TO DO:

- Cloud Formation para criação de recursos e execução de testes.
- Converter este READ ME para o inglẽs.
- Criação automática de recursos na AWS para testes.

## I hope you enjoy it and feel free to colaborate!!!