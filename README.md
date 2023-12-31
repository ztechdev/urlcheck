# Техническое задание

Нужно реализовать утилиту для обработки списка URL адресов.
Для каждого URL адреса нужно получить:

* время выполнения запроса
* время обработки
  Эти данные нужно сохранить в отчёт, для последующей аналитики.

## Список URL адресов

Сейчас наш список адресов это файл с примерно 250 записями.
Но отдел аналитики планирует, через полгода начать его дополнять от 200 до 1_500 записей в месяц.

Формат файла:

```
https://www.zonatelecom.ru/
https://www.zonatelecom.ru/support/services/pochta
...
https://www.zonatelecom.ru/fsin-services/services-institutions
```

Предусмотреть возможность перехода с файла на БД.

## Отчёт

Пример отчёта

| URL                                                            | Ошибка   | Кол-во байт ответа | Время обработки в сек |
|----------------------------------------------------------------|----------|--------------------|-----------------------|
| https://www.zonatelecom.ru/                                    |          | 5426321            | 0.678                 |
| https://www.zonatelecom.ru/fsin-services/services-institutions | some err | 0                  | 0                     |

На первом этапе, можно сохранять статистику в файл, но скорее всего потом понадобится база данных для удобства обработки
и анализа.
Формат файла не принципиален, главное чтобы каждая запись содержала все необходимые данные.
