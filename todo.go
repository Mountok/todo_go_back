package todoapp

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        string `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	Itemid int
}
