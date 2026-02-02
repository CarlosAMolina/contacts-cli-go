# contacts-cli-go

## Introduction

Contacts CLI. Go version.

## Configuration

### VPS configuration

First, compile the binary:

```bash
make build
```

Send the binary to your VPS:

```bash
make send
```

You can create an alias in your VPS `~/.bashrc`:

```bash
# Contacts
alias c='$HOME/Software/contacts-cli-go /tmp/contacts.json'
```

Run:

```bash
c
```

## Run locally

```bash
make run
```
