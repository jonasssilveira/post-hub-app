create table if not exists `post-hub-app`.post_moderation
(
    post_moderation_id bigint unsigned auto_increment
        primary key,
    moderation_id      bigint unsigned null,
    post_id            bigint unsigned null,
    state              boolean      null
);

alter table `post-hub-app`.post_moderation
    add constraint fk_user_moderation_id
        foreign key (moderation_id) references `post-hub-app`.user (user_id);

alter table `post-hub-app`.post_moderation
    add constraint fk_post_moderation_id
        foreign key (post_id) references `post-hub-app`.post (post_id);


