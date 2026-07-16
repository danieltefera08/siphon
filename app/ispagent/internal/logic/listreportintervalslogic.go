package logic

import (
	"context"

	"siphon/app/ispagent/internal/ispclient"
	"siphon/app/ispagent/internal/svc"
	"siphon/app/ispagent/ispagent"
	"siphon/common/isp"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListReportIntervalsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListReportIntervalsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListReportIntervalsLogic {
	return &ListReportIntervalsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListReportIntervalsLogic) ListReportIntervals(in *ispagent.ListReportIntervalsReq) (*ispagent.ListReportIntervalsRes, error) {
	reportIntervals := l.svcCtx.IspClient.ReportIntervals()

	categories := make([]*ispagent.ReportCategoryInfo, 0, len(reportIntervals))
	for cat, d := range reportIntervals {
		typ, cmd := isp.DecodeMessageID(int(cat))
		categories = append(categories, &ispagent.ReportCategoryInfo{
			Category:        int32(cat),
			Name:            ispclient.CategoryMessageName(cat),
			IntervalSeconds: int64(d.Seconds()),
			NoFreshCheck:    l.svcCtx.IspClient.CategoryNoFreshCheck(cat),
			Type:            typ,
			Command:         cmd,
			KeyAttrs:        ispclient.CategoryKeyAttrs(cat),
		})
	}

	return &ispagent.ListReportIntervalsRes{
		Categories: categories,
	}, nil
}
