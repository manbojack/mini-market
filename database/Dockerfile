FROM postgres:13

# Копируем все init sql скрипты в контейнер
COPY ./init-scripts/ /docker-entrypoint-initdb.d/

# Можно добавить ENV переменные или другие настройки при необходимости

EXPOSE 5432
