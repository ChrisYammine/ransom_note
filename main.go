package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

// Input format:
// The first line contains two space-separated integers describing the respective values of  (the number of words in the magazine) and  (the number of words in the ransom note).
// The second line contains  space-separated strings denoting the words present in the magazine.
// The third line contains  space-separated strings denoting the words present in the ransom note.

func main() {
  scanner := bufio.NewScanner(os.Stdin);
  scanner.Scan();

  scanner.Scan();
  magazine_words := strings.Split(scanner.Text(), " ");
  scanner.Scan();
  ransom_words := strings.Split(scanner.Text(), " ");

  var magazine_map = make(map[string]int);

  for _, value := range(magazine_words) {
    magazine_map[value]++;
  }

  for _, value := range(ransom_words) {
    if magazine_map[value]--; magazine_map[value] < 0 {
      fmt.Println("No");
      return;
    }
  }
  fmt.Println("Yes");
}
