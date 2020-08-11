
-- +migrate Up
INSERT INTO `users` (`name`, `password`, `created_at`, `updated_at`) VALUES ("admin", "$2a$10$FRc9ICx.bln.SEjlWlDu1.wqVs1BUkiVcmfb1jJxeAdOCvak5dvCe", "2020-08-08 00:00:00", "2020-08-08 00:00:00");
-- +migrate Down
DELETE FROM `users` WHERE `name` = "admin";
