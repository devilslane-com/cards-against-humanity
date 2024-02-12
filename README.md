# Cards Against Humanity

Thanks to Celsius at https://www.restagainsthumanity.com (see also https://github.com/celsiusnarhwal/rest-against-humanity)!

This sets up a database for an electronic version of the game. You should set up a MongoDB instance with a database, username, and password.

## Importer

THis will create the `packs` and `cards` collections.

### Build the importer

```bash
go mod tidy
cd src/import
go build -o ../../../bin/import.exe .
```

### Run the importer 

```bash
cd bin/
import "../resources/data/cards.json" "mongodb://user:pass@localhost:27017/db_name" db_name
```

## Seeder

This will create the `games`, `players`, `rounds`, and `responses` collections, and an example structure.

### Build the seeder

```bash
go mod tidy
cd src/seed
go build -o ../../../bin/seed.exe .
```

### Run the seeder 

```bash
cd bin/
seed "../resources/data/cards.json" "mongodb://user:pass@localhost:27017/db_name" db_name
```