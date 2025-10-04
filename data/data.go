package data

/*
	type todo struct {
		id        int
		task      string
		completed bool
	}

if use small letter then it will become private and c.json will not able to export it and api response will be empty.
*/
/*
Without json tag it will also work but mainly in DB camelcase or snakecase so that is why  we explicitly provide names for it.
*/

/*
If we dont put binding and wrong json or no json is provided still it will add with default values.
*/
type Todo struct {
	Id        int    `json:"id" binding:"required"`
	Task      string `json:"task" binding:"required"`
	Completed bool   `json:"completed"`
}

var Todos = []Todo{
	{Id: 1,
		Task:      "go",
		Completed: false},
}
