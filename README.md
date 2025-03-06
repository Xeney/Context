# Context

Этот код на Go демонстрирует управление горутинами с помощью контекста. Программа запускает 5 задач, каждая из которых выполняется случайное время (от 1 до 5 секунд). Если задача не завершается за 3 секунды, она отменяется через контекст. Используется sync.WaitGroup для синхронизации завершения всех задач.

## Как это работает?
1. Создается контекст с таймаутом 3 секунды.
2. Запускаются 5 горутин, каждая выполняет задачу.
3. Если задача завершается раньше таймаута, выводится "Task X finished".
4. Если таймаут истекает, незавершенные задачи отменяются, и выводится "Task X canceled".

## Пример вывода
```
Task 1 started
Task 2 started
Task 3 started
Task 4 started
Task 5 started
Task 2 finished
Task 5 finished
Task 1 canceled
Task 3 canceled
Task 4 canceled
Program finished
```
