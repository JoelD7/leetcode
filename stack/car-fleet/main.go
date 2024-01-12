package main

func carFleet(target int, position []int, speed []int) int {
	fleets := make([]*Stack, 0)

	var fleetMotion, carMotion int

	for i := 0; i < len(position); i++ {
		if len(fleets) == 0 {
			stack := NewStack()
			stack.Push(&Motion{position: position[i], speed: speed[i]}, 0, position[i])
			fleets = append(fleets, stack)
			continue
		}

		for _, fleet := range fleets {
			if fleet.length == 1 {
				fleetMotion = fleet.Peek().position + fleet.Peek().speed
				carMotion = position[i] + speed[i]

				if fleetMotion == carMotion && (fleetMotion+carMotion) <= target {
					fleet.Push(&Motion{position: position[i], speed: speed[i]}, 1, fleetMotion)
					break
				}

				continue
			}

			fleetMotion = fleet.Peek().position + fleet.Peek().speed
			carMotion = position[i] + speed[i]

			//For a case like: [10,8,6] [2,4,6]
			if fleet.Peek().position == carMotion && carMotion <= target {
				fleet.Push(&Motion{position: position[i], speed: speed[i]}, 0, fleet.Peek().position)
				continue
			}

			steps, posAfterSteps, ok := willMeet(fleet, position[i], speed[i])
			if ok && posAfterSteps < target {
				fleet.Push(&Motion{position: position[i], speed: speed[i]}, steps, posAfterSteps)
				continue
			}

			stack := NewStack()
			stack.Push(&Motion{position: position[i], speed: speed[i]}, 0, position[i])
			fleets = append(fleets, stack)
		}
	}

	return len(fleets)
}

func willMeet(fleet *Stack, carPos, carSpeed int) (int, int, bool) {
	if !isMoreAdvancedSlower(fleet, carPos, carSpeed) {
		return 0, 0, false
	}

	fleetSum := fleet.Peek().position + fleet.Peek().speed
	carSum := carPos + carSpeed
	steps := fleetSum + carSum + 1

	fleetPosAfterSteps := fleet.Peek().position + (steps * fleet.Peek().speed)
	carPosAfterSteps := carPos + (steps * carSpeed)

	if fleetPosAfterSteps == carPosAfterSteps {
		return steps, carPosAfterSteps, true
	}

	return 0, 0, false
}

// isMoreAdvancedSlower indicates if the object that is farther ahead is slower than the one that's behind.
func isMoreAdvancedSlower(fleet *Stack, carPos, carSpeed int) bool {
	f := fleet.Peek()

	return (f.position > carPos && f.speed < carSpeed) ||
		(f.position < carPos && f.speed > carSpeed)
}

type Motion struct {
	position int
	speed    int
	//The amount of times the fleet has moved
	steps int
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

func (s *Stack) Push(m *Motion, steps, posAfterSteps int) {
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
		posAfterSteps,
		fleetSpeed,
		s.head.fleetMotion.steps + steps,
	}
}

func (s *Stack) Peek() *Motion {
	return s.head.fleetMotion
}
