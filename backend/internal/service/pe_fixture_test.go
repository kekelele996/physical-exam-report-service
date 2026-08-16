package service

import (
	"log/slog"
	"os"
	"testing"

	"github.com/blueship581/gbcheckup/internal/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newPeDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil { t.Fatalf("open sqlite: %v", err) }
	if err := db.AutoMigrate(&model.User{}, &model.Enterprise{}, &model.Examinee{}, &model.Package{}, &model.PackageItem{}, &model.Registration{}, &model.ExamResult{}, &model.AbnormalMetric{}, &model.Report{}, &model.GroupOrder{}); err != nil { t.Fatalf("migrate: %v", err) }
	return db
}

func peLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
}
