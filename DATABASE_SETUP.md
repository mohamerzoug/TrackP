# PostgreSQL Setup for TrackP

## Installation

### Windows
1. Download PostgreSQL from https://www.postgresql.org/download/windows/
2. Run the installer and follow the setup wizard
3. Remember the password you set for the 'postgres' user
4. Default port is 5432

### Alternative: Using Docker
```bash
docker run --name trackp-postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:15
```

## Database Setup

1. **Connect to PostgreSQL**
   ```bash
   psql -U postgres -h localhost -p 5432
   ```

2. **Create the TrackP database**
   ```sql
   CREATE DATABASE trackp;
   \q
   ```

3. **Run the setup script**
   ```bash
   psql -U postgres -d trackp -f database/setup.sql
   ```

## Configuration

Update database credentials in `backend/main.go` or set environment variables:

```go
// Default values in code:
dbHost := "localhost"
dbPort := "5432" 
dbUser := "postgres"
dbPassword := "040601"  // Change this!
dbName := "trackp"
```

## Alternative: In-Memory Development Mode

If you don't want to set up PostgreSQL immediately, you can modify the backend to use SQLite for development:

1. Install SQLite driver: `go get github.com/mattn/go-sqlite3`
2. Modify the connection string in `main.go`

## Verification

Test the connection:
```bash
cd backend
go run main.go
```

You should see: "Connected to PostgreSQL database" and "Server starting on :8080"
