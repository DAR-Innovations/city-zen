-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    role VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Create departments table
CREATE TABLE IF NOT EXISTS departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    phone VARCHAR(20)
);

-- Create issues table
CREATE TABLE IF NOT EXISTS issue (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    tags VARCHAR(255),
    status VARCHAR(50) CHECK (status IN ('PENDING', 'IN PROGRESS', 'DONE')) DEFAULT 'PENDING',
    is_public BOOLEAN DEFAULT TRUE,
    longitude DECIMAL(9, 6),
    latitude DECIMAL(9, 6),
    user_id INT REFERENCES users(id),
    image VARCHAR(255)
);

-- Create reports table
CREATE TABLE IF NOT EXISTS reports (
    id SERIAL PRIMARY KEY,
    issue_id INT REFERENCES issue(id),
    description TEXT,
    image VARCHAR(255),
    department_id INT REFERENCES departments(id),
    user_id INT REFERENCES users(id)
);

-- Create task_types table
CREATE TABLE IF NOT EXISTS task_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    department_id INT REFERENCES departments(id)
);

-- Create tasks table
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    issue_id INT REFERENCES issue(id),
    task_type_id INT REFERENCES task_types(id),
    urgency VARCHAR(50) CHECK (urgency IN ('HIGH', 'MEDIUM', 'LOW')),
    complexity VARCHAR(50) CHECK (complexity IN ('HIGH', 'MEDIUM', 'LOW')),
    processed_at DATE
);
