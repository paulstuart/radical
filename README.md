# radical
Store floating point latitude and longitude as a sortable uint64

This allows indexing geo coordinates on a single column and also trims floating point fuzziness to a respectable 7 significant digits, which is ~4 inches of accuracy.