package memory

import (
	"testing"

	"github.com/pschlump/gatus/core"
	"github.com/pschlump/gatus/storage/store/common"
	"github.com/pschlump/gatus/storage/store/common/paging"
)

func BenchmarkShallowCopyEndpointStatus(b *testing.B) {
	endpoint := &testEndpoint
	status := core.NewEndpointStatus(endpoint.Group, endpoint.Name)
	for i := 0; i < common.MaximumNumberOfResults; i++ {
		AddResult(status, &testSuccessfulResult)
	}
	for n := 0; n < b.N; n++ {
		ShallowCopyEndpointStatus(status, paging.NewEndpointStatusParams().WithResults(1, 20))
	}
	b.ReportAllocs()
}
