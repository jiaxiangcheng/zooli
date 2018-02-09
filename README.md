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

### Installing

A step by step series of examples that tell you have to get a development env running

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [repository tags](https://github.com/Qiaorui/zooli/tags).

## Authors

* **项翘睿** - *Initial work* -
