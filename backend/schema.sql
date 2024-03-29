-- *
-- * USER TABLES
-- *
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  verified_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
  banned_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TYPE ROLE AS ENUM ('free', 'premium', 'administrator');
CREATE TABLE users_packages (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT,
  tier ROLE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  valid_from DATE NOT NULL DEFAULT NOW(),
  valid_until DATE NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_user_packages
    FOREIGN KEY(user_id)
      REFERENCES users(id)
      ON UPDATE CASCADE
      ON DELETE CASCADE
);

CREATE TYPE VERIFY_EMAIL_TYPE AS ENUM ('emailVerification', 'emailChange', 'passwordReset');
CREATE TABLE verification (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT,
  task_type VERIFY_EMAIL_TYPE NOT NULL,
  code UUID NOT NULL DEFAULT gen_random_uuid(),
  used BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  valid_until TIMESTAMP WITH TIME ZONE DEFAULT (NOW() + interval '15 minutes'),
  CONSTRAINT fk_user_verification
    FOREIGN KEY(user_id)
      REFERENCES users(id)
      ON UPDATE CASCADE
      ON DELETE SET NULL
);

-- *
-- * VIDEO RELATED TABLES
-- *

CREATE TABLE categories (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE tags (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE video (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description VARCHAR(2048) NOT NULL,
  category_id BIGINT NOT NULL,
  upvotes BIGINT NOT NULL DEFAULT 0,
  downvotes BIGINT NOT NULL DEFAULT 0,
  views BIGINT NOT NULL DEFAULT 0,
  is_premium BOOLEAN NOT NULL DEFAULT FALSE,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
  CONSTRAINT fk_category_video
    FOREIGN KEY (category_id)
    REFERENCES categories(id)
      ON DELETE SET NULL
      ON UPDATE CASCADE
);

CREATE TABLE video_tags (
  id BIGSERIAL PRIMARY KEY,
  video_id BIGSERIAL NOT NULL,
  tag_id BIGSERIAL NOT NULL,
  CONSTRAINT fk_video_tags
    FOREIGN KEY (video_id)
    REFERENCES video(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);


CREATE TYPE RESOLUTION AS ENUM ('360p', '480p', '720p');
CREATE TABLE video_files (
  id BIGSERIAL PRIMARY KEY,
  file_path VARCHAR(255) NOT NULL UNIQUE,
  video_id BIGINT NOT NULL,
  file_size BIGINT NOT NULL,
  duration BIGINT NOT NULL,
  resolution RESOLUTION NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
  CONSTRAINT fk_video_files
    FOREIGN KEY (video_id)
    REFERENCES video(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

CREATE TABLE video_thumbnails (
  id BIGSERIAL PRIMARY KEY,
  video_id BIGINT NOT NULL,
  file_size INT NOT NULL,
  file_name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
  CONSTRAINT fk_video_thumbnails
    FOREIGN KEY (video_id)
    REFERENCES video(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

-- *
-- * USER ACTIVITY
-- *

CREATE TABLE user_video_metrics (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  video_id BIGINT NOT NULL,
  time_spent_watching INT NOT NULL,
  stopped_at INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  CONSTRAINT fk_metrics_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
  CONSTRAINT fk_metrics_video
    FOREIGN KEY (video_id)
    REFERENCES video(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

CREATE TABLE comments (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  video_id BIGINT NOT NULL,
  comment VARCHAR(250) NOT NULL,
  upvotes BIGINT NOT NULL DEFAULT 0,
  downvotes BIGINT NOT NULL DEFAULT 0,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  visable BOOLEAN DEFAULT true,
  CONSTRAINT fk_comments_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
      ON DELETE SET NULL
      ON UPDATE CASCADE,
  CONSTRAINT fk_comments_video
    FOREIGN KEY (video_id)
    REFERENCES video(id)
      ON DELETE SET NULL
      ON UPDATE CASCADE
);

CREATE TABLE comments_reports (
  id BIGSERIAL PRIMARY KEY,
  reporter_id BIGINT NOT NULL,
  comment_id BIGINT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  message VARCHAR(255) NOT NULL,
  closed BOOLEAN DEFAULT false,
  CONSTRAINT fk_comments_reports
    FOREIGN KEY (comment_id)
    REFERENCES comments(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
  CONSTRAINT fk_comments_reports_users
    FOREIGN KEY (reporter_id)
    REFERENCES users(id)
      ON DELETE CASCADE
      ON UPDATE CASCADE
);

CREATE TYPE vote AS ENUM('up','down','none');
CREATE TABLE video_reaction (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  video_id BIGINT NOT NULL,
  value vote NOT NULL DEFAULT 'none',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  CONSTRAINT fk_video_reaction_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
      ON DELETE SET NULL
      ON UPDATE CASCADE,
  CONSTRAINT fk_video_reaction_video
    FOREIGN KEY (video_id)
    REFERENCES video(id)
      ON DELETE SET NULL
      ON UPDATE CASCADE
);
