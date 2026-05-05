package domain

type RobotPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type RobotMotionDomain struct {
	Dev       string          `json:"dev"`
	Positions []RobotPosition `json:"positions"`
}
