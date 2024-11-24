package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.save(todos)
}
