package kata

import "strconv"

func HowMuch(m int, n int) [][3]string {
    var results [][3]string

    if (m > n) {
        m, n = n, m
    }

    for i := m; i <= n; i++ {
        if (i - 1) % 9 == 0 && (i - 2) % 7 == 0 {
            results = append(results, MakeString(i, (i-2)/7, (i-1)/9))
        }
    }

    return results
}

func MakeString(int money, int car, int boat) [3]string {
  return [3]string {
        "M: " + strconv.Itoa(money),
        "B: " + strconv.Itoa(car),
        "C: " + strconv.Itoa(boat)
  }
}
