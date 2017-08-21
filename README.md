# Wreck REST Request Tool

A simple command line tool for interacting with REST APIs
---

## Building and Installing

In order to build this tool you should have it checked out locally into your GOPATH.  Once it is checked out you should be able to compile and install wreck by using:

```bash
$ make install
```

To install the configuration files you will need to copy them to the locations where wreck will look for them.

```bash
$ cp global.yaml /etc/wreck.yaml
$ cp user.yaml ~/.wreck.yaml
```

## Running Tests

```bash
$ make test
```

## Usage

Wreck can make both GET requests and POST requests to an API.  You can also pass username and password credentials on the command line or in the user-specific configuration file located at `~/.wreck.yaml`.

### Example Usage

#### Getting Help

```bash
$ wreck help
```

#### Making a GET Request
```bash
$ wreck get https://httpbin.org/get
```

#### Making a POST Request with Credentials
```bash
$ wreck post --username foo --password bar https://httpbin.org/post
```