# GO Fast AWS Connections API
| Data | Autor | Comentarios | Versao |
| --- | --- | --- | --- |
| 10/01/2019 | Wagner Alves (aka Barão) | Initial Version | 0.0.3 |

É uma API de adapters desenvolvida para integração com alguns serviços da [Amazon Web Services](https://aws.amazon.com), DynamoDB, SQS, SNS, S3, Lambda) para aplicações feitas em [GO/GOLANG](https://golang.org/)

## O que a API fornece:

> Fornece conexões e abstrações para os serviços citados acima, centralizando todo codigo de infraestrutura e facilitando a interação com os recursos e/ou serviços pois a api automatiza a obtenção de credenciais por meio de uma cadeia (environment, shared credentias, ec2 metadata, etc...).

## IMPORTANTE:

> As conexões estão configuradas com um provider para pegar o valor diretamente das variáveis de ambiente OU do arquivos 'credentials' e 'config' que fica localizado na pasta: usuario/.aws/

(http://docs.aws.amazon.com/pt_br/cli/latest/userguide/cli-chap-getting-started.html) configurado na máquina).

## Exemplo do arquivo 'credentials' em: usuario/.aws/credentials

```
[company]
aws_access_key_id = AKIAJ2TPF...
aws_secret_access_key = fjGtNv3Jidr9dYzS7...
region = us-east-1

```

## Exemplo do arquivo 'config' em: usuario/.aws/config
> para interação em várias contas por meio do cross account.

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
role_arn = arn:aws:iam::ACCOUNT_ID:role/IamRoleRO
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


# SQS - How to use:

Inserting dependencies:
 
``` 
    package awssqs

    import (
        "go-fast-aws-connections/fac_sqs"
    )

```

Starting the client:

```
    func init() {
        profile := "company-dev"
        facsqs.Start(profile)
    }

```

Method to send messages to SQS Queue:

```
    func SendMessage(queueName string, message string) {

        //DelaySeconds: aws.Int64(10),
        result, err := facsqs.SendMessage(queueName, message)

        if err != nil {
            golog.Infof("Error sending message to queue: %s , %s ", queueName, err)
            return
        }

        fmt.Println("Success", *result.MessageId)
    }

```

# DynamoDB - How to use:
```

```
# S3 - How to use:
```

```

## Change log:

| Data | Autor | Descrição | Versao |
| --- | --- | --- | --- |
| 03/01/2019 | Wagner Alves (aka Barão) | Initial Version | 0.0.1 | 

## Restrições:

- Testes de integração não são habilitados por default, pois necessitam de criação prévia de recursos na AWS.

## TO DO:

- Cloud Formation para criação de recursos e execução de testes.
- Converter este READ ME para o inglẽs.
- Criação automática de recursos na AWS para testes.

## I hope you enjoy it and feel free to colaborate!!!
