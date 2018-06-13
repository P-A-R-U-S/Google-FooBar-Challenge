# The Google Foo Bar Challenge in Golang 

[![Go Report Card](https://goreportcard.com/badge/github.com/P-A-R-U-S/Golang-CQRS)](https://goreportcard.com/report/github.com/P-A-R-U-S/Golang-CQRS) [![Travis-CI](https://travis-ci.org/P-A-R-U-S/Golang-CQRS.svg?branch=master)](https://travis-ci.org/P-A-R-U-S/Golang-CQRS)

## Description

The Google Foo bar challenge has been known for the last 5 years or more as a secret process of hiring developers and programmers all over the world.
It is a secret process and the challenge consists of coding challenges of increasing difficulty as you go along. All challenges has implemented in Go with conversions to Go language 

![](screenshot.png)

### Level 1

#### Minion Labor Shifts

```
Minion Labor Shifts
===================
Commander Lambda’s minions are upset! They’re given the worst jobs on the whole space station, and some of them are starting to complain that even those worst jobs are being allocated unfairly. If you can fix this problem, it’ll prove your chops to Commander Lambda so you can get promoted!
Minions’ tasks are assigned by putting their ID numbers into a list, one time for each day they’ll work that task. As shifts are planned well in advance, the lists for each task will contain up to 99 integers. When a minion is scheduled for the same task too many times, they’ll complain about it until they’re taken off the task completely. Some tasks are worse than others, so the number of scheduled assignments before a minion will refuse to do a task varies depending on the task. You figure you can speed things up by automating the removal of the minions who have been assigned a task too many times before they even get a chance to start complaining.
Write a function called answer(data, n) that takes in a list of less than 100 integers and a number n, and returns that same list but with all of the numbers that occur more than n times removed entirely. The returned list should retain the same ordering as the original list — you don’t want to mix up those carefully-planned shift rotations! For instance, if data was [5, 10, 15, 10, 7] and n was 1, answer(data, n) would return the list [5, 15, 7] because 10 occurs twice, and thus was removed from the list entirely.

```

#### Prison Labor Dodgers
```
# Quiz: Prison Labor Dodgers
# ====================
Commander Lambda is all about efficiency, including using her bunny prisoners for manual labor. But no one's been properly monitoring the labor shifts for a while, and they've gotten quite mixed up. You've been given the task of fixing them, but after you wrote up new shifts, you realized that some prisoners had been transferred to a different block and aren't available for their assigned shifts. And manually sorting through each shift list to compare against prisoner block lists will take forever - remember, Commander Lambda loves efficiency!
Given two almost identical lists of prisoner IDs x and y where one of the lists contains an additional ID, write a function answer(x, y) that compares the lists and returns the additional ID.
For example, given the lists x = [13, 5, 6, 2, 5] and y = [5, 2, 5, 13], the function answer(x, y) would return 6 because the list x contains the integer 6 and the list y doesn't. Given the lists x = [14, 27, 1, 4, 2, 50, 3, 1] and y = [2, 4, -4, 3, 1, 1, 14, 27, 50], the function answer(x, y) would return -4 because the list y contains the integer -4 and the list x doesn't.
In each test case, the lists x and y will always contain n non-unique integers where n is at least 1 but never more than 99, and one of the lists will contain an additional unique integer which should be returned by the function.  The same n non-unique integers will be present on both lists, but they might appear in a different order, like in the examples above. Commander Lambda likes to keep her numbers short, so every prisoner ID will be between -1000 and 1000.
```