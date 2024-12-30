# Используйте официальный образ PostgreSQL
FROM postgres:15

# Установите необходимые зависимости (если нужно)
# RUN apt-get update && apt-get install -y <ваши_пакеты>

# Копируйте скрипты инициализации (если есть)
# COPY ./init.sql /docker-entrypoint-initdb.d/

# Установите переменные окружения
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=12345
ENV POSTGRES_DB=kirkagram

# Открываем порт
EXPOSE 5432

# По умолчанию запускается сервер PostgreSQL
CMD ["postgres"]
