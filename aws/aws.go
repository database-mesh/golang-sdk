package aws
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
)

type Sessions map[string]*session.Session

type aws struct {
    credentials []credential
}

type credential struct {
    region string
    accessKey string
    secretAccessKey string
}

func NewSessions() *aws {
    return &aws{credentials: []credential{}}
}

func (s *aws)SetCredential(region, accessKey, secretAccessKey string) *aws{
    s.credentials = append(s.credentials, credential{
        region: region,
	accessKey: accessKey,
	secretAccessKey: secretAccessKey,
    }) 
    return s
}

func (s *aws)Build() Sessions {
    sess := map[string]*session.Session{}
    for _, v := range s.credentials{
        as, err := newAWSSession(v.region, v.accessKey, v.secretAccessKey)
	if err != nil {
            continue
	}
        sess[v.region] = as 
    }
    return sess
}

func newAWSSession(region, ak, sk string) (*session.Session, error) {
    c := credentials.NewStaticCredentials(ak, sk, "")
    awscfg := &aws.Config{Credentials: c}
    awscfg.WithRegion(v.region)
    return session.NewSessionWithOptions(session.Options{
        Config: *awscfg,
        AssumeRoleTokenProvider: token.StdinStderrTokenProvider,
        SharedConfigState:       session.SharedConfigEnable,
    })
}
