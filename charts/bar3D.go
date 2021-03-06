package charts

import (
	"github.com/chenjiandongx/go-echarts/common"
)

type Bar3D struct {
	Chart3D
}

func (Bar3D) chartType() string { return common.ChartType.Bar3D }

type Bar3DOpts struct {
	Shading string
}

func (Bar3DOpts) markSeries() {}

func (opt *Bar3DOpts) setChartOpt(s *singleSeries) {
	s.Shading = opt.Shading
}

func NewBar3D(routers ...RouterOpts) *Bar3D {
	chart := new(Bar3D)
	chart.initBaseOpts(false, routers...)
	chart.initChart3D()
	return chart
}

func (c *Bar3D) AddXYAxis(xAxis, yAxis interface{}) *Bar3D {
	c.xData = xAxis
	c.yData = yAxis
	return c
}

func (c *Bar3D) AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Bar3D {
	c.addZAxis(common.ChartType.Bar3D, name, zAxis, options...)
	return c
}
