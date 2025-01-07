package rgb

import "time"

// Colorpulse will run RGB function
func (r *ActiveRGB) Colorpulse(startTime *time.Time) {
	buf := map[int][]byte{}
	elapsed := time.Since(*startTime).Milliseconds()

	// Calculate progress and reset when it exceeds 1.0
	progress := float64(elapsed) / (r.RgbModeSpeed * 1000)
	if progress >= 1.0 {
		*startTime = time.Now() // Reset startTime to the current time
		elapsed = 0             // Reset elapsed time
		progress = 0            // Reset progress
	}
	color := interpolateColor(r.RGBStartColor, r.RGBEndColor, progress, r.RGBBrightness)

	// Update LED channels
	for j := 0; j < r.LightChannels; j++ {
		buf[j] = []byte{
			byte(color.Red),
			byte(color.Green),
			byte(color.Blue),
		}
		if r.IsAIO && r.HasLCD {
			if j > 15 && j < 20 {
				buf[j] = []byte{0, 0, 0}
			}
		}
	}
	if r.Inverted {
		r.Output = SetColorInverted(buf)
	} else {
		r.Output = SetColor(buf)
	}
}
