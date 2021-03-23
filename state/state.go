package state

type photos struct {
	Small string `json:"small"`
	Large string `json:"large"`
}

type contacts struct {
	Github    string `json:"github"`
	Youtube   string `json:"youtube"`
	Twitter   string `json:"twitter"`
	Instagram string `json:"instagram"`
}

type User struct {
	Id                        int      `json:"id"`
	Followed                  bool     `json:"followed"`
	LookingForAJob            bool     `json:"lookingForAJob"`
	Name                      string   `json:"name"`
	FullName                  string   `json:"fullName"`
	Status                    string   `json:"status"`
	AboutMe                   string   `json:"aboutMe"`
	LookingForAJobDescription string   `json:"lookingForAJobDescription"`
	Contacts                  contacts `json:"contacts"`
	Photos                    photos   `json:"photos"`
}

var Users []User

func createUser(
	id int,
	followed, lookingForAJob bool,
	name, fullName, status, photo, aboutMe, lookingForAJobDescription string,
) {
	user := User{
		Id:                        id,
		Name:                      name,
		FullName:                  fullName,
		Status:                    status,
		AboutMe:                   aboutMe,
		Followed:                  followed,
		LookingForAJob:            lookingForAJob,
		LookingForAJobDescription: lookingForAJobDescription,
		Photos:                    photos{Small: photo},
		Contacts: contacts{
			Github:    "https://github.com",
			Youtube:   "https://youtube.com",
			Instagram: "https://instagram.com",
			Twitter:   "https://twitter.com",
		},
	}
	Users = append(Users, user)
}

func InitTempState() {
	Users = make([]User, 0, 12)

	createUser(0, false, true, "igor", "igor valerievich", "single",
		"https://i0.wp.com/blog.logrocket.com/wp-content/uploads/2019/10/golang.png?fit=730%2C412&ssl=1",
		"cool guy", "html=))")

	createUser(1, false, true, "igor", "igor valerievich", "single",
		"https://computingforgeeks.com/wp-content/uploads/2018/09/install-latest-golang-centos7-ubuntu-18.04-01.png",
		"cool guy", "html=))")

	createUser(2, false, true, "igor", "igor valerievich", "single",
		"https://upload.wikimedia.org/wikipedia/commons/1/1c/Dmitry_Nagiev_2017_3.jpg",
		"cool guy", "html=))")

	createUser(3, false, true, "igor", "igor valerievich", "single",
		"https://i0.wp.com/blog.logrocket.com/wp-content/uploads/2019/10/golang.png?fit=730%2C412&ssl=1",
		"cool guy", "html=))")

	createUser(4, false, true, "igor", "igor valerievich", "single",
		"https://i0.wp.com/blog.logrocket.com/wp-content/uploads/2019/10/golang.png?fit=730%2C412&ssl=1",
		"cool guy", "html=))")

	createUser(5, false, true, "igor", "igor valerievich", "single",
		"https://i0.wp.com/blog.logrocket.com/wp-content/uploads/2019/10/golang.png?fit=730%2C412&ssl=1",
		"cool guy", "html=))")

	createUser(6, false, true, "igor", "igor valerievich", "single",
		"https://i0.wp.com/blog.logrocket.com/wp-content/uploads/2019/10/golang.png?fit=730%2C412&ssl=1",
		"cool guy", "html=))")
}

// there were 12 users
