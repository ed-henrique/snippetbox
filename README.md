# Snippetbox

Web app similar to Gist or Pastebin, where you can store, share and create snippets.

## How to Setup Environment

> [!IMPORTANT]
> Install [Go >=1.22.1][1] before proceeding.

1. Clone this repo

```bash
git clone git@github.com:ed-henrique/snippetbox.git
cd snippetbox
```

2. Run the server

```bash
go run ./cmd/web
```

> [!NOTE]
> If needed, use `go run ./cmd/web -help` to check the available flags for the app

[1]: https://go.dev/doc/install
