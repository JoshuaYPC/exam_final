package eval

func Ft_missing(arr []int, n int) int {

	somme := 0
	somme_n := 0

	for i := 0; i < n; i++ {
		somme += arr[i]
	}

	somme_n = ((n + 1) * (n + 2)) / 2
	return (somme_n - somme)
}
