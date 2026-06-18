package logic

import (
	"context"
	"siphon/app/gis/gis"

	"siphon/app/gis/internal/svc"
	"siphon/common/tool"
	"siphon/third_party/extproto"

	"github.com/mmcloughlin/geohash"
	"github.com/zeromicro/go-zero/core/logx"
)

type DecodeGeoHashLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecodeGeoHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecodeGeoHashLogic {
	return &DecodeGeoHashLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DecodeGeoHash 将 geohash 字符串解码为中心点坐标及边界框（bounding box）。
func (l *DecodeGeoHashLogic) DecodeGeoHash(in *gis.DecodeGeoHashReq) (*gis.DecodeGeoHashRes, error) {
	if in.Geohash == "" {
		return nil, tool.NewErrorByPbCode(extproto.Code__1_01_PARAM_MISSING, "geohash")
	}
	// 中心点
	lat, lon := geohash.DecodeCenter(in.Geohash)
	// bbox
	box := geohash.BoundingBox(in.Geohash)
	return &gis.DecodeGeoHashRes{
		Point: &gis.Point{
			Lat: lat,
			Lon: lon,
		},
		LatMin: box.MinLat,
		LatMax: box.MaxLat,
		LonMin: box.MinLng,
		LonMax: box.MaxLng,
	}, nil
}
