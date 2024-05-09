<!-- markdownlint-disable MD041 MD033 -->

<p align="center">
  <br>
  <img src="https://github.com/kilianc/base-golang-cli/assets/385716/a04a114e-6f17-4160-a768-43291f3639de" alt="cli-name" width="150">
</p>

<p>
  <h1 align="center"><code>cli-name</code></h1>
</p>

<p align="center">
  TBD - Description of this project
  <br><br><br>
</p>

```sh
❯ bin/cli-name

                                          
                                          
                                          
            put ascii art here            
                                          
                                          
                                          
                                           v0.0.1

• starting `cli-name` (description)
• press ctrl+c to stop
````

### Usage

> [!NOTE]
> This repository is a base template, to use it clone it, edit the `BINARY_NAME` value in the `Makefile` and run `make rename-cli`.

### Install

With `go install`

```sh
go install https://github.com/kilianc/cli-name/cmd/cli-name@v0.0.1
```

With `docker`

```sh
docker run --rm -it cli-name:v0.0.1 --version
```

With `curl`

```sh
# change the os (linux or darwin) and arch (amd64 or arm64) based on your machine
curl -O https://https://github.com/kilianc/cli-name/releases/download/v0.0.1/cli-name-darwin-arm64.tar.gz
tar -xzf cli-name-darwin-arm64.tar.gz
./cli-name --version
```

### Usage

> [!NOTE]
> TODO

### Local Development

These are the usual suspects

```sh
make run
make build
make test
make cover
```

After running `make build` the binary available in the `bin/` folder

```sh
bin/cli-name --version
```

### Docker

If you prefer to build and run `cli-name` in a docker container, just on of these commands

````sh
make docker-build
make docker-run
````
