package server

func (s *FiberServer) RegisterRoutes() {
	api := s.App.Group("/api")

    userGroup := api.Group("/users")
    userHandler := NewUserHandler(s.db)

	userGroup.Get("/", userHandler.GetUsers)
	userGroup.Get("/:id", userHandler.GetUserById)
	userGroup.Patch("/:id", userHandler.UpdateUser)
    userGroup.Delete("/:id", userHandler.DeleteUser)
}
