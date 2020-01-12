# gopi-app

This repository contains a template gopi application. You should
clone it and then use it as a basis for your own applications:

```bash
bash% git clone https://github.com/djthorpe/gopi-app helloworld
bash% cd helloworld
bash% git remote remove origin
```

The repository contains a `Makefile` which you can use to compile
the application. The targets are:

  * `make test` will run the application tests;
  * `make linux` will install the application for Linux;
  * `make darwin` will install the application for Darwin (MacOS);
  * `make rpi` will install the application for Raspberry Pi;
  * `make clean` will run `go clean`.

Without any modification a command `helloworld` is compiled and installed.
You need to edit the files under `cmd/helloworld` in order to start
developing your application.


