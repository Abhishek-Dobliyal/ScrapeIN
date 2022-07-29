# ScrapeIN
A simple CLI tool written in GoLang to scrape jobs from Indeed.


### Installations

- [GoLang](https://go.dev/)

### Dependencies

- Navigate to the project directory and initialise a `go.mod` file using

```bash
go mod init
```

- To install the dependencies, run

```bash
go mod tidy
```

### Usage

```bash
go run main.go --tag "software engineer intern" --name "output.json"
```

### To-Do

- [ ] Add notifier 
- [ ] Add advanced flag searches and filters
