package state

type photosStruc struct {
	Small string `json:"small"`
}

type User struct {
	Id       int         `json:"id"`
	Photos   photosStruc `json:"photos"`
	Followed bool        `json:"followed"`
	Name     string      `json:"name"`
	Status   string      `json:"status"`
}

var Users = []User{
	{
		Id:       0,
		Followed: false,
		Name:     "bro",
		Status:   "zzzzzzz",
		Photos:   photosStruc{"https://i0.wp.com/blog.logrocket.com/wp-content/uploads/2019/10/golang.png?fit=730%2C412&ssl=1"},
	},
	{
		Id:       1,
		Followed: false,
		Name:     "igor",
		Status:   "uuuu",
		Photos:   photosStruc{Small: "https://computingforgeeks.com/wp-content/uploads/2018/09/install-latest-golang-centos7-ubuntu-18.04-01.png"},
	},
	{
		Id:       2,
		Followed: false,
		Name:     "igor",
		Status:   "ssssss",
		Photos:   photosStruc{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       3,
		Followed: false,
		Name:     "igor",
		Status:   "vvvvvvv",
		Photos:   photosStruc{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       4,
		Followed: false,
		Name:     "foxie",
		Status:   "jjjjjjj",
		Photos:   photosStruc{""},
	},
	{
		Id:       5,
		Followed: false,
		Name:     "bobbie",
		Status:   "mmmmmmm",
		Photos:   photosStruc{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       6,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   photosStruc{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       7,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   photosStruc{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       8,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   photosStruc{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       9,
		Followed: false,
		Name:     "igor",
		Status:   "single",
		Photos:   photosStruc{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       10,
		Followed: false,
		Name:     "igor",
		Status:   "xxxxxxxx",
		Photos:   photosStruc{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
	{
		Id:       11,
		Followed: false,
		Name:     "jasmin",
		Status:   "mmmmmmmm",
		Photos:   photosStruc{Small: "https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg"},
	},
}
