create table if not EXISTS `users` (
  id string PRIMARY KEY NOT NULL,
  username string UNIQUE NOT NULL,
  password string NOT NULL,
  reset_token UNQUE NOT NULL,
  role string NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME,
  deleted_at DATETIME
)