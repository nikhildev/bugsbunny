# BugsBunny

Bug tracking and management application.

## How to Run

### Run without building

```bash
go run ./api run server
```

```bash
go run ./api migrate
```

### Build first, then run

```bash
go build -o bugsbunny ./api
./bugsbunny run server
./bugsbunny migrate
```
