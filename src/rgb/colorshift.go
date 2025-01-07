package rgb

import "time"

// Colorshift will run RGB function
func (r *ActiveRGB) Colorshift(startTime *time.Time, activeRgb *ActiveRGB) {
	buf := map[int][]byte{}
	elapsed := time.Since(*startTime).Milliseconds()
	color := &Color{}

	// Calculate progress and reset when it exceeds 1.0
	progress := float64(elapsed) / (r.RgbModeSpeed * 1000)
	if progress >= 1.0 {
		*startTime = time.Now() // Reset startTime to the current time
		elapsed = 0             // Reset elapsed time
		progress = 0            // Reset progress
		if activeRgb.Phase == 1 {
			activeRgb.Phase = 0
		} else {
			activeRgb.Phase = 1
		}
	}

	if activeRgb.Phase == 0 {
		color = interpolateColor(r.RGBStartColor, r.RGBEndColor, progress, r.RGBBrightness)
	} else {
		color = interpolateColor(r.RGBEndColor, r.RGBStartColor, progress, r.RGBBrightness)
	}

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
