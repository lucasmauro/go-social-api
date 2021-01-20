INSERT INTO users (name, nickname, email, password)
	VALUES
    	("John Doe", "JD", "johndoe@mail.com", "$2a$10$amAKB9/UadeTcO1QqgBpPesJ5kMbmnSqvW8w2fqA86.LFNbGxFEgm"),
        ("John Toe", "JT", "johntoe@mail.com", "$2a$10$amAKB9/UadeTcO1QqgBpPesJ5kMbmnSqvW8w2fqA86.LFNbGxFEgm"),
        ("John Foe", "JF", "johnfoe@mail.com", "$2a$10$amAKB9/UadeTcO1QqgBpPesJ5kMbmnSqvW8w2fqA86.LFNbGxFEgm");
-- mySuperSecretPassword!123

INSERT INTO followers (user_id, follower_id)
	VALUES
    	(1, 2),
        (1, 3),
        (2, 3);