create table public.responses
(
    id             text    not null
        primary key
        unique,
    input_text     text    not null,
    response_text  text    not null,
    is_reported    boolean default false,
    response_token integer not null,
    created_date   integer not null
);

alter table public.responses
    owner to postgres;

