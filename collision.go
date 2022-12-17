package main

// CheckCollision 检查两个物体间是否碰撞
func CheckCollision(bullet, alien Entity) bool {
	alienTop, alienLeft := alien.Y(), alien.X()
	alienBottom, alienRight := alien.Y()+float64(alien.Height()), alien.X()+float64(alien.Width())

	// 左上角
	x, y := bullet.X(), bullet.Y()
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	// 右上角
	x, y = bullet.X()+float64(bullet.Width()), bullet.Y()
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	// 左下角
	x, y = bullet.X(), bullet.Y()+float64(bullet.Height())
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	// 右下角
	x, y = bullet.X()+float64(bullet.Width()), bullet.Y()+float64(bullet.Height())
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	return false
}
