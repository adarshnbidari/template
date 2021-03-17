package components

import (
	"embed"
	"errors"
	"os"
)

//go:embed assets
var f embed.FS

//Setup sets up the template
func Setup(pl string) error {

	var ResErr error

	switch pl {

	case "node":
		fetchTemplate(pl, &ResErr)
	default:
		ResErr = errors.New("project not supported")
	}

	return ResErr

}

func fetchTemplate(dir string, ResErr *error) {

	dirData, err := f.ReadDir("assets/" + dir)

	handleError(err, ResErr)

	err = os.Mkdir(dir+"_template", 0777)

	handleError(err, ResErr)

	err = os.Chdir(dir + "_template")

	handleError(err, ResErr)

	for _, v := range dirData {

		fileData, err := f.ReadFile("assets/" + dir + "/" + v.Name())

		handleError(err, ResErr)

		err = os.WriteFile(v.Name(), fileData, 0777)

		handleError(err, ResErr)

	}

}

func handleError(err error, ResErr *error) {

	if err != nil {

		*ResErr = err

	}

}
