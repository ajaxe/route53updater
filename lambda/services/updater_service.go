package services

type UpdaterService interface {
	Update(clientIdentifier string) error
}

func NewUpdaterService() UpdaterService {
	return &route53UpdaterService{}
}

type route53UpdaterService struct {
	route53Client route53
}

func (u *route53UpdaterService) Update(clientIdentifier string) error {
	// 1: read s3 config
	// 2: using client identifier get domain list
	// 3:
}
