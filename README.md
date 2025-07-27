# TrackP - Lightweight Project Management Tool

TrackP is a modern, single-user project management application designed for clarity and efficiency. It provides a clean, visual interface to organize projects and track tasks from start to finish using a Kanban-style board.

![TrackP Screenshot](https://via.placeholder.com/800x400/667eea/ffffff?text=TrackP+Dashboard)

## ✨ Features

### Core Functionality
- **📊 Project Dashboard**: Central view listing all created projects with enhanced UI
- **📝 Create and Manage Projects**: Add new projects with titles and descriptions
- **📋 Kanban-Style Task Board**: Visual workflow with "To Do," "In Progress," and "Done" columns
- **🔄 Full Task Management (CRUD)**: Create, view, update, and delete tasks
- **🖱️ Drag-and-Drop Interface**: Intuitive task status updates by dragging between columns
- **📅 Due Date Tracking**: Set and track task due dates with overdue indicators
- **📱 Responsive Design**: Works seamlessly on desktop and mobile devices

### Enhanced Features (v1.1.0)
- **🔍 Real-time Search**: Find projects instantly with debounced search functionality
- **📊 Smart Sorting**: Sort projects by name (A-Z) or creation date (newest first)
- **💾 Persistent Preferences**: Search and sort preferences saved automatically
- **⚡ Performance Optimized**: Debounced inputs and optimized re-renders
- **🎨 Enhanced UI/UX**: Better visual hierarchy, loading states, and error handling
- **🛠️ Developer Experience**: Custom hooks, utilities, and comprehensive testing tools
- **🔧 Demo Mode**: Run without database setup using in-memory storage

## 🚀 Quick Start (Demo Mode)

The fastest way to try TrackP without setting up a database:

```bash
# Clone or extract the project
cd TrackP

# Start demo backend (in-memory storage)
cd backend
go run demo.go

# In another terminal, start frontend
cd ../frontend
npm install
npm start
```

Visit `http://localhost:3000` to see TrackP in action!

## 🏗️ Technology Stack

### Frontend (React + TypeScript)
- **React 19** with TypeScript for type safety
- **@dnd-kit** for modern drag-and-drop functionality
- **Axios** for API communication
- **CSS3** for styling with modern design patterns
- **Custom Hooks**: useLocalStorage, useAsync, useDebounce, useForm, and more
- **Utility Functions**: Date formatting, validation, error handling, and performance optimizations

### Backend (Go + Gin)
- **Go** with Gin framework for high-performance API
- **PostgreSQL** driver for database connectivity (production)
- **In-Memory Storage** option for demo mode
- **CORS** middleware for frontend-backend communication
- RESTful API design with comprehensive error handling

### Database (PostgreSQL)
- **Projects table**: Stores project information
- **Tasks table**: Stores task details with foreign key relationship to projects
- Simple, efficient relational schema

## 🛠️ Production Setup

### Prerequisites

- **Node.js** (v16 or higher)
- **Go** (v1.19 or higher)
- **PostgreSQL** (v12 or higher)

### Database Setup

1. Install PostgreSQL and create a database named `trackp`
2. Run the setup script:
   ```bash
   psql -U postgres -d trackp -f database/setup.sql
   ```
3. Update database connection settings in `backend/main.go` if needed

### Installation

1. **Install all dependencies**:
   ```bash
   npm run install:all
   ```

2. **Start production backend**:
   ```bash
   cd backend
   go run main.go
   ```

3. **Start frontend development server**:
   ```bash
   cd frontend
   npm start
   ```

## 📖 Usage

### Getting Started
1. **Creating Projects**: Click "Add New Project" on the dashboard
2. **Searching Projects**: Use the search box to find projects by name or description
3. **Sorting Projects**: Click "Sort by Name" or "Sort by Date" to organize your view
4. **Managing Tasks**: Select a project to open its Kanban board

### Task Management  
5. **Adding Tasks**: Click "Add Task" and fill in the details
6. **Moving Tasks**: Drag and drop tasks between columns to update their status
7. **Editing Tasks**: Click on task details to edit (feature can be extended)
8. **Deleting**: Use the "×" button on projects or tasks to delete them

### Advanced Features
- **Persistent Search**: Your search preferences are saved automatically
- **Keyboard Navigation**: Tab through interface elements for accessibility
- **Responsive Design**: Use on any device - desktop, tablet, or mobile
- **Error Recovery**: Helpful error messages guide you through any issues

## 🔌 API Endpoints

### Projects
- `GET /api/projects` - Get all projects
- `POST /api/projects` - Create a new project
- `GET /api/projects/:id` - Get a specific project
- `PUT /api/projects/:id` - Update a project
- `DELETE /api/projects/:id` - Delete a project

### Tasks
- `GET /api/projects/:id/tasks` - Get all tasks for a project
- `POST /api/projects/:id/tasks` - Create a new task
- `PUT /api/tasks/:id` - Update a task
- `DELETE /api/tasks/:id` - Delete a task

## 📁 Project Structure

```
TrackP/
├── frontend/                 # React TypeScript frontend
│   ├── src/
│   │   ├── components/      # React components
│   │   │   ├── ProjectDashboard.tsx    # Enhanced with search & sort
│   │   │   ├── KanbanBoard.tsx         # Drag & drop task board
│   │   │   ├── TaskCard.tsx            # Individual task component
│   │   │   ├── TaskColumn.tsx          # Task column component
│   │   │   ├── LoadingSpinner.tsx      # Reusable loading component
│   │   │   └── LoadingSpinner.css
│   │   ├── api.ts          # API service functions
│   │   ├── types.ts        # TypeScript type definitions
│   │   ├── config.ts       # Application configuration
│   │   ├── utils.ts        # Utility functions
│   │   ├── hooks.ts        # Custom React hooks
│   │   └── App.tsx         # Main application component
│   └── package.json
├── backend/                  # Go Gin backend
│   ├── main.go             # Production server with PostgreSQL
│   ├── demo.go             # Demo server with in-memory storage
│   ├── .env.example        # Environment configuration template
│   └── go.mod
├── database/                 # Database setup scripts
│   └── setup.sql
├── docs/                     # Documentation
│   ├── DEVELOPMENT.md      # Detailed development guide
│   ├── DATABASE_SETUP.md   # Database setup instructions
│   ├── TESTING.md          # Comprehensive testing guide
│   └── CHANGELOG.md        # Version history and features
├── .vscode/                  # VS Code tasks and settings
│   └── tasks.json
├── .github/                  # GitHub and Copilot configuration
│   └── copilot-instructions.md
└── README.md               # This file
```

## 🔧 Development

See [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md) for detailed development instructions and [docs/TESTING.md](docs/TESTING.md) for comprehensive testing guidelines.

### Available Scripts

```bash
# Root level convenience scripts
npm start                    # Start both frontend and backend
npm run install:all         # Install all dependencies  
npm run build:all           # Build for production

# Individual components
npm run dev:frontend        # Start React dev server
npm run dev:backend         # Start Go server
npm run build:frontend      # Build React for production
npm run build:backend       # Build Go binary
```

## 🎯 VS Code Integration

TrackP includes VS Code tasks for easy development:

1. Open Command Palette (`Ctrl+Shift+P`)
2. Type "Tasks: Run Task"
3. Select "Start TrackP Development Environment"

This will start both backend and frontend servers simultaneously.

## 🚀 Production Deployment

### Frontend
```bash
cd frontend
npm run build
# Deploy the build/ folder to any static hosting
```

### Backend
```bash
cd backend
go build -o trackp-server main.go
# Deploy the binary with PostgreSQL connection
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🔮 Future Enhancements

### Version 1.2.0 (Planned)
- [ ] Task priority levels with color coding
- [ ] Task due date reminders and notifications
- [ ] Keyboard shortcuts for power users
- [ ] Dark mode theme toggle
- [ ] Export/import functionality (JSON/CSV)
- [ ] Task templates for common workflows
- [ ] Better mobile experience and PWA support

### Version 1.3.0 (Planned)
- [ ] Time tracking and reporting
- [ ] Task comments and file attachments
- [ ] Project templates and cloning
- [ ] Advanced filtering and search
- [ ] Analytics dashboard
- [ ] Multi-user support with permissions

### Version 2.0.0 (Future)
- [ ] Real-time collaboration
- [ ] API integrations (GitHub, Slack, etc.)
- [ ] Custom fields and task types
- [ ] Automation rules and workflows
- [ ] Mobile native applications

---

## 📈 Version Information

**Current Version**: TrackP v1.1.0 - Enhanced Features  
**Release Date**: July 27, 2025  
**Compatibility**: React 19, Go 1.19+, PostgreSQL 12+

### What's New in v1.1.0
✨ **Real-time search** with debounced performance  
📊 **Smart sorting** with persistent preferences  
🛠️ **Custom hooks** and utility functions  
🎨 **Enhanced UI/UX** with loading states  
🔧 **Demo mode** for easy testing  
📖 **Comprehensive documentation** and testing guides

See [docs/CHANGELOG.md](docs/CHANGELOG.md) for complete version history.
- [ ] Export/import functionality
- [ ] Real-time collaboration
- [ ] Mobile app version
- [ ] Dark mode theme
- [ ] Advanced filtering and search
- [ ] Reporting and analytics

## 🐛 Troubleshooting

**Backend won't start?**
- Check if PostgreSQL is running
- Try demo mode: `go run demo.go`
- Verify Go version: `go version`

**Frontend compilation errors?**
- Run `npm install` in frontend directory
- Clear cache: `rm -rf node_modules package-lock.json && npm install`

**CORS issues?**
- Ensure backend is running on port 8080
- Check CORS configuration in `main.go` or `demo.go`

## 📞 Support

- 📖 Check [DEVELOPMENT.md](DEVELOPMENT.md) for detailed setup
- 🗄️ See [DATABASE_SETUP.md](DATABASE_SETUP.md) for database configuration
- 🐛 Report issues in the GitHub Issues section
- 💡 Feature requests are welcome!

---

**Made with ❤️ using React, Go, and PostgreSQL**
