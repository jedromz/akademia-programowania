# Podstawy Go

W pliku `academy.go` znajdują się cztery funkcje wymagające 
zaimplementowania. Należy implementować je w kolejności, w 
jakiej są zdefiniowane, wykorzystując poprzednie funkcje tam, 
gdzie ma to sens.

By sprawdzić poprawność implementacji należy wykonać testy. 
Służy do tego narzędzie `go test`, które uruchamiamy w katalogu, 
w którym znajduje się plik `academy_test.go`. Pliku tego nie 
należy modifkować, ale ciekawskich zapraszam do przyjrzenia się 
mu. Więcej o testach będzie w ramach późniejszego wykładu.

W trakcie rozwiązywania tego zadania pomocna może okazać się 
flaga `failfast`, która spowoduje zakończenie wykonywania testów 
po pierwszym błędzie. Dzięki temu komunikat błędu będzie związany 
z tym, nad czym w danym momencie pracujecie.

Domyślnie `go test` nie wyświetla testów zakończonych powodzeniem.
By wyświetlić wszystkie testy, niezalenie od wyniku, trzeba dodać
flagę `v` (verbose).

`go test -failfast -v`
