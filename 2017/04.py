def no_duplicates(words):
  valid = True
  while len(words) >= 1:
    word = words.pop()
    if word in words:
      valid = False
  return valid

def sort_words(phrase):
  sortedWords = []
  words = phrase.split()
  valid = True
  for word in words:
    sortedLetters = []
    for letter in word:
      sortedLetters.append(letter)
    sortedLetters.sort()
    sortedWords.append("".join(sortedLetters))
  return sortedWords


def valid_passphrase(phrase):
  words = sort_words(phrase)
  return no_duplicates(words)


def count_valid_passphrases(filename):
  file = open(filename, "r")
  passphrases = []
  for line in file:
    passphrases.append(line)


  count = 0
  for passphrase in passphrases:
    if valid_passphrase(passphrase):
      count += 1

  print count

count_valid_passphrases("2017-04.txt")
