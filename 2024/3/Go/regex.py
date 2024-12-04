import sys
import re

def main():
  text = sys.argv[1]

  pattern = r"(?s)don't(.*?)do(?!n't)"
  processed_text = re.sub(pattern, "", text)

  print(processed_text)

if __name__ == "__main__":
  main()