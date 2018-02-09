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

A mysql admin account is required

```sql
CREATE USER 'zooli'@'localhost' IDENTIFIED BY '1234';
GRANT ALL PRIVILEGES ON zooli.* TO 'zooli'@'localhost';
```

## Architecture
A typical [MVC](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller) model.

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


## Development Methodology

The methodology we are going to follow is Agile-like method. We publish **bugs**, **tasks** and **issues** as cards in [project panel](https://github.com/Qiaorui/zooli/projects/1). Then each member has to pick his task and drag into corresponding state column.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [repository tags](https://github.com/Qiaorui/zooli/tags).