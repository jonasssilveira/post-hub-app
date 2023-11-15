-- Modify post_id column to disallow null values
ALTER TABLE `post-hub-app`.comment
    MODIFY post_id bigint unsigned NOT NULL;

-- Modify user_id column to disallow null values
ALTER TABLE `post-hub-app`.comment
    MODIFY user_id bigint unsigned NOT NULL;

-- Modify message column to disallow null values
ALTER TABLE `post-hub-app`.comment
    MODIFY message longtext NOT NULL;