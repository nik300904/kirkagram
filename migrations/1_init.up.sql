-- Создание таблицы "user"
CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,                -- Автоинкрементный идентификатор
    username VARCHAR(50) NOT NULL UNIQUE, -- Уникальное имя пользователя
    email VARCHAR(100) NOT NULL UNIQUE,    -- Уникальный адрес электронной почты
    password VARCHAR(255) NOT NULL,         -- Хэш пароля
    bio TEXT,                               -- Биография пользователя
    profile_pic VARCHAR(255),               -- URL профиля пользователя
    followers INTEGER[],                        -- Массив ID подписчиков
    following INTEGER[]                         -- Массив ID пользователей, на которых подписан
);

-- Создание таблицы "post"
CREATE TABLE IF NOT EXISTS "post" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    image_url TEXT NOT NULL,
    caption TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

-- Создание таблицы "like"
CREATE TABLE IF NOT EXISTS "like" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES "post"(id) ON DELETE CASCADE,
    UNIQUE (user_id, post_id) -- Пользователь может лайкнуть пост только один раз
);

-- Создание таблицы "comment"
CREATE TABLE IF NOT EXISTS "comment" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES "post"(id) ON DELETE CASCADE
);
