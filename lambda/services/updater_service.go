package services

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

type UpdaterService interface {
	Update(clientIdentifier string) error
}

const (
	awsRegion    = "us-east-1"
	hostedZoneID = "ZHQI7S8NKQBRD"
)

func NewUpdaterService() UpdaterService {
	return &route53UpdaterService{}
}

type route53UpdaterService struct{}

func (u *route53UpdaterService) Update(ip string) error {
	sess, err := session.NewSession(aws.NewConfig().WithRegion(awsRegion))
	if err != nil {
		return err
	}
	svc := route53.New(sess)
	input := u.getChangeInput(ip)
	_, err = svc.ChangeResourceRecordSets(input)
	return err
}

func (u *route53UpdaterService) getChangeInput(ip string) *route53.ChangeResourceRecordSetsInput {
	return &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String(hostedZoneID),
		ChangeBatch: &route53.ChangeBatch{
			Changes: u.getDomainUpsertChanges(ip),
		},
	}
}

func (u *route53UpdaterService) getDomainUpsertChanges(ip string) []*route53.Change {
	apogeeDomain := ".apogee-dev.com"
	subDomains := []string{"docker-registry", "faas", "offsite"}
	var changes []*route53.Change
	for _, a := range subDomains {
		changes = append(changes, &route53.Change{
			Action: aws.String("UPSERT"),
			ResourceRecordSet: &route53.ResourceRecordSet{
				Name: aws.String(fmt.Sprintf("%s%s", a, apogeeDomain)),
				ResourceRecords: []*route53.ResourceRecord{
					{
						Value: aws.String(ip),
					},
				},
				TTL:  aws.Int64(300),
				Type: aws.String("A"),
			},
		})
	}
	return changes
}
