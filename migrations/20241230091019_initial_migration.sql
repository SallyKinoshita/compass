-- Create "teachers" table
CREATE TABLE `teachers` (`id` int NOT NULL AUTO_INCREMENT, `name` varchar(255) NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "classrooms" table
CREATE TABLE `classrooms` (`id` int NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `teacher_id` int NOT NULL, PRIMARY KEY (`id`), INDEX `teacher_id` (`teacher_id`), CONSTRAINT `teacher_id` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "students" table
CREATE TABLE `students` (`id` int NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `login_id` varchar(255) NOT NULL, `classroom_id` int NOT NULL, PRIMARY KEY (`id`), INDEX `classroom_id` (`classroom_id`), UNIQUE INDEX `index_login_id` (`login_id`), CONSTRAINT `classroom_id` FOREIGN KEY (`classroom_id`) REFERENCES `classrooms` (`id`) ON UPDATE RESTRICT ON DELETE RESTRICT) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
