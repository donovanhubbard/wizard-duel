package app

type GridUpdateMsg struct {
	Grid Grid
}

type PlanMoveMsg struct {
	ID        string
	Direction string
}

type CastSpellMsg struct {
	ID        string
	Type      string
	Direction string
}
