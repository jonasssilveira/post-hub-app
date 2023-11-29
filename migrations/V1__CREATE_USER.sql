create table if not exists `post-hub-app`.user
(
    user_id   varchar(36) not null
        primary key,
    email     longtext null,
    password  longtext null,
    full_name longtext null
);

