CREATE TABLE user_activity(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    post_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    is_liked BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL,
    CONSTRAINT fk_postIdUserActivity FOREIGN KEY (post_id) REFERENCES post(id),
    CONSTRAINT fk_userIdUserActivity FOREIGN KEY (user_id) REFERENCES users(id)
);