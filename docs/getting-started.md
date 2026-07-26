---
layout: default
title: Getting Started
---

# Getting Started

## Installation

Build from source with Go 1.22 or later:

```sh
git clone https://github.com/reuben-emmens/revisio.git
cd revisio
go build -o revisio ./cmd/revisio
```

Or install directly with `go install`:

```sh
go install github.com/reuben-emmens/revisio/cmd/revisio@latest
```

A `.deb` package is also available from the [releases page](https://github.com/reuben-emmens/revisio/releases).

## Usage

Create a flashcard:

```sh
revisio create --key "capital of France" --value "Paris"
```

Read a flashcard back:

```sh
revisio read --key "capital of France"
```

Check the installed version:

```sh
revisio version
```

See the [Command Reference](commands.html) for the full list of flags.
