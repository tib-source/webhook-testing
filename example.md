# Markdown CLI

A simple command-line tool that converts **Markdown** to **HTML**.

## Features

- Converts markdown to clean HTML
- Extracts the document title
- Counts words and estimates reading time

## Usage

```bash
markdowncli example.md
cat example.md | markdowncli -
```

## Why Go?

Go compiles to a single binary with no runtime dependencies. The module
cache is small (a few MB), making it ideal for CI/CD pipelines with
limited cache storage.
