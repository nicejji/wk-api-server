begin;

insert into "user" (nickname, password, bio) values 
  ('wkiskas', 'sdf@#$f$$#23', 'shh... senior'),
  ('wissa', 'mypass_1', 'that is nice!'),
  ('empfity', '%#@genXl', 'here it is...');

insert into "chat" (title, private) values
  (null, true),
  (null, true),
  ('Group: wkiskas, wissa, empfity', false);

insert into "user_chat" (user_id, chat_id, is_creator, blocked) values
  (1, 1, true, false),
  (2, 1, false, false),
  (3, 2, true, false),
  (1, 2, false, false),
  (3, 3, true, false),
  (1, 3, false, false),
  (2, 3, false, false);

insert into "message" (user_id, chat_id, content) values 
  (1, 1, 'hello here!'),
  (2, 1, 'nice to see you :)'),
  (3, 2, 'wow you here'),
  (1, 2, 'yesss'),
  (3, 3, 'welcome guys'),
  (2, 3, 'that is nice'),
  (1, 3, 'supper nice');

commit;

