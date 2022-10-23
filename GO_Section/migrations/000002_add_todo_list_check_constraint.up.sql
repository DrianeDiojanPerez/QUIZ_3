-- Filename: migrations/000002_add_todo_list_check_constraint.up.sql

ALTER TABLE todo_list ADD CONSTRAINT status_length_check CHECK (array_length(status, 1) BETWEEN 1 AND 5);