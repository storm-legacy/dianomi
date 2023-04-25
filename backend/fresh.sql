-- CREATE ADMINISTRATOR ACCOUNT
INSERT INTO users(
  email,
  password,
  verified_at
) VALUES (
  'admin@legacystorm.com',
  '$argon2id$v=19$m=65536,t=3,p=2$DZ1Zfefj6zjvyWYVqRnk+Q$xIYVNrhbM/tCPK5PIRVQwlarg8H/QuUVCeNqYi2MS4M',
  NOW()
),
(
  'premium@legacystorm.com',
  '$argon2id$v=19$m=65536,t=3,p=2$DZ1Zfefj6zjvyWYVqRnk+Q$xIYVNrhbM/tCPK5PIRVQwlarg8H/QuUVCeNqYi2MS4M',
  NOW()
);
-- GIVE IT PERMISSIONS
INSERT INTO users_packages(
  user_id,
  tier,
  valid_from,
  valid_until
) VALUES (
  (SELECT id FROM users WHERE email='admin@legacystorm.com'),
  'administrator',
  NOW(),
  (NOW() + interval '100 years')
),
(
  (SELECT id FROM users WHERE email='premium@legacystorm.com'),
  'premium',
  NOW(),
  (NOW() + interval '6 months')
)
;