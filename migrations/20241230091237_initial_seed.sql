-- Insert data into "facilitators" table
INSERT INTO facilitators (id, name) VALUES
(1, 'Facilitator1'),
(2, 'Facilitator2');

-- Insert data into "classrooms" table
INSERT INTO classrooms (id, name, facilitator_id) VALUES
(1, 'クラスA', 1),
(2, 'クラスB', 2),
(3, 'クラスC', 1);

-- Insert data into "students" table
INSERT INTO students (id, name, login_id, classroom_id) VALUES
(1, '佐藤', 'foo123', 1),
(2, '鈴木', 'bar456', 2),
(3, '田中', 'baz789', 3),
(4, '加藤', 'hoge0000', 1),
(5, '太田', 'fuga1234', 2),
(6, '佐々木', 'piyo5678', 3),
(7, '小田原', 'fizz9999', 1),
(8, 'Smith', 'buzz777', 2),
(9, 'Johnson', 'fizzbuzz#123', 3),
(10, 'Williams', 'xxx:42', 1);
