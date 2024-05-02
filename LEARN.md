# AWS
## General
### aws icons
- https://aws-icons.com

## Cost
### To be extreme stingy
- ecs
    - lower cpu
    - lower memory
    - limit the number of containers (task)
- rds
    - choose cheap model (this time t3 micro)
    - lower allocated storage
    - use magnetic storage
- alb
    - use only one alb to handle multiple domains

## VPC
### what is ip v4 cidr blocks?
    - https://docs.aws.amazon.com/vpc/latest/userguide/vpc-cidr-blocks.html
    - this is to specify certain range of ip v4 address

### ALB vs ELB
- related:
    - https://www.logicmonitor.com/blog/alb-vs-elb#:~:text=One%20of%20the%20most%20significant,number%2C%20hostname%2C%20and%20path.
- Basically:
    - ALB has more functionalites like routing based on url components
    - ELB is more simple

### why the fluff i cannot open public ips?
- it was the security group
- the default one does not allow inbound http request

### ssh to ec2
- https://qiita.com/takuma-jpn/items/b2c04b7a271a4472a900

## ECS
### resource initialization error on ecs task
- should be pretty much about not being able to http to internet
- easiest way is assigning public ip

### ecs deploy fails with api
```
ResourceInitializationError: unable to pull secrets or registry auth: execution resource retrieval failed: unable to retrieve ecr registry auth: service call has been retried 3 time(s)
```
- related:
    - https://repost.aws/questions/QUTTXzLlU_T72po4QROT1c2w/resourceinitializationerror-unable-to-pull-secrets-or-registry-auth-execution-resource-retrieval-failed-unable-to-retrieve-ecr-registry-auth-service-call-has-been-retried-3-time-s-requesterror
- just building infrastructure without any preceeding experience is too hard to handle
- will just go with setup wizard of vpc to create related resources (igw, subnets)

## Route53
### configure domain linked to ecs containers
- related:
    - https://zenn.dev/ttani/articles/aws-ecs-hostbase-routing
    - https://qiita.com/sugimount-a/items/c3dd0c177461d6b5131b
    - https://zenn.dev/taiki_asakawa/books/dfc00287d5b8c7/viewer/a9d77b

### Create sub domain of domain bought in route 53
- https://repost.aws/ja/knowledge-center/create-subdomain-route-53

## SDK
### go sdk
- https://github.com/aws/aws-sdk-go-v2
- https://aws.github.io/aws-sdk-go-v2/docs/getting-started/

## WorkMail
### Getting started
- https://qiita.com/ysKey2/items/2b019337772f8499beec

# HTTP
## Status code
### Which code to use for redirects?
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections
- for post, 303 see other should be preferable
