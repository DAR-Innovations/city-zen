-- Inserting mock data
-- Departments
INSERT INTO departments (name, description, phone, created_at, updated_at) VALUES
    ('IT Department', 'Responsible for all technical issues', '1234567890', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('HR Department', 'Handles human resources management', '0987654321', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Task Types
INSERT INTO task_types (name, description, created_at, updated_at) VALUES
    ('Bug', 'Fix issues or errors in the system', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Feature Request', 'Implement new features or improvements', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Users
INSERT INTO users (first_name, last_name, phone, is_verified, role, password, created_at, updated_at) VALUES
    ('John', 'Doe', '1234567890', TRUE, 'ADMIN', 'test1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Jane', 'Smith', '0987654321', TRUE, 'USER', 'test1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Employees
INSERT INTO employees (first_name, last_name, phone, is_verified, department_id, role, password, created_at, updated_at) VALUES
    ('Alice', 'Johnson', '1112223333', TRUE, 1, 'ADMIN', 'test1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Bob', 'Brown', '4445556666', FALSE, 2, 'EMPLOYEE', 'test1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Issues
INSERT INTO issues (name, description, is_completed, longitude, latitude, author_id, image_url, created_at, updated_at) VALUES
    ('Website Bug', 'The website has a bug when loading the homepage.', FALSE, 45.0, 60.0, 1, 'image_url_1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Database Issue', 'The database is returning error codes.', TRUE, 46.0, 61.0, 2, 'image_url_2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Department Tasks
INSERT INTO department_tasks (issue_id, title, description, task_type_id, status, urgency, complexity, department_id, created_at, updated_at) VALUES
    (1, 'Fix homepage bug', 'Resolve the bug affecting the homepage.', 1, 'PENDING', 'HIGH', 'MEDIUM', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Fix database errors', 'Resolve the database connection errors.', 1, 'IN PROGRESS', 'HIGH', 'HIGH', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Volunteer Tasks
INSERT INTO volunteer_tasks (issue_id, title, description, status, urgency, complexity, volunteer_id, created_at, updated_at) VALUES
    (1, 'Bug Testing', 'Test the homepage bug fix after it is applied.', 'PENDING', 'MEDIUM', 'LOW', NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Database Testing', 'Test the database after fixing errors.', 'PENDING', 'HIGH', 'HIGH', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- User Reports
INSERT INTO user_reports (issue_id, description, image_url, reported_by, created_at, updated_at) VALUES
    (1, 'Found a critical bug on the homepage.', 'report_image_1', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Found a bug in the database.', 'report_image_2', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Department Reports
INSERT INTO department_reports (issue_id, description, image_url, reported_by, created_at, updated_at) VALUES
    (1, 'Department report about the homepage bug.', 'department_report_1', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'Department report about database issues.', 'department_report_2', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);