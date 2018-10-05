package Prac2

func sorting(list [10]int) [10]int{

    for i := 0 ; i < len(list) ; i++ {
        for j := 0 ; j < len(list)-1 ; j++ {
            if list[j] > list[j+1] {
                list[j], list[j+1] = list[j+1], list[j]
            }
        }
    }
    return list
}
