1. 滑动窗口 or 动态规划：
滑动窗口和动态规划乍眼一看都能解决，子数组或者子字符串的最值问题，那两种方法解决的问题有什么区别对于我们何时采用何种方法起到了比较重要的作用  
1). 滑动窗口，用来解决满足指定条件的最值问题，并且当某个子数组或者子字符串不满足这个条件时，并且就算继续延长该子数组或者子字符串的长度，它仍然不会满足要求，这个时候就需要调整窗口左边界，使得子数组或者子字符串可以重新满足条件,eg. leetcode3 longest-substring-without-repeating-characters   
2). 动态规划，某个子数组或子字符串不满足条件时，延长长度是有可能能满足条件的, 并且子数组或者子字符串之间有某种变化关系，eg. leetcode5 longest-palindromic-substring