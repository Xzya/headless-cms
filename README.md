# Headless CMS

This is an attempt to make a Headless CMS in `go`.

## API documentation

You can find the API documentation [here](http://rebilly.github.io/ReDoc/?url=https://raw.githubusercontent.com/Xzya/headless-cms/master/swagger/swagger.yml).

## Requirements

- Go 1.10
- [dep](https://github.com/golang/dep) for installing the dependencies
- [go-swagger](https://github.com/go-swagger/go-swagger) for generating the server stubs
- A database, e.g. PostgreSQL

## Running the project

First, generate the server stub from the swagger definition

```bash
./gen.sh
```

Install the dependencies using dep

```bash
dep ensure
```

Build the executable

```bash
./build.sh
```

Run the binary

```bash
./main
```

Run `--help` to check the available parameters

```bash
./main --help
Usage of main:
  -database string
        The database name (default "cms")
  -db-host string
        The database hostname (default "localhost")
  -db-port string
        The database port (default "5432")
  -dialect string
        The database dialect (default "postgres")
  -pass string
        The database password (default "test")
  -port int
        Port to run this service on (default 3000)
  -user string
        The database username (default "test")
```

## Dashboard

You can find a dashboard for the project [here](https://github.com/Xzya/headless-cms-dashboard).

## Built using

- [Swagger](https://swagger.io/), API definition
- [go-swagger](https://github.com/go-swagger/go-swagger), generate server stub based on swagger definition
- [gorm](https://github.com/jinzhu/gorm), database operations

## License

Open sourced under the [MIT license](./LICENSE.md).