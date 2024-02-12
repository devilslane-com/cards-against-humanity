# Cards Against Humanity

Thanks to Celsius at https://www.restagainsthumanity.com (see also https://github.com/celsiusnarhwal/rest-against-humanity)!

This sets up a database for an electronic version of the game. You should set up a MongoDB instance with a database, username, and password.

## Build the importer

```bash
go mod tidy
cd src/import
go build -o ../../../bin/import.exe .
```

## Run the importer 

For dev:

```bash
export 
go run ./main.go "../../resources/data/cards.json" "mongodb://user:pass@localhost:27017/db_name" db_name
```

For prod:

```bash
cd bin/
import "../resources/data/cards.json" "mongodb://user:pass@localhost:27017/db_name" db_name
```