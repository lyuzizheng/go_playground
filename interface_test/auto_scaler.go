package interface_test

const (
	CheckPSMUsage = "check-psm"
)

var MethodMap = map[string]func(service AutoScalerService){
	"aaa": AutoScalerService.CheckPSMUsage,
}

type AutoScalerService interface {
	CheckPSMUsage()
	CheckStatus()
}

type AutoScalerServiceImpl struct {
	RedirectServiceImpl
	PushServiceImpl
}

func (a *AutoScalerServiceImpl) CheckPSMUsage() {
	panic("implement me")
}

