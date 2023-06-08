package vertex

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Nickname string
	Age      int
	Birth    string
	State    string
}

func CreateVertex() {
	reader := bufio.NewReader(os.Stdin)

	// 1行目の入力を受け取る
	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimSpace(line1)
	numUsers, _ := strconv.Atoi(line1)

	// Userのスライスを作成
	users := make([]User, numUsers)

	// ユーザー情報を受け取る
	for i := 0; i < numUsers; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		values := strings.Split(line, " ")

		age, _ := strconv.Atoi(values[1])
		users[i] = User{
			Nickname: values[0],
			Age:      age,
			Birth:    values[2],
			State:    values[3],
		}
	}
	for _, user := range users {
		output := fmt.Sprintf(
			"User{\nnickname : %s\nold : %d\nbirth : %s\nstate : %s\n}\n",
			user.Nickname, user.Age, user.Birth, user.State)
		fmt.Print(output)
	}

}
