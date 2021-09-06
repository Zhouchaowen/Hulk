package generates

import (
	"github.com/srlemon/gen-id/utils"
	"time"
)

func generatorRandTime() time.Time {
	begin, _ := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	end, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 00:00:00")
	return time.Unix(utils.RandInt64(begin.Unix(), end.Unix()), 0)
}

type TimeRule struct {
}

func (s *TimeRule) GetParamType() ParamType {
	return Time
}

func (s *TimeRule) GetNonComplianceCount() int {
	return 0
}

func (s *TimeRule) GetNonComplianceOtherTypes() []ParamType {
	return []ParamType{
		Bool,
		Int,
		Float64,
		String,
	}
}
