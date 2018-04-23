// Package space has methods to give space age
package space

func Age(seconds float64, planet string) float64 {

	check := make(map[string]float64)
	check["Mercury"] = 0.2408467
	check["Venus"] = 0.61519726
	check["Earth"] = 1.00
	check["Mars"] = 1.8808158
	check["Jupiter"] = 11.862615
	check["Saturn"] = 29.447498
	check["Uranus"] = 84.016846
	check["Neptune"] = 164.79132

	earthYearSeconds := 31557600.0

	planetYearSeconds := earthYearSeconds * check[planet]

	return round(seconds / planetYearSeconds)
}

func round(n float64) float64 {
	return float64(int64(n/0.01+0.5)) * 0.01
}
