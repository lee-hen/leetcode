package moving_average_from_data_stream

import "container/list"

type MovingAverage struct {
	queue *list.List
	size int

	lastSum int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		queue: list.New(),
	}
}

func (m *MovingAverage) Next(val int) float64 {
	if m.queue.Len() < m.size {
		m.queue.PushBack(val)

		m.lastSum += val
	} else {
		front := m.queue.Front()
		m.queue.Remove(front)
		m.queue.PushBack(val)

		first := front.Value.(int)
		m.lastSum -= first
		m.lastSum += val
	}

	return float64(m.lastSum) / float64(m.queue.Len())
}
