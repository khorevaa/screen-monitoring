CREATE TABLE IF NOT EXISTS `pages` (
	`id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`title`	TEXT NOT NULL UNIQUE,
	`visible`	INTEGER NOT NULL DEFAULT 0
);