package common

import (
	"context"
	"english_app/pkg/errs"
	"gopkg.in/vansante/go-ffprobe.v2"
	"time"
)

func GetVideoData(url string) (float64, errs.MessageErr) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, url)
	if err != nil {
		return 0, errs.NewBadRequest("Cannot get video minutes")
	}

	duration := data.Format.Duration().Minutes()

	return duration, nil
}
