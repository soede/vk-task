# Алгоритм Дейкстры на Go
Этот проект реализует алгоритм Дейкстры для нахождения кратчайшего пути от стартовой точки до конечной точки в графе, представленном в виде двумерного поля.

## Как запустить
Для того чтобы запустить код:
1. Клонируйте репозиторий:
    ```bash
   git clone https://github.com/soede/vk-task.git
   cd vk-task
   ```
2. Установите зависимости:
    ```bash 
    go mod tidy
   ```
3. Запустите ```main.go```
   ```bash
   go run main.go
   ```
## Cложность
- Время работы алгоритма ```O(V log(V))```, где ```V``` – это количество клеток на поле (размер поля rows * cols).
- Пространственная сложность ```O(V)``` так как хранятся все элементы поля, стоимости и предыдущие точки.
