package carbon

// Element is an interface implemented in every ui object. It allows objects to FitInto their states in tree-like
// order. Commonly every ui element needs to locate, transmit events, FitInto sprites and draw.
type Element interface {

	// Contains must tell if given point is inside the object's Absolute bounds
	Contains(x, y float64) bool

	// FitInto must Update object's Absolute bounds. It receives a container's location
	FitInto(x1, y1, x2, y2 float64)

	// Rasterize is needed to FitInto graphic sprites contained in frames. It's usually being called after resizing of
	// any object, group or the whole window.
	Rasterize()

	// Handle is needed to react on a mouse interaction with the object.
	Handle(event Event, x, y float64)

	// Draw is needed to simply show rendered sprites in the window. Calling each frame.
	Draw(window *Window)
}
