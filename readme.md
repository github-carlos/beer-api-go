API de estudo em Go

### Routes
GET http://localhost:4000/beer \
POST http://localhost:4000/beer \
GET http://localhost:4000/beer/1 \
PUT http://localhost:4000/beer/1 \
DELETE http://localhost:4000/beer/1

### db
Para teste foi utilizado o sqlite3

comando para criar o banco
```
sqlite3 data/beer.db
```
CREATE TABLE beer ( id INTEGER PRIMARY KEY AUTOINCREMENT, name text NOT NULL, type integer NOT NULL, style integer NOT NULL );
### Biblioteca para trabalhar com Sqlite
```
go get github.com/mattn/go-sqlite3
```
