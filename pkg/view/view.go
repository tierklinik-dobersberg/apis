package view

import (
	"github.com/mennanov/fmutils"
	commonv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
	"google.golang.org/protobuf/proto"
)

func Apply(msg proto.Message, view *commonv1.View) {
	if view.Prune {
		fmutils.Prune(msg, view.GetFieldMask().GetPaths())
	} else {
		fmutils.Filter(msg, view.GetFieldMask().GetPaths())
	}
}

func ApplyList[E proto.Message, T []E](list T, view *commonv1.View) {
	for _, msg := range list {
		Apply(msg, view)
	}
}
