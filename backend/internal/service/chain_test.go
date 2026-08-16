package service
import (
 "context"
 "testing"
 "github.com/blueship581/gbcheckup/internal/model"
 "github.com/blueship581/gbcheckup/internal/repository"
)
func TestFollowUpChain(t *testing.T){
 db:=newPeDB(t)
 m:=&model.AbnormalMetric{ExamineeID:1,PackageItemID:1,AbnormalLevel:"mild",Value:"300",RefValueRange:"10-20",FollowUpStatus:"pending"}; db.Create(m)
 r:=repository.NewAbnormalMetricRepository(db); svc:=NewAbnormalMetricService(r,peLogger())
 if _,err:=svc.UpdateFollowUp(context.Background(),m.ID,"bogus","");err==nil{t.Fatal("invalid status should be rejected")}
 got,err:=svc.UpdateFollowUp(context.Background(),m.ID,"done","复查")
 if err!=nil{t.Fatalf("UpdateFollowUp: %v",err)}
 if got.FollowUpStatus!="done"{t.Fatalf("follow_up_status = %s, want done",got.FollowUpStatus)}
}
