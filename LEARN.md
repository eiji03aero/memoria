# Overall
## General
### Dev productivity is pretty low
- it takes 1-2 hours to implement api with test
    - creating boilerplate files
    - not intuitive error message on test
- it takes 3-4 hours to complete one screen on frontend
    - need to get used to tools
    - storybook needs to be restarted pretty often
    - nextjs taking fluffying minutes to reflect changes

# RDB
## SQL
### to merge two datasets
- https://qiita.com/tarosuke777000/items/391b0291faae45974be1
- use union or union all

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
- fargate
    - use fargate spote
    - this time it reduced (per day):
        - before: $0.87
        - after: $0.15
        - ratio: down to 17.2%, cutting 82,8%
- rds
    - choose cheap model (this time t3 micro)
    - lower allocated storage
    - use magnetic storage
    - seems to be we can run rds on ec2 spot instances
        - https://reintech.io/blog/managing-costs-aws-rds-reserved-spot-instances
- ec2
    - stop a basion server if you are not using
- vpc
    - for now the usage is on:
        - public ip ($0.005 / hour)
        - bastion, client, api, db-migration, alb?
- alb
    - use only one alb to handle multiple domains
- eventbridge scheduler
    - use this to schedule certain operations
    - for this time, I used it to
        - stop resources at 0:00 (rds, ecs services)
        - start resources at 18:00 (rds, ecs services)

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

## S3
### 400 for uploading with presigned url
- https://scrapbox.io/kidaaam-92022284/Presigned_URL%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6S3%E3%81%AB%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%92put%E3%81%99%E3%82%8B%E3%81%A8%E3%81%8D%E3%81%AB400%E3%82%A8%E3%83%A9%E3%83%BC%E3%81%AB%E3%81%AA%E3%82%8B%E4%BB%B6
- Authentication header was getting in the way

### Key name
- https://docs.aws.amazon.com/AmazonS3/latest/userguide/upload-objects.html
- it should not contain slach at the beginning of key

### Access denied when accessed from browser
- https://karukichi-blog.netlify.app/blogs/resolve-403-error-in-s3

---

# HTTP
## Status code
### Which code to use for redirects?
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections
- for post, 303 see other should be preferable


---

# Go
## Basics
### nested struct will be initialized with zero-values
- Thought it become nil, but not
```
type Parent struct {
	Child struct {
		Value string
	}
}

func main() {
	d := &Parent{}
	// Can set value, it is initialized with zero-values
	d.Child.Value = "Hoge child"
}
```

### zero-value slice
```
type Parent struct {
	Slice []string
}

func main() {
	d := &Parent{}
	log.Println(d.Slice)    // []
	log.Println(d.Slice[0]) // panics saying index out of range
}
```

### declaration of variables in nested that has same name as return variable
- will be treated as new variable
    - well it is surrouned by block so :(
```

func NestedReturnVar() (user error, str string) {
	str = "base"

	if true {
		user, str := nil, "from child func
		return
	}
	return
}

func main() {
	str := NestedReturnVar()
	log.Println(str)
}
```

### To define custom error
- Error method for error interface
- Is method for errors.Is
```

type Internal struct {
	message string
}

func NewInternal(message string) error {
	return &Internal{message: message}
}

func (e Internal) Error() string {
	return fmt.Sprintf("internal error: %s", e.message)
}

func (e *Internal) Is(err error) bool {
	other, ok := err.(*Internal)
	if !ok {
		return false
	}

	return e.message == other.message
}
```

### Get stuct name from interface value
```
if t := reflect.TypeOf(err); t.Kind() == reflect.Ptr {
    log.Println("pointer name:", t.Elem().Name())
} else {
    log.Println("value name:", t.Name())
}
```

### using errors.Is
- whether value or pointer affects the result
```
type CError struct{}

func (e CError) Error() string {
	return "let's go"
}

func main() {
	err := CError{}
	log.Println(errors.Is(err, CError{})) // true

	err2 := CError{}
	log.Println(errors.Is(err2, &CError{})) // false

	err3 := &CError{}
	log.Println(errors.Is(err3, CError{})) // false

	err4 := &CError{}
	log.Println(errors.Is(err4, &CError{})) // true
}
```

- seems to be the return value of Error method needs to be same to be true for errors.Is comparison
- seems to be:
    - errors.Is is pretty much for constant error the ones that are static and do not have dynamic properties
    - to cover above case, you should use errors.As
```
type CError struct{ message string }

func (e CError) Error() string {
	return "let's go " + e.message
}

func main() {
	err := CError{}
	log.Println(errors.Is(err, CError{message: "1"})) // false

	err2 := CError{}
	log.Println(errors.Is(err2, CError{})) // true

	err3 := CError{}
	log.Println(errors.As(err3, &CError{message: "3"})) // true
}
```

### Make new slice with capacity
- 2nd arg is for length
- 3rd arg is for capacity
```
userSpaces = make([]*model.UserSpace, 0, len(userSpaceTbls))
```

### Zero-value for func is nil
```
func main() {
	var d func()
	log.Println(d)
}
```

### nil pointer deference panics
```
func main() {
	var d *bool
	log.Println(*d) // panics with invalid memory address or nil pointer dereference
}
```

### Buffered channel
- In this example, it won't block the main thread as we have go-ed 3 funcs
```

// You can edit this code!
// Click here and start typing.
package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

type ch chan int
type rch chan string

func work(wg *sync.WaitGroup, c ch, rc rch) {
	wg.Add(1)
	defer wg.Done()
	log.Println("going to for in work")
	for i := range c {
		log.Println("received i in for work")
		time.Sleep(time.Duration(i) * time.Second)
		log.Println("hogeee: " + strconv.Itoa(i))
	}

	rc <- "done working"
}

func main() {
	wg := &sync.WaitGroup{}
	c := make(ch, 1)
	rc := make(rch)
	go work(wg, c, rc)
	go work(wg, c, rc)
	go work(wg, c, rc)

	log.Println("going to work")
	c <- 1
	c <- 2
	c <- 3
	close(c)
	wg.Wait()

	for r := range rc {
		log.Println("doneee: " + r)
	}
	log.Println("done")
}
```

---

# Frontend
## next.js
### How to import external css with layer specified
- https://zenn.dev/swallowtail62/articles/4d81a75a6be2e7
- use postcss's plugin

## tanstack-query
### Difficulty with handling error response body
- want to run logic with onSuccess
- error response body can be obtained only when mutateFn is run successfully
- this leads to running onSuccess thought request was an error

## Motherfxcking react-spectrum
### where is typography component?
- it has been discussed for years

### Just gave up utilizing this as design sytem
- functionality is very limited (should be as intended though)
- size attributes are very not intuitive
- cannot reuse color, size in other place
- thanks for wasting my time
- just gonna go with panda and use this as base ui component library

### who the fluff uses content-box ... ?
- why, just why

## pandacss
### so far so good
- css in js supporting rsc
- the css function needs to be imported via direct path to styled-system/css
    - you cannot reexport, the stylesheet would not be attached

### Difficulty with using dynaic values
- especially specifying size like height
- cannot really get work-arounded yet

---

# Makefile
## Difficulty with exntending
- cannot pass arguments like sh file can
