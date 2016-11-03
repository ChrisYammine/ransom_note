package main
import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

// Input format:
// The first line contains two space-separated integers describing the respective values of  (the number of words in the magazine) and  (the number of words in the ransom note).
// The second line contains  space-separated strings denoting the words present in the magazine.
// The third line contains  space-separated strings denoting the words present in the ransom note.

func main() {
  scanner := bufio.NewScanner(os.Stdin);
  scanner.Scan();
  word_counts := strings.Split(scanner.Text(), " "); // first is magazine word count | second is ransom note word count
  m_word_count, _ := strconv.Atoi(word_counts[0]);
  r_word_count, _ := strconv.Atoi(word_counts[1]);

  scanner.Scan();
  magazine_words := strings.Split(scanner.Text(), " ");

  scanner.Scan();
  ransom_words := strings.Split(scanner.Text(), " ");

  if (r_word_count > m_word_count) {
    fmt.Println("No");
    return
  }

  // fmt.Println(magazine_words);
  // fmt.Println(ransom_words);
}
