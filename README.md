

# Gratitude Log

Gratitude Log is a program to keep a diary of things you are grateful for, and to transform these data into insights.

## Goals of the program:

* Cultivate the attention to gratitude
* Keep track of goal-oriented behavior
* Support a data-driven lifestyle
* Reinforce smart behavior

## Log file formatting example

There is a gratitude log to JSON format converter code available. It only works when you stick to the following formatting:

```
M21.11.2016 - Title for the day

category1Name: gratitude keywords, gratitude keywords, ..., gratitude keywords
category2Name: weather lovely, SOCIAL > friends > John Doe time, Amsterdam,
programming Go, Gratitude Log,
category...Name: keywords
categoryNName: keywords
```

* First letter of the day of week (Estonian and English are supported)
* Date in the following format day.month.year
* Title
* Category name (e.g. joy, develop, tip, location) followed by a colon. You can have as many categories as you wish.
* Keywords associated with that category. 
..* Everies are separated by a comma.
..* First word of each entry has the highest weight.
..* "SOCIAL > friends > John Doe time" will count John Doe time also as an entry to friends, and friends as an entry to SOCIAL. This makes it easy to keep track of where your joy comes from.
