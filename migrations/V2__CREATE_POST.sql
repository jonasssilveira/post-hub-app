
create table if not exists `post-hub-app`.post
(
    post_id       bigint unsigned auto_increment
    primary key,
    title         longtext        null,
    message       longtext        null,
    user_id       bigint unsigned null,
    moderation_id bigint unsigned null
);
alter table `post-hub-app`.post
    add constraint fk_users_posts
        foreign key (user_id) references `post-hub-app`.user (user_id);

alter table `post-hub-app`.post
    add constraint fk_users_post
        foreign key (moderation_id) references `post-hub-app`.user(user_id);