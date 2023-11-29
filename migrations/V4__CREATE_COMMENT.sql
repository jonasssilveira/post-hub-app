create table if not exists `post-hub-app`.comment
(
    comment_id varchar(36) not null
        primary key,
    user_id    varchar(36) null,
    message    longtext    null,
    post_id    varchar(36) null
);
alter table `post-hub-app`.comment
    add constraint fk_post_comment
        foreign key (post_id) references `post-hub-app`.post (post_id);

alter table `post-hub-app`.comment
    add constraint fk_post_comments
        foreign key (user_id) references `post-hub-app`.user (user_id);
