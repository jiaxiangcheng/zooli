# zooli

The **web app** for the project Zooli. This part we're going to develop a platform

## Getting Started

```bash
# clone the project
git clone https://github.com/Qiaorui/zooli.git $GOPATH/src/github.com/Qiaorui/zooli/

cd $GOPATH/src/github.com/Qiaorui/zooli/

bee run
```
Then your app should be able to run in http://localhost:8090/
> Note: if any error happens please check prerequisites first.

### Prerequisites

You need to install [golang](https://golang.org/) and [mysql](https://www.mysql.com/) locally.

Then you can install beego package

```bash
go get github.com/astaxie/beego
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
* https://golang.org/doc/effective_go.html  Effective writing Go and coding style
* http://gopl-zh.b0.upaiyun.com/

### Beego
* https://beego.me/docs/intro/

### Frontend
General
* https://github.com/hacke2/hacke2.github.io/issues/1  front end resources
* https://github.com/hacke2/hacke2.github.io/issues/3  front end resources
* http://coderlmn.github.io/code-standards/ front end coding style

UI
* https://beego.me/docs/mvc/view/tutorial.md  Go template guide
* https://semantic-ui.com/ Semantic UI
* https://www.w3schools.com/ basic html, css, js knowledge

Javascript

* http://bq69.com/blog/articles/script/868/google-javascript-style-guide.html js coding style
* http://tc9011.com/ Very interesting blog about js

### Docker
Web
* https://github.com/wsargent/docker-cheat-sheet/tree/master/zh-cn#docker-cheat-sheet
* https://github.com/widuu/chinese_docker

Book

* https://book.douban.com/subject/27082348/ 自己动手写docker
* https://github.com/yeasy/docker_practice 从入门到实践
* https://book.douban.com/subject/26780404/ 第一本Docker

## Development Methodology

The methodology we are going to follow is Agile-like method. We publish **bugs**, **tasks** and **issues** as cards in [project panel](https://github.com/Qiaorui/zooli/projects/1). Then each member has to pick his task and drag into corresponding state column.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [repository tags](https://github.com/Qiaorui/zooli/tags).