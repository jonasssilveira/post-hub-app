create table if not exists `post-hub-app`.comment
(
    comment_id bigint unsigned auto_increment
    primary key,
    user_id    bigint unsigned null,
    message    longtext        null,
    post_id    bigint unsigned null
);
alter table `post-hub-app`.comment
    add constraint fk_post_comment
        foreign key (post_id) references `post-hub-app`.post (post_id);

alter table `post-hub-app`.comment
    add constraint fk_post_comments
        foreign key (user_id) references `post-hub-app`.user (user_id);
