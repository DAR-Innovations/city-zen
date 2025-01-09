-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    role VARCHAR(50) CHECK (role IN ('ADMIN', 'USER')) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Create employees table
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    department_id INT REFERENCES departments(id),
    role VARCHAR(50) CHECK (role IN ('ADMIN', 'EMPLOYEE')) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Create departments table
CREATE TABLE IF NOT EXISTS departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    phone VARCHAR(20),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Create task_types table
CREATE TABLE IF NOT EXISTS task_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Create department_task_types table
CREATE TABLE IF NOT EXISTS department_task_types (
    id SERIAL PRIMARY KEY,
    department_id INT REFERENCES departments(id),
    task_type_id INT REFERENCES task_types(id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Create issues table
CREATE TABLE IF NOT EXISTS issue (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    is_completed BOOLEAN DEFAULT FALSE,
    longitude DECIMAL(9, 6),
    latitude DECIMAL(9, 6),
    author_id INT REFERENCES users(id),
    image_url VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Create reports table
CREATE TABLE IF NOT EXISTS user_reports (
    id SERIAL PRIMARY KEY,
    issue_id INT REFERENCES issue(id),
    description TEXT,
    image_url VARCHAR(255),
    reported_by INT REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- Create reports table
CREATE TABLE IF NOT EXISTS department_reports (
    id SERIAL PRIMARY KEY,
    issue_id INT REFERENCES issue(id),
    description TEXT,
    image_url VARCHAR(255),
    reported_by INT REFERENCES departments(id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
    );

-- Create volunteer_tasks table
CREATE TABLE IF NOT EXISTS volunteer_tasks (
    id SERIAL PRIMARY KEY,
    issue_id INT REFERENCES issue(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) CHECK (status IN ('PENDING', 'IN PROGRESS', 'DONE')) DEFAULT 'PENDING',
    urgency VARCHAR(50) CHECK (urgency IN ('HIGH', 'MEDIUM', 'LOW')),
    complexity VARCHAR(50) CHECK (complexity IN ('HIGH', 'MEDIUM', 'LOW')),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
    );

-- Create department_tasks table
CREATE TABLE IF NOT EXISTS department_tasks (
    id SERIAL PRIMARY KEY,
    issue_id INT REFERENCES issue(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    task_type_id INT REFERENCES task_types(id),
    status VARCHAR(50) CHECK (status IN ('PENDING', 'IN PROGRESS', 'DONE')) DEFAULT 'PENDING',
    urgency VARCHAR(50) CHECK (urgency IN ('HIGH', 'MEDIUM', 'LOW')),
    complexity VARCHAR(50) CHECK (complexity IN ('HIGH', 'MEDIUM', 'LOW')),
    department_id INT REFERENCES departments(id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);
