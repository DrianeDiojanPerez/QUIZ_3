-- Filename: migrations/000003_add_todo_indexes.up.sql

CREATE INDEX IF NOT EXISTS todo_list_task_name_idx ON todo_list USING GIN(to_tsvector('simple', task_name));
CREATE INDEX IF NOT EXISTS todo_list_priority_idx ON todo_list USING GIN(to_tsvector('simple', priority));
CREATE INDEX IF NOT EXISTS todo_list_status_idx ON todo_list USING GIN(status);