# go-dealership-api

REST API implementation to manage cars data in dealership. Implemeted with Go using fiber with MongoDB Database.

## How To Use

1. Clone the repository into your local machine.

2. Copy the `.env.example` file.

```
cp .env.example .env
```

3. Fill the database name and mongo URI in `.env` file.

4. Make sure the database is online then start the application.

```
go run main.go
```

5. For docker user, fill the `VOLUME` field in `.env` file with the correct directory.

6. Build the application.

```
docker compose build
```

7. Run the application.

```
docker compose up
```