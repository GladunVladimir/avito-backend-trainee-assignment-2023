# Тестовое задание для стажёра Backend
# Сервис динамического сегментирования пользователей

### Проблема:

В Авито часто проводятся различные эксперименты — тесты новых продуктов, тесты интерфейса, скидочные и многие другие.
На архитектурном комитете приняли решение централизовать работу с проводимыми экспериментами и вынести этот функционал в отдельный сервис.

### Задача:

Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

# Запуск проекта

```sh
git clone git@github.com:GladunVladimir/json-api.git
docker compose build
docker compose up
```
# Работа с API

Создание сегмента

В случае удачного добавления выводится "Segment created"
В случае наличия вводимого сегмента в базе выводится "Segment is already exists"


![image](https://github.com/GladunVladimir/json-api/blob/main/Screenshots/%D0%94%D0%BE%D0%B1%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D1%81%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82%D0%B0.png)


Удаление сегмента

В случае удачного удаления выводится "Segment removed"
В случае отсутствия вводимого сегмента в базе выводится "No such segment exists"

![image](https://github.com/GladunVladimir/json-api/blob/main/Screenshots/%D0%A3%D0%B4%D0%B0%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D1%81%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82%D0%B0.png)


Добавление пользователя в сегмент

В случае удачного добавления выводится "User/Users added to segment"
В случае отсутствия нужного сегмента в базе будет выведено "Segment not found"
В случае наличия пользователя в вводимом сегменте будет выведено "User already in this segment"

![imae](https://github.com/GladunVladimir/json-api/blob/main/Screenshots/%D0%94%D0%BE%D0%B1%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8F%20%D0%B2%20%D1%81%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82.png)


Удаление пользователя из сегмента

В случае удачного удаления выводится "User/Users added to segment"
В случае отсутствия пользователя с таким сегментом выводится "User not found"
В случае отсутствия сегмента у пользователя "Segment not found"

![image](https://github.com/GladunVladimir/json-api/blob/main/Screenshots/%D0%A3%D0%B4%D0%B0%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8F%20%D0%B8%D0%B7%20%D1%81%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82%D0%B0.png)


Сегменты выбранного пользователя

В случае удачного нахождения нужной записи выводится список slug сегментов пользователя
В случае отсутствия данного пользователя в базе данных выводится "User not found"

![image](https://github.com/GladunVladimir/json-api/blob/main/Screenshots/%D0%A1%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82%D1%8B%20%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8F.png)
