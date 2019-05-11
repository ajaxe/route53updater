package services

import (
	"encoding/json"
	"fmt"

	"github.com/ajaxe/route53updater/pkg/logging"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

const (
	awsRegion    = "us-east-1"
	hostedZoneID = "ZHQI7S8NKQBRD"
)

var subDomainsToUpdate = []string{"docker-registry", "faas", "offsite"}

type UpdaterService interface {
	Update(clientIdentifier string) error
}

func NewUpdaterService() UpdaterService {
	return &route53UpdaterService{}
}

type route53UpdaterService struct{}

func (u *route53UpdaterService) Update(ip string) error {
	logging.DBGLogger.Printf("update: start")
	sess, err := session.NewSession(aws.NewConfig().WithRegion(awsRegion))
	if err != nil {
		return fmt.Errorf("update error: %v", err)
	}
	logging.DBGLogger.Printf("update: session started")
	svc := route53.New(sess)
	input := u.getChangeInput(ip)
	output, err := svc.ChangeResourceRecordSets(input)
	if b, err := json.MarshalIndent(output, "", "  "); err != nil {
		logging.DBGLogger.Printf("update: run result: %s", string(b))
	}
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
	var changes []*route53.Change
	for _, a := range subDomainsToUpdate {
		d := fmt.Sprintf("%s%s", a, apogeeDomain)
		logging.DBGLogger.Printf("getDomainUpsertChanges: adding domain: %s", d)
		changes = append(changes, &route53.Change{
			Action: aws.String("UPSERT"),
			ResourceRecordSet: &route53.ResourceRecordSet{
				Name: aws.String(d),
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
