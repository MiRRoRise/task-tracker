package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func add(task string, tasks *[]Task) {
	task = strings.Trim(task, "\"")
	t1 := Task{len(*tasks) + 1, task, "todo"}
	*tasks = append(*tasks, t1)

	data, err := json.MarshalIndent(*tasks, "", "  ")
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}

	err = os.WriteFile("file.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл: ", err)
		return
	}

	fmt.Println("Output: Task added successfully (ID: ", t1.Id, ")")
}

func update(id int, task string, tasks *[]Task) {
	found := false
	task = strings.Trim(task, "\"")
	for ind, elem := range *tasks {
		if elem.Id == id {
			(*tasks)[ind].Description = task
		}
	}

	if !found {
		fmt.Print("Задача с таким id не найдена")
	}

	data, err := json.MarshalIndent(*tasks, "", " ")
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}

	err = os.WriteFile("file.json", data, 0644)
	if err != nil {
		fmt.Print("Ошибка записи в файл: ", err)
	}

	fmt.Println("Update succesfully")
}

func delete(id int, tasks *[]Task) {
	found := false
	for ind, elem := range *tasks {
		if elem.Id == id {
			*tasks = append((*tasks)[:ind], (*tasks)[ind+1:]...)
			found = true
		}
	}

	if !found {
		fmt.Print("Задача с таким id не найдена")
	}

	data, err := json.MarshalIndent(*tasks, "", " ")
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}

	err = os.WriteFile("file.json", data, 0644)
	if err != nil {
		fmt.Print("Ошибка записи в файл: ", err)
	}

	fmt.Println("Delete succesfully")
}

func make_progress(id int, tasks *[]Task) {
	found := false
	for ind, elem := range *tasks {
		if elem.Id == id {
			(*tasks)[ind].Status = "in-progress"
			found = true
		}
	}

	if !found {
		fmt.Print("Задача с таким id не найдена")
	}

	data, err := json.MarshalIndent(*tasks, "", " ")
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}

	err = os.WriteFile("file.json", data, 0644)
	if err != nil {
		fmt.Print("Ошибка записи в файл: ", err)
	}
}

func make_done(id int, tasks *[]Task) {
	found := false
	for ind, elem := range *tasks {
		if elem.Id == id {
			(*tasks)[ind].Status = "done"
			found = true
		}
	}

	if !found {
		fmt.Print("Задача с таким id не найдена")
	}

	data, err := json.MarshalIndent(*tasks, "", " ")
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}

	err = os.WriteFile("file.json", data, 0644)
	if err != nil {
		fmt.Print("Ошибка записи в файл: ", err)
	}
}

func see_all(tasks *[]Task) {

}

func main() {
	_, err := os.Stat("file.json")
	if err != nil {
		_, err := os.Create("file.json")
		if err != nil {
			fmt.Println("Ошибка создания файла: ", err)
			return
		}
	}

	file, err := os.ReadFile("file.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла: ", err)
		return
	}
	fmt.Print(string(file))
	var tasks []Task
	if len(file) != 0 {
		err = json.Unmarshal(file, &tasks)
		if err != nil {
			fmt.Println("Ошибка десереализации: ", err)
			return
		}
	}

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		data := strings.Split(scanner.Text(), "\"")
		option := strings.Split(data[0], " ")

		switch option[0] {
		case "add":
			task := data[1]
			add(task, &tasks)
		case "update":
			id, task := option[1], data[1]
			id1, err := strconv.Atoi(id)
			if err != nil {
				fmt.Print("Ошибка конвертации: ", err)
				return
			}
			if id1 < 1 {
				fmt.Print("id должен быть натуральным числом")
				return
			}
			update(id1, task, &tasks)
		case "delete":
			id := option[1]
			id1, err := strconv.Atoi(id)
			if err != nil {
				fmt.Print("Ошибка конвертации: ", err)
				return
			}
			if id1 < 1 {
				fmt.Print("id должен быть натуральным числом")
				return
			}
			delete(id1, &tasks)
		case "mark-in-progress":
			id := option[1]
			id1, err := strconv.Atoi(id)
			if err != nil {
				fmt.Print("Ошибка конвертации: ", err)
				return
			}
			if id1 < 1 {
				fmt.Print("id должен быть натуральным числом")
				return
			}
			make_progress(id1, &tasks)
		case "mark-done":
			id := option[1]
			id1, err := strconv.Atoi(id)
			if err != nil {
				fmt.Print("Ошибка конвертации: ", err)
				return
			}
			if id1 < 1 {
				fmt.Print("id должен быть натуральным числом")
				return
			}
			make_done(id1, &tasks)
		case "list":
			see_all(&tasks)
		case "exit":
			return
		default:
			fmt.Println("Такой команды нет")
		}
	}
}
