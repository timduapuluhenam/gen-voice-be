package idgenerator

import (
	"github.com/sony/sonyflake"
	// "github.com/sony/sonyflake/awsutil"
)

var SF *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	// st.MachineID = awsutil.AmazonEC2MachineID
	SF = sonyflake.NewSonyflake(st)
	if SF == nil {
		panic("sonyflake not created")
	}
}
