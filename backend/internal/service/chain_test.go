package service
import (
 "testing"
 "time"
 "github.com/blueship581/gbcheckup/internal/model"
 "github.com/blueship581/gbcheckup/internal/repository"
)
func TestStatsAggregation(t *testing.T){
 db:=newPeDB(t)
 p:=&model.Package{Name:"基础体检",PackageType:"standard",Price:100,Status:"active"}; db.Create(p)
 pi:=&model.PackageItem{PackageID:p.ID,ItemName:"血常规"}; db.Create(pi)
 e:=&model.Examinee{Name:"张三"}; db.Create(e)
 now:=time.Now()
 db.Create(&model.Registration{ExamineeID:e.ID,PackageID:p.ID,GuideNo:"G1",Status:"registered",RegisteredAt:now})
 db.Create(&model.Registration{ExamineeID:e.ID,PackageID:p.ID,GuideNo:"G2",Status:"registered",RegisteredAt:now})
 db.Create(&model.ExamResult{RegistrationID:1,ExamineeID:e.ID,PackageItemID:pi.ID,IsAbnormal:true,Status:"entered"})
 db.Create(&model.ExamResult{RegistrationID:2,ExamineeID:e.ID,PackageItemID:pi.ID,IsAbnormal:true,Status:"entered"})
 db.Create(&model.ExamResult{RegistrationID:1,ExamineeID:e.ID,PackageItemID:pi.ID,IsAbnormal:false,Status:"entered"})
 rr:=repository.NewRegistrationRepository(db); er:=repository.NewExamResultRepository(db)
 pkg,err:=rr.CountGroupByPackage(); if err!=nil{t.Fatal(err)}
 if len(pkg)!=1 || pkg[0].Count!=2{t.Fatalf("pkg count wrong: %+v",pkg)}
 ab,err:=er.CountAbnormalGroupByItem(); if err!=nil{t.Fatal(err)}
 if len(ab)!=1 || ab[0].Count!=2{t.Fatalf("abnormal count wrong: %+v",ab)}
}
