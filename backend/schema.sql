CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  verified_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TYPE ROLE AS ENUM ('free', 'premium', 'administrator');
CREATE TABLE users_packages (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
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
  user_id BIGINT NOT NULL,
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

CREATE TABLE revoked_tokens (
  id BIGSERIAL PRIMARY KEY,
  token VARCHAR(1024) NOT NULL UNIQUE,
  user_id BIGINT NOT NULL,
  valid_until TIMESTAMP WITH TIME ZONE NOT NULL,
  CONSTRAINT fk_user_revoked_token
    FOREIGN KEY(user_id)
      REFERENCES users(id)
      ON UPDATE CASCADE
      ON DELETE CASCADE
);