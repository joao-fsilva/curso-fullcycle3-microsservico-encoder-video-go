package domain

import "time"

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}


type Job struct {
	ID string `valid:"uuid"`
	OutputBucketPath string `valid:"notnull"`
	Status string `valid:"notnull"`
	Video *Video `valid:"-"`
	VideoID string `valid:"-"`
	Error string `valid:"-"`
	CreatedAt time.Time `valid:"-"`
	UpdatedAt time.Time `valid:"-"`
}

func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job)

	if err !== nil {
		return 
	}
}