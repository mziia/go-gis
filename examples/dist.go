package examples

import (
	"fmt"

	"github.com/mziia/gogis"
)

// FindDistance creating two Points and finds distance between them
func FindDistance() {

	point1 := gogis.NewPoint([]float64{37.6175, 55.752})
	point2 := gogis.NewPoint([]float64{37.6048, 55.7649})

	feature1 := gogis.NewFeature()
	feature1.SetProperty("локация", "Кремль")
	feature1.SetID("0001")
	feature1.SetGeometry(point1)

	feature2 := gogis.NewFeature()
	feature2.SetProperty("локация", "Тверская")
	feature2.SetID("0002")
	feature2.SetGeometry(point2)

	fmt.Println("Distance Point - Point: ", gogis.Distance(feature1, feature2))

}
