-- Modify state column to disallow null values
ALTER TABLE `post-hub-app`.user
    MODIFY email longtext NOT NULL;

-- Modify post_id column to disallow null values
ALTER TABLE `post-hub-app`.user
    MODIFY password longtext NOT NULL;

-- Modify moderation_id column to disallow null values
ALTER TABLE `post-hub-app`.user
    MODIFY full_name longtext NOT NULL;