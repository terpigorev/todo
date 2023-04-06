package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Создаем новый сканнер для чтения ввода пользователя
    scanner := bufio.NewScanner(os.Stdin)

    // Создаем пустой список для хранения задач
    var todos []string

    // Бесконечный цикл для ввода задач
    for {
        // Запрашиваем ввод новой задачи у пользователя
        fmt.Print("Введите новое дело или введите 'q', чтобы выйти: ")
        scanner.Scan()

        // Сохраняем введенную задачу в переменную todo
        todo := scanner.Text()

        // Если пользователь ввел "q", прерываем цикл
        if todo == "q" {
            break
        }

        // Добавляем введенную задачу в список todos
        todos = append(todos, todo)
    }

    // Выводим список задач на экран
    fmt.Println("Ваши задачи:")
    for i, todo := range todos {
        fmt.Printf("%d. %s\n", i+1, todo)
    }
}
