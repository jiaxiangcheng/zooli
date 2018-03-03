# Zooli 酥梨

The **web app** for the project Zooli. This part we're going to develop a platform

## Getting Started

```bash
# clone the project
git clone https://github.com/Qiaorui/zooli.git $GOPATH/src/github.com/Qiaorui/zooli/

cd $GOPATH/src/github.com/Qiaorui/zooli/

# update package
go get -u

# run web app
bee run
```
Then your app should be able to run in http://localhost:8090/

The default account:
```
username: admin
password: 1234
```

If you generate random data set, all fake account has password: 111111

> Note: if any error happens please check prerequisites first.

### Prerequisites

You need to install [golang](https://golang.org/) and [mysql](https://www.mysql.com/) locally.

Then you can install beego package

```bash
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
```

A mysql admin account is required. You should enter mysql shell and then use following command to create a mysql user for this project.

```sql
CREATE USER 'zooli'@'localhost' IDENTIFIED BY '1234';
GRANT ALL PRIVILEGES ON zooli.* TO 'zooli'@'localhost';
```

## Architecture

In this application we are using **beego** framework + **semantic ui**.
Beego is typical [MVC](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller) model. So you understand why we have controller, model, view directories .

```
.
├── conf/                       Config files for different environments
├── controllers/                Controllers to handle all http requests
├── models/                     Database and model definition
├── routers/                    Routers and filters
├── tests/                      For now we have enough time do this, just ignore it
├── static                      All static resource should store in here, a example would be organize css, js, etc by directory
│   ├── css/
│   ├── js/
│   └── ...
├── views                       Layout, template, html files
├── main.go                     Main entry point
└── semantic.json               Semantic UI config file
```

## Tutorials
Here I'm going to give some useful link to study all required knowledge

### Golang
Basic
* https://tour.golang.org/welcome/1 Golang tutorial

Advance
* http://gorm.io/  GORM guide
* https://github.com/asaskevich/govalidator  GO validator
* https://golang.org/doc/effective_go.html  Effective writing Go and coding style
* http://gopl-zh.b0.upaiyun.com/

### Beego
* https://beego.me/docs/intro/
* https://drive.google.com/open?id=172OSzfEmixDe3zIsZ4b7umJSD9YEjcWX beego-in-action
* https://lei-cao.gitbooks.io/beego-in-action/content/zh/index.html Beego 实战

### Frontend
General
* https://github.com/hacke2/hacke2.github.io/issues/1  front end resources
* https://github.com/hacke2/hacke2.github.io/issues/3  front end resources
* http://coderlmn.github.io/code-standards/ front end coding style
* http://codeguide.bootcss.com/  html css code style
* http://alloyteam.github.io/CodeGuide/#html  yes, code style again

UI
* https://beego.me/docs/mvc/view/tutorial.md  Go template guide
* https://semantic-ui.com/ Semantic UI
* https://www.w3schools.com/ basic html, css, js knowledge
* https://facebook.github.io/flux/ FLUX

Javascript

* http://bq69.com/blog/articles/script/868/google-javascript-style-guide.html js coding style
* http://tc9011.com/ Very interesting blog about js

Jquery

* https://jquery.com/
* https://www.w3schools.com/jquery/default.asp

### Backend
* http://www.cnblogs.com/samlin/archive/2010/02/08/log-operation-management.html Logging system
* http://www.cnblogs.com/hooray/archive/2012/09/05/2672133.html Logging system
* http://blog.csdn.net/jackljf/article/details/5750577 Logging system


### Docker
Training
* https://www.katacoda.com/courses/docker
* https://training.docker.com/
* https://hackr.io/tutorials/learn-docker

Web
* https://github.com/wsargent/docker-cheat-sheet/tree/master/zh-cn#docker-cheat-sheet
* https://github.com/widuu/chinese_docker

Book

* https://book.douban.com/subject/27082348/ 自己动手写docker
* https://github.com/yeasy/docker_practice 从入门到实践
* https://book.douban.com/subject/26780404/ 第一本Docker

## Use cases

A company like BWM has a lot of stores. Each store can provide some kind of services (repair, wash, etc)
and some products that belong to one service. The client can order a product and pay.
Each store will have some manager (user) who using this application to monitoring the workflow and receive money.
For now we are putting all kinds of user (admin, manager) in the same User table, differ by Role. This is easiest way to do so, but not correct.

![alt text](https://image.ibb.co/dZ2ckx/network_1.png "class diagram")

Following we define some use cases. The necessary fields to create a model are defined in Go source code file under **/model** directory.

| User    | Action          | Pre          | Summary                                       | Post                    |
| ------- |:---------------:| :-----------:| --------------------------------------------- | :---------------------: |
| All     | Login           | Not login    | User enters username and password             | Login and save session cookie if correct        |
| All     | Logout          | Login        | Click logout                                  | Delete user session cookie |
| Admin   | Show company list |            | Show all companies                            |                         |
| Admin   | CRUD company    | Company list | Create, Read, Update or Delete a company      | Check and do CRUD       |
| Admin   | Show store list |              | Show all stores                               |                         |
| Admin   | CRUD store      | Store list   | CRUD a store                                  | Check and do CRUD       |
| Admin   | Show service list |            | Show all services                             |                         |
| Admin   | CRUD service    | Service list | CRUD a service                                | Check and do CRUD |
| Admin   | Show user list  |              | Show all users                                |                         |
| Admin   | CRUD user       | User list    | Create, Read, Update or Delete a user         | Check and do CRUD |
| User    | Show order list |              | Show all orders made by corresponding store   |                         |
| User    | RU order        | Order list   | Read detail of a order and/or update          |                         |
| User    | Update store    | Store has a user | User update his store information         |                         |
| User    | Show product list |            | Show all product of his store                 |                         |
| User    | CRUD product    | Product list | CRUD a product                                | Check and do CRUD |


## Development Methodology

The methodology we are going to follow is Agile-like method. We publish **bugs**, **tasks** and **issues** as cards in [project panel](https://github.com/Qiaorui/zooli/projects/1). Then each member has to pick his task and drag into corresponding state column.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [repository tags](https://github.com/Qiaorui/zooli/tags).