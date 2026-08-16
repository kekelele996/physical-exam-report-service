package service
import "testing"
func TestAbnormalJudgementAndLevel(t *testing.T){
 if IsAbnormal("10-20","10"){t.Fatal("10 should not be abnormal in 10-20")}
 if !IsAbnormal(">10","10"){t.Fatal("10 should be abnormal for >10")}
 if !IsAbnormal("<5","5"){t.Fatal("5 should be abnormal for <5")}
 if GuessAbnormalLevel("300")!="moderate"{t.Fatalf("300 should be moderate, got %s",GuessAbnormalLevel("300"))}
 if GuessAbnormalLevel("3000")!="severe"{t.Fatalf("3000 should be severe, got %s",GuessAbnormalLevel("3000"))}
}
