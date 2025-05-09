-- Drop triggers
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
DROP TRIGGER IF EXISTS update_tasks_updated_at ON tasks;

-- Drop trigger function
DROP FUNCTION IF EXISTS update_updated_at_column;

-- Drop tables
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;

-- Drop extension
DROP EXTENSION IF EXISTS "uuid-ossp";
