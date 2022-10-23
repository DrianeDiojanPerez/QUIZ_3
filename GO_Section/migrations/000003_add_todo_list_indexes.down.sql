-- Filename: migrations/000003_add_todo_indexes.down.sql

DROP INDEX IF EXISTS todo_list_task_name_idx;
DROP INDEX IF EXISTS todo_list_priority_idx;
DROP INDEX IF EXISTS todo_list_status_idx;