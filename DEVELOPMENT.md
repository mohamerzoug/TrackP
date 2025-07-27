# TrackP Development Guide

## Quick Start

### Prerequisites Check
Before starting, ensure you have:
- **PostgreSQL** installed and running
- **Node.js 16+** and npm
- **Go 1.19+**

### 1. Database Setup
```sql
-- Create database in PostgreSQL
CREATE DATABASE trackp;
```

Run the setup script:
```bash
psql -U postgres -d trackp -f database/setup.sql
```

### 2. Install Dependencies
```bash
# Install all dependencies
npm run install:all

# Or install individually:
npm run install:frontend
npm run install:backend
```

### 3. Start Development Environment
```bash
# Start both frontend and backend simultaneously
npm start

# Or use VS Code tasks:
# Ctrl+Shift+P -> "Tasks: Run Task" -> "Start TrackP Development Environment"
```

## Architecture Overview

### Frontend (http://localhost:3000)
- **React 19** with TypeScript
- **@dnd-kit** for drag-and-drop
- **Axios** for API calls
- Modern component architecture

### Backend (http://localhost:8080)
- **Go + Gin** REST API
- **PostgreSQL** database
- CORS enabled for frontend

### Database Schema
```sql
projects (id, title, description, created_at)
tasks (id, title, description, status, due_date, project_id, created_at)
```

## Available Scripts

### Root Level
- `npm start` - Start both frontend and backend
- `npm run build:all` - Build production versions
- `npm run install:all` - Install all dependencies

### Frontend
- `npm run dev:frontend` - Start React dev server
- `npm run build:frontend` - Build for production
- `npm run test:frontend` - Run tests

### Backend
- `npm run dev:backend` - Start Go server
- `npm run build:backend` - Build binary

## Development Workflow

1. **Start Development**: Run `npm start` or use VS Code task
2. **Make Changes**: Edit files in `frontend/src/` or `backend/`
3. **Test**: Both servers auto-reload on file changes
4. **Build**: Use `npm run build:all` for production

## Project Structure

```
TrackP/
├── frontend/                 # React TypeScript app
│   ├── src/
│   │   ├── components/      # UI components
│   │   ├── api.ts          # API calls
│   │   ├── types.ts        # TypeScript types
│   │   └── App.tsx         # Main app
│   └── package.json
├── backend/                  # Go API server
│   ├── main.go             # Server + handlers
│   └── go.mod
├── database/                 # SQL scripts
│   └── setup.sql
├── .vscode/                  # VS Code tasks
├── package.json             # Root scripts
└── README.md
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/projects` | List all projects |
| POST | `/api/projects` | Create project |
| GET | `/api/projects/:id` | Get project |
| PUT | `/api/projects/:id` | Update project |
| DELETE | `/api/projects/:id` | Delete project |
| GET | `/api/projects/:id/tasks` | Get project tasks |
| POST | `/api/projects/:id/tasks` | Create task |
| PUT | `/api/tasks/:id` | Update task |
| DELETE | `/api/tasks/:id` | Delete task |

## Environment Configuration

Create `backend/.env` from `backend/.env.example`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=trackp
```

## Troubleshooting

### Database Connection Issues
1. Ensure PostgreSQL is running
2. Check connection details in `backend/main.go`
3. Verify database `trackp` exists

### Frontend Build Errors
1. Run `npm run install:frontend`
2. Clear node_modules: `rm -rf frontend/node_modules && npm run install:frontend`

### Backend Compilation Issues
1. Run `npm run install:backend`
2. Check Go version: `go version`

### CORS Issues
- Backend includes CORS middleware for `http://localhost:3000`
- Modify in `main.go` if using different ports

## Production Deployment

### Build Applications
```bash
npm run build:all
```

### Frontend Deployment
- Static files in `frontend/build/`
- Deploy to any static hosting (Netlify, Vercel, S3)

### Backend Deployment
- Binary: `backend/trackp-server`
- Set environment variables for production database
- Deploy to cloud provider (AWS, GCP, Heroku)

## Contributing

1. Fork repository
2. Create feature branch
3. Make changes with tests
4. Submit pull request

## VS Code Integration

Recommended extensions:
- Go (Google)
- ES7+ React/Redux/React-Native snippets
- TypeScript Importer
- PostgreSQL (Chris Kolkman)

Tasks available in Command Palette:
- Start TrackP Development Environment
- Start Frontend Development Server
- Start Backend Server
- Build Frontend for Production
- Build Backend Binary
