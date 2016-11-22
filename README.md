# Gratitude Log

Gratitude Log is a program to keep a diary of things you are grateful for, and to transform these data into insights.

## Goals of the program:

* Cultivate the attention to gratitude
* Keep track of goal-oriented behavior
* Support a data-driven lifestyle
* Reinforce smart behavior

# Writing your Gratitude Log

You can either write it as a structured .txt or .json file.

* The priority for now is to keep the program working with a JSON file.

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

 * Everies are separated by a comma.
 * First word of each entry has the highest weight.
 * "SOCIAL > friends > John Doe time" will count John Doe time also as an entry to friends, and friends as an entry to SOCIAL. This makes it easy to keep track of where your joy comes from.

## Gratitude JSON formatting example

```
{
    "Relations": null,
    "Pages": [
        {
            "Date": "2016-11-07T00:00:00Z",
            "Title": "Title That Summarizes Your Day",
            "Categories": {
                "gratitude": [
                    "SOCIAL > friends > John time",
					"SOCIAL > friends > Nele time",
            		"SOCIAL > friends > Jane time",
					"weather nice",
					"Python time",
                    "lunch tasty"
                 ],
                "geo": [
                    "Amsterdam"
                ]
            }
        },
        {
            "Date": "2016-11-08T00:00:00Z",
            "Title": "Title That Summarizes Your Day 2",
            "Categories": {
                "gratitude": [
                    "sunrise amazing",
					"weather nice",
					"SOCIAL > friends > Nele time",
            		"SOCIAL > friends > Jane time",
			        "lunch tasty",
                    "evening relaxed"
                 ],
                "geo": [
                    "Amsterdam"
                ]
            }
        },
        {
            "Date": "2016-11-09T00:00:00Z",
            "Title": "Title That Summarizes Your Day 3",
            "Categories": {
                "newcategory": [
			        "categories can be added"
                ],
                "gratitude": [
			        "entry 1",
                    "last entry without a comma"
                 ],
                "geo": [
                    "Amsterdam"
                ]
            }
        }
    ]
}
```
