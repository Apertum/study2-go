package main

import (
	"fmt"
	"study2-go/internal/config"
	"time"

	"study2-go/internal/pattern"

	"github.com/sirupsen/logrus"
)

func init() {
	config.LogsInit()
	logo()
}

func logo() {
	logrus.Info("  ▄██████▄   ▄██████▄     ████████▄  ")
	logrus.Info("▄█▀▀██████ ▄██▀▀████▄    ▀▀▀▀▀▀███")
	logrus.Info("██    ▀▀▀  ██    ████          ███")
	logrus.Info("██   ▄███▄ ██    ████      ▄████▀")
	logrus.Info("██  ▀█████ ██    ████    ▄████▀")
	logrus.Info("██    ████ ██    ████  ▄████▀")
	logrus.Info("██▄▄▄█████ ▀██▄▄████▀  ████████████")
	logrus.Info("▀▀▀▀▀▀▀    ▀▀████▀    ████████████")
}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	cfg := config.Load()
	logrus.WithField("address", cfg.Address).Info("загрузка конфигурации")

	logrus.Info("\n\n=== === === === === Singleton")

	for i := range 10 {
		go func(i int) {
			fmt.Printf("Адрес singleton: %p\n", pattern.GetSingleton())
		}(i)
	}
	time.Sleep(500 * time.Millisecond)

	logrus.Info("\n\n=== === === === === PRoto :")

	folder := &pattern.Folder{
		Children: []pattern.Node{
			&pattern.File{Name: "file1"},
			&pattern.File{Name: "file2"},
			&pattern.Folder{
				Children: []pattern.Node{
					&pattern.File{Name: "file3"},
					&pattern.File{Name: "file4"},
				},
				File: pattern.File{Name: "subfolder"},
			},
		},
		File: pattern.File{Name: "root"},
	}
	logrus.Info("Original:\n%s\n", folder)

	clone := folder.Clone()
	logrus.Info("Clone:\n%s\n", clone)

	logrus.Info("\n\n=== === === === === Builder :")

	o := pattern.NewObject().SetMode(10).SetPath(`root`)
	logrus.Info("Builder: ", o)

	oBest := pattern.NewObjectBest(pattern.WithMode(10), pattern.WithPath(`root`))
	logrus.Info("BuilderBest: ", oBest)

	logrus.Info("\n\n=== === === === === FAbric :")

	mysql := pattern.NewConnector("mysql://...")
	mysql.Query("SELECT something FROM list")

	pg := pattern.NewConnector("postgresql://...")
	pg.Query("SELECT something FROM list")

	logrus.Info("\n\n=== === === === === FAbric Abstract :")

	artdecoFactory := pattern.GetFactory("art-deco")
	modernFactory := pattern.GetFactory("modern")

	artdecoChair := artdecoFactory.MakeChair("дуб")
	artdecoTable := artdecoFactory.MakeTable("дуб")

	modernChair := modernFactory.MakeChair("ясень")
	modernTable := modernFactory.MakeTable("ясень")

	logrus.Info(artdecoChair.Print(), artdecoTable.Print())
	logrus.Info(modernChair.Print(), modernTable.Print())

	logrus.Info("сервер остановлен")
}
