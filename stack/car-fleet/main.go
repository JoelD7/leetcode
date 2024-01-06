package main

func carFleet(target int, position []int, speed []int) int {
	fleets := make([]Stack, 0)

	var fleetMotion, carMotion int

	for i := 0; i < len(position); i++ {
		if len(fleets) == 0 {
			stack := NewStack()
			stack.Push(&Motion{position: position[i], speed: speed[i]})
			continue
		}

		for _, fleet := range fleets {
			if fleet.length == 1 {
				fleetMotion = fleet.Peek().position + fleet.Peek().speed
				carMotion = position[i] + speed[i]

				if fleetMotion == carMotion && (fleetMotion+carMotion) <= target {
					fleet.Push(&Motion{position: position[i], speed: speed[i]})
					break
				}

				continue
			}

			carMotion = position[i] + speed[i]

			if fleet.Peek().position == carMotion && (fleet.Peek().position+carMotion) <= target {
				fleet.Push(&Motion{position: position[i], speed: speed[i]})
				break
			}

		}
	}
}

type Motion struct {
	position int
	speed    int
}

type Node struct {
	val         *Motion
	fleetMotion *Motion
	prev        *Node
}

type Stack struct {
	head   *Node
	length int
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Push(m *Motion) {
	s.length++

	if s.head == nil {
		s.head = &Node{
			val:         m,
			fleetMotion: m,
			prev:        nil,
		}
		return
	}

	curHead := &Node{s.head.val, s.head.fleetMotion, s.head.prev}
	s.head.val = m
	s.head.prev = curHead
	fleetSpeed := m.speed
	if s.head.fleetMotion.speed < m.speed {
		fleetSpeed = s.head.fleetMotion.speed
	}

	s.head.fleetMotion = &Motion{
		s.head.fleetMotion.position + m.position,
		fleetSpeed,
	}
}

func (s *Stack) Peek() *Motion {
	return s.head.fleetMotion
}
