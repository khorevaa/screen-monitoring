CREATE TABLE IF NOT EXISTS `pages` (
	`id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`title`	TEXT NOT NULL CHECK(title <> '') UNIQUE,
	`visible`	INTEGER NOT NULL DEFAULT 0
);
