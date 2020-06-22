package carbon

// type Hoverable interface {
//   Contains(x, y float64) bool
//   Land()
// }
//
// var hovered []Hoverable
//
// // OnHover adds the object to wide handlers and returns its index > 0
// func OnHover(object Hoverable) {
//   for i, handler := range hovered {
//     if handler == nil || handler == object {
//       hovered[i] = object
//     }
//   }
//   hovered = append(hovered, object)
// }
//
// // HandleHovered transmits the event to all wide handlers
// func HandleHovered(event Event, x, y float64) {
//   if event.Action != Move && event.Action != OnRelease {
//     return
//   }
//   for _, object := range hovered {
//     if object != nil && !object.Contains(x, y) {
//       object.Land()
//     }
//   }
// }

type Focusable interface {
	Defocus(active bool)
}

var active Focusable

// UpdateActive sets a new activated object and deactivates previous
func UpdateActive(object Focusable) {
	if active != object && active != nil {
		active.Defocus(false)
	}
	active = object
}
