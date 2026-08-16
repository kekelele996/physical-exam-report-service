package service
import (
 "context"
 "errors"
 "testing"
 "github.com/blueship581/gbcheckup/internal/repository"
 "github.com/blueship581/gbcheckup/internal/util"
)
func expectNoPanicPe(t *testing.T,name string,fn func() error) error { t.Helper(); var err error; func(){ defer func(){ if r:=recover(); r!=nil { t.Fatalf("%s panicked: %v",name,r) } }(); err=fn() }(); return err }
func TestMissingResultAndRegistration(t *testing.T){
 db:=newPeDB(t)
 er:=repository.NewExamResultRepository(db); rr:=repository.NewRegistrationRepository(db)
 erSvc:=NewExamResultService(er,rr,repository.NewAbnormalMetricRepository(db),peLogger())
 regSvc:=NewRegistrationService(rr,repository.NewExamineeRepository(db),repository.NewPackageRepository(db),repository.NewPackageItemRepository(db),er,peLogger())
 err:=expectNoPanicPe(t,"Review",func() error { return erSvc.Review(context.Background(),999) })
 if !errors.Is(err,util.ErrNotFound){t.Fatalf("Review=%v",err)}
 _,err=regSvc.Get(context.Background(),999)
 if !errors.Is(err,util.ErrNotFound){t.Fatalf("Get=%v",err)}
}
