create table if not exists `post-hub-app`.user
(
    user_id   bigint unsigned auto_increment
        primary key,
    email     longtext null,
    password  longtext null,
    full_name longtext null
);

