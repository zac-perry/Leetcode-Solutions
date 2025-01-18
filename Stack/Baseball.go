// 682. Baseball Game
// Pretty easy, just used a FILO stack (first in, last out)
// Define my own push and pop functions
func push(stack []int, val int) []int {
    return append(stack, val)
}

func pop(stack []int) ([]int, int) {
    length := len(stack)
    if length == 0 {
        return []int{}, 0
    }

    return stack[:length-1], stack[length-1]
}

func calPoints(operations []string) int {
    // FILO? First in last out
    var stack []int
    var total int

    for i := 0; i < len(operations); i++ {
        if operations[i] == "+" {
            newVal := stack[(len(stack)-1)] + stack[(len(stack)-2)]
            stack = push(stack, newVal)
            total += newVal

        } else if operations[i] == "D" {
            newVal := stack[len(stack)-1] * 2
            stack = push(stack, newVal)
            total += newVal

        } else if operations[i] == "C" {
            val := 0
            stack, val = pop(stack)
            total -= val
        } else {
            num, _ := strconv.Atoi(operations[i])
            stack = push(stack, num)
            total += num
        }
    }

    return total
}
