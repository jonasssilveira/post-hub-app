-- Modify state column to disallow null values
ALTER TABLE `post-hub-app`.post_moderation
    MODIFY state boolean NOT NULL DEFAULT false;

-- Modify post_id column to disallow null values
ALTER TABLE `post-hub-app`.post_moderation
    MODIFY post_id varchar (36) NOT NULL;

-- Modify moderation_id column to disallow null values
ALTER TABLE `post-hub-app`.post_moderation
    MODIFY moderation_id varchar (36) NOT NULL;