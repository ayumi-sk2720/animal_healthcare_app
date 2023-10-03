
-- +migrate Up
-- ==========================
-- master
-- ==========================
CREATE TABLE IF NOT EXISTS `pet` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` VARCHAR(255) NULL COMMENT '名前',
  `birth_day` VARCHAR(255) NULL COMMENT '誕生日',
  `sex` VARCHAR(255) NULL COMMENT '性別',
  `created_at` DATETIME NOT NULL COMMENT '作成日時',
  `updated_at` DATETIME NOT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'ペットマスタ';

CREATE TABLE IF NOT EXISTS `schedule` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pet_id` BIGINT NOT NULL COMMENT 'ペットID',
  `title` VARCHAR(255) NULL COMMENT 'タイトル',
  `date` DATETIME NOT NULL COMMENT '日時',
  `location` VARCHAR(45) NULL COMMENT '場所',
  `created_at` DATETIME NOT NULL COMMENT '作成日時',
  `updated_at` DATETIME NOT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_schedule1`
    FOREIGN KEY (`pet_id`)
    REFERENCES `pet` (`id`)
    -- ON DELETE CASCADE
    -- ON UPDATE NO ACTION
  )
ENGINE = InnoDB
COMMENT = '予定';


-- +migrate Down
DROP TABLE `schedule`;
DROP TABLE `pet`;