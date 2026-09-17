# Problem
https://leetcode.com/problems/design-add-and-search-words-data-structure/description/

Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:

- `WordDictionary()` Initializes the object. 
- `void addWord(word)` Adds word to the data structure, it can be matched later. 
- `bool search(word)` Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.



### Example:

    Input
    ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
    [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
    Output
    [null,null,null,null,false,true,true,true]
    
    Explanation
    WordDictionary wordDictionary = new WordDictionary();
    wordDictionary.addWord("bad");
    wordDictionary.addWord("dad");
    wordDictionary.addWord("mad");
    wordDictionary.search("pad"); // return False
    wordDictionary.search("bad"); // return True
    wordDictionary.search(".ad"); // return True
    wordDictionary.search("b.."); // return True



### Constraints:

* 1 <= word.length <= 25
* word in addWord consists of lowercase English letters.
* word in search consist of '.' or lowercase English letters.
* There will be at most 2 dots in word for search queries.
* At most 104 calls will be made to addWord and search.

# Solution
The most efficient data structure for word search and prefix matching is the **trie**, so most of the problem is solved by implementing one.

What makes this problem interesting is what you do with wildcards. Note that wildcards aren’t prefixes all the time. If they were, the problem would practically be solved by just implementing the trie alone, as this is a data structure built for finding words with prefixes. However, this is not the case. Wildcards may appear at any position in `word`. So the approach is this:

1. Do a regular trie search for fully formed words
2. For words with wildcards, when the current character being checked(`char`) is wildcard, instead of picking a single path from the current node that matches `char`, you can pick any path because this is what the wildcard means: a “.” stands for any character. Then, recursively do the same for the remaining characters starting at the current node, until reaching the end of the word.
    1. By searching only with the *remaining* characters, we guarantee that wildcard evaluation doesn’t exceed the character count of `word`. If `word = "b.d"` and `char = .`, the search will continue from the current node(the one that “.” represents) with `word = d`, so now we’ll search for a **single**(`word` is now of length 1) node that is a `d` and is also a word.