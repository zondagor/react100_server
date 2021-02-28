package state

type User struct {
	Id       int
	Photos   struct{ Small string }
	Followed bool
	Name     string
	Status   string
}

var Users = []User{
	{
		Id:       1,
		Followed: false,
		Name:     "bro",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://i0.wp.com/blog.logrocket.com/wp-content/uploads/2019/10/golang.png?fit=730%2C412&ssl=1"},
	},
	{
		Id:       2,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://computingforgeeks.com/wp-content/uploads/2018/09/install-latest-golang-centos7-ubuntu-18.04-01.png"},
	},
	{
		Id:       3,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       4,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       5,
		Followed: false,
		Name:     "foxie",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       6,
		Followed: false,
		Name:     "bobbie",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       7,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       8,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       9,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       10,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       11,
		Followed: false,
		Name:     "jasmin",
		Status:   "single",
		Photos:   struct{ Small string }{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
}
