package web

import (
	"backup-x/entity"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed dist
var indexEmbedFile embed.FS

// IndexConfig 填写配置信息
func IndexConfig(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFS(indexEmbedFile, "dist/index.html")
	if err != nil {
		log.Println(err)
		return
	}

	conf, err := entity.GetConfigCache()
	if err == nil {
		tmpl.Execute(writer, &writtingData{Config: conf, Version: os.Getenv(VersionEnv)})
		return
	}

	// default config
	// 获得环境变量
	backupConf := []entity.BackupConfig{}
	for i := 0; i < 16; i++ {
		backupConf = append(backupConf, entity.BackupConfig{SaveDays: 30, SaveDaysS3: 60, StartTime: 1, Period: 1440, BackupType: 0})
	}
	conf = entity.Config{
		BackupConfig: backupConf,
	}

	tmpl.Execute(writer, &writtingData{Config: conf, Version: os.Getenv(VersionEnv)})
}
