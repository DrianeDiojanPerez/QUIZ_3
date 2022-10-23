-- Filename: migrations/000002_add_todo_check_constraint.down.sql

ALTER TABLE todo_list DROP CONSTRAINT IF EXISTS status_length_check;