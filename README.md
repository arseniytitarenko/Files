# Files

### Архитектура
Архитектура программы и запросы соответствует предложенной схеме, а именно: \
`storage-service`, `minio` и `postgres` для него \
`analysis-service`, `minio` и `postgres` для него \
`nginx` в качестве `API Gateway` \
Сервисы соответствуют чистой архитектуре

### Спецификация
Примеры всех запросов по сервисам есть в postman коллекции: \
https://www.postman.com/olympguide/files/collection/nz65tz5/files?action=share&creator=40644038

### Запуск
Достаточно склонировать репозиторий и запустить из корня командой: \
`docker-compose up --build`