-- migrate:up
ALTER TABLE dbmate
ADD COLUMN age INT DEFAULT 0,
    ADD COLUMN is_active BOOLEAN DEFAULT TRUE;
-- migrate:down
ALTER TABLE dbmate DROP COLUMN age,
    DROP COLUMN is_active;