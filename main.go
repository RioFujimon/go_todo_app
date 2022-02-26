package main

import "udemy/go_todo_app/app/controllers"

func main() {
	/*
	// Configのテスト
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	// Logのテスト
	log.Println("test")

	fmt.Println(models.Db)
	*/

	/*
	// Userデータの作成
	u := &models.User{}
	u.Name = "test2"
	u.Email = "test2@example.com"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()
	*/

	/*
	// Userデータの取得
	u, _ := models.GetUser(1)
	fmt.Println(u)

	// Userデータの更新
	u.Name = "Test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

	// Userデータの削除
	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
	*/

	/*
	// Todoデータの作成
	user, _ := models.GetUser(2)
	user.CreateTodo("First Todo")
	*/

	/*
	// 1件のTodoデータを取得
	t, _ := models.GetTodo(1)
	fmt.Println(t)
	*/


	/*
	user, _ := models.GetUser(3)
	user.CreateTodo("Third Todo")
	*/

	/*
	// 複数のTodoデータを取得
	todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	} 
	*/


	/*
	// UserのTodoを取得
	user2, _ := models.GetUser(2)
	todos, _ := user2.GetTodoByUser()
	for _, v := range todos {
		fmt.Println(v)
	}
	*/

	/*
	// Todoの更新
	t, _ := models.GetTodo(1)
	t.Content = "Update Todo"
	t.UpdateTodo()
	*/

	/*
	// Todoの削除
	t, _ := models.GetTodo(3)
	t.DeleteTodo()
	*/

	controllers.StartMainServer()

	/*
	// EmailによるUserの取得テスト
	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)
	*/

	/*
	// Sessionの作成のテスト
	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)
	*/

	/*
	// Sessionが存在するかのテスト
	valid, _ := session.CheckSession()
	fmt.Println(valid)
	*/
}