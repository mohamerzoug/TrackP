-- TrackP Database Setup
-- Run these commands in your PostgreSQL database

-- Create the database (run this as a superuser)
-- CREATE DATABASE trackp;

-- Connect to the trackp database and run the following:

-- Create projects table
CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create tasks table
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'To Do',
    due_date DATE,
    project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_tasks_project_id ON tasks(project_id);
CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);
CREATE INDEX IF NOT EXISTS idx_projects_created_at ON projects(created_at);
CREATE INDEX IF NOT EXISTS idx_tasks_created_at ON tasks(created_at);

-- Insert sample data (optional)
INSERT INTO projects (title, description) VALUES 
('Sample Project', 'This is a sample project to get you started with TrackP'),
('Website Redesign', 'Complete redesign of the company website with modern UI/UX');

INSERT INTO tasks (title, description, status, project_id) VALUES 
('Set up development environment', 'Install all necessary tools and dependencies', 'Done', 1),
('Create database schema', 'Design and implement the database structure', 'Done', 1),
('Implement user interface', 'Build the React frontend components', 'In Progress', 1),
('API development', 'Create RESTful API endpoints', 'To Do', 1),
('Research design trends', 'Look into current web design trends and best practices', 'To Do', 2),
('Create wireframes', 'Design the layout and structure of new pages', 'To Do', 2);
