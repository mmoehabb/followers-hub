package ancillaries

import (
  "fmt"
)

// ancillary function used by db subpackages in order to generate update sql queries.
// This makes update functions more readable and generic.
func GenUpdateQuery(table string, modelmap map[string]string, pk string) (string, []any) {
  var i int = 1
  var values []any

  q := "UPDATE streamers SET "
  for k, v := range modelmap {
    if i > 1 {
      q = fmt.Sprintf("%s, ", q)
    }
    q = fmt.Sprintf("%s%s=$%d", q, k, i)
    values = append(values, v)
    i = i + 1
  }
  q = fmt.Sprintf("%s WHERE %s='%s'", q, pk, modelmap[pk])

  return q, values
}

