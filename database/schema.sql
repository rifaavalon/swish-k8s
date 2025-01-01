CREATE TABLE environment_usage (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    environment_name VARCHAR(255) NOT NULL,
    cpu_requested VARCHAR(50),
    memory_requested VARCHAR(50),
    cpu_used VARCHAR(50),
    memory_used VARCHAR(50),
    start_time TIMESTAMP,
    end_time TIMESTAMP
);
