package dal

import (
	"fmt"
	"os"
	"testing"

	"github.com/gocode/learnpack/config"
	"github.com/gocode/learnpack/models"
)

func TestMain(m *testing.M) {
	initErr := config.InitConfig()
	if initErr != nil {
		fmt.Fprintf(os.Stderr, "Config init failed %v\n", initErr)
		os.Exit(1)
	}
	initErr = InitSQL()
	if initErr != nil {
		fmt.Fprintf(os.Stderr, "DB init failed %v\n", initErr)
		os.Exit(1)
	}
	defer UninitSQL()
	exitVal := m.Run()
	fmt.Println("After")
	os.Exit(exitVal)
}

func TestDalUserNew(t *testing.T) {
	userDal, err := NewUserDal()

	user := models.User{
		Name:  "xyz",
		Email: "abc@xyz.com",
		Pass:  "Pass",
	}

	createdUser, err := userDal.NewUser(&user)
	defer userDal.DeleteUser(createdUser.ID)

	if err != nil {
		t.Errorf("NewUser Insert Failed %s", err)
	}
	if int64(createdUser.ID) <= 0 {
		t.Errorf("NewUser User ID not present")
	}
}

func TestTwo(t *testing.T) {
	t.Log("Test 2")
}
