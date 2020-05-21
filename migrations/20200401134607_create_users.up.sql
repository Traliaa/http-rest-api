create TABLE users(
    id bigserial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);
create TABLE devise(
    id bigserial not null primary key
);
create TABLE devise_iot(
    id bigserial not null primary key
);
