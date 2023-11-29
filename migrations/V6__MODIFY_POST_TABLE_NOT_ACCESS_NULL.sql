-- Modify title column to disallow null values
ALTER TABLE `post-hub-app`.post
    MODIFY title longtext NOT NULL;

-- Modify message column to disallow null values
ALTER TABLE `post-hub-app`.post
    MODIFY message longtext NOT NULL;

-- Modify user_id column to disallow null values
ALTER TABLE `post-hub-app`.post
    MODIFY user_id varchar(36)  NOT NULL;