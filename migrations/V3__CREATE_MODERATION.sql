create table if not exists `post-hub-app`.post_moderation
(
    post_moderation_id varchar(36) not null
        primary key,
    moderation_id      varchar(36) null,
    post_id            varchar(36) null,
    state              boolean     null
);

alter table `post-hub-app`.post_moderation
    add constraint fk_user_moderation_id
        foreign key (moderation_id) references `post-hub-app`.user (user_id);

alter table `post-hub-app`.post_moderation
    add constraint fk_post_moderation_id
        foreign key (post_id) references `post-hub-app`.post (post_id);


