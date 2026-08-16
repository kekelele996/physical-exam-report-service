package service
import (
 "errors"
 "testing"
 "github.com/blueship581/gbcheckup/internal/repository"
 "github.com/blueship581/gbcheckup/internal/util"
)
func TestErrorChainSentinels(t *testing.T){
 db:=newPeDB(t)
 er:=repository.NewExamResultRepository(db); ex:=repository.NewExamineeRepository(db); rr:=repository.NewRegistrationRepository(db)
 if _,err:=er.FindByID(999);!errors.Is(err,util.ErrNotFound){t.Fatalf("result FindByID=%v",err)}
 if _,err:=ex.FindByID(999);!errors.Is(err,util.ErrNotFound){t.Fatalf("examinee FindByID=%v",err)}
 if _,err:=rr.FindByID(999);!errors.Is(err,util.ErrNotFound){t.Fatalf("registration FindByID=%v",err)}
}
