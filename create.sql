begin;

create table "user" (
  id        serial primary key,
  nickname  varchar(20),
  password  varchar(128),
  bio       varchar(150),
  image     varchar(700),
  created   timestamp default now()
);

create table "chat" (
  id        serial primary key,
  title     varchar(50),
  private   boolean,
  created   timestamp default now()
);

create table "user_chat" (
  id            serial primary key,
  user_id       int references "user"(id),
  chat_id       int references "chat"(id),
  is_creator    boolean,
  blocked       boolean,
  created       timestamp default now()
);

create table "message" (
  id        serial primary key,
  user_id   int references "user"(id),
  chat_id   int references "chat"(id),
  content   varchar(300),
  created   timestamp default now()
);

create table "attachment" (
  id          serial primary key,
  message_id  int references "message"(id),
  image       varchar(700)
);

commit;
