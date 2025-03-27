const texts = {
	"2023": {
		"1" : {
			part1: "For each line we check each character and check if it's a digit.<br>" +
				"We store the first and the last one, join them together and convert it to a number.<br>" +
				"We add this number to the total and we get the resut.",
			part2: "This can be done in many ways.<br>" +
				"Instead of replacing the digits in words with their equivalent number and repeat what we dit on part 1,<br>" +
				"I decided to work with indexes.<br>"+
				"I declared an array/slice of the valid \"strings\". This is:<br>" +
				"'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',<br>" +
				"'zero', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine'<br>" +
				"Then, for each line, I searched for this strings and stored the first and last index and their value" +
				" (the value is the index of the array%10).<br>" +
				"The number we are looking for is the first value * 10 + the last value.<br>" +
				"We add the value of each line and we get the answer."},
		"2" : {
			part1: "The only possible issue here is by commiting an error processing the data.<br>" +
				"When having a game loaded we check if all sets are correct by checking if" +
				" the amount of each cube is higher than what we have been told at the start.<br>" +
				"At the end of each game, if the game was valid, we add the game ID to the answer.",
			part2: "This second part is easy if we process the data correctly, as on the first part.<br>" +
				"I had a record of the maximum of each type of cube for each game, then multiply this" +
				" maximums and add the result to the answer."},
		"3" : {
			part1: "I first created a list of the numbers and thier position.<br>" +
				"Then, with a function, I checked if there is a symbol around it" +
				" (around = Up, Down, Left and Right the number).<br>" +
				"If this function returns true, then we have to add the value of this number to the answer.",
			part2: "Now, the list of numbers not only contain their position, but also if they have" +
				" a gear or not and the position of this gear.<br>" +
				"Then I iterate all numbers with gears and for each one I search for it's counterpart" +
				" on the rest of the list to avoid counting a pair twice" +
				" (It's counterpart will be a number with a gear at the same position).<br>" +
				"Once found it's counterpart I multiply both values and add the result to the answer."},
		"4" : {
			part1: "I isoleted each list of numbers in two slices.<br>" +
				"Then, for each number of the first list I checked if it is present on the second one.<br>" +
				"If the number is present I set the points of this card to 1 the first time and double" +
				" it the other times.<br>" +
				"When we are done with a card we add the points to the answer and take another one.",
			part2: "This time I had to declare a card type that contains the two lists of numbers" +
				" and a multiplyer.<br>" +
				"After loading all the cards into a slice we proceed in a similar way to part 1 but" +
				" instead of counting points we count the number of matches. <br>" +
				"This means the next N cards (being N the number of matches) will have thier multiplyer" +
				" increased by the current card multiplyer.<br>" +
				"To get the answer I had to iterate all cards again and add all multiplyers."},
		"5" : {
			part1: "First I defined several states to load the data. Each state represents what we" +
				" are going to read from the input file.<br>" +
				"The structure that will contain the data is just a struct with 3 integers:<br>" +
				" Destination, source and range (on this order).<br>" +
				"Each conversion \"map\" is a slice of this structure, except the seeds, this is a" +
				" simple slice of integers.<br>" +
				"I load the information on this structures with 8 functions, one for the seeds and" +
				" the rest for the conversion maps.<br>" +
				"I use the states mentioned before to invoque this functions in order (and seems the" +
				" states are not really necessary).<br>" +
				"Now we need to define the convert function, that takes an int and a map and returns" +
				" the proper destination. To find the destination we iterate the map searching if the" +
				" given int falls into the range (between source and source+range) if that's the case" +
				" we return the destination + range, if we don't find any range that contains the given" +
				" int, we return the int directly because: \"Any source numbers that aren't mapped" +
				" correspond to the same destination number\".<br>" +
				"Once all information is loaded I iterated the seeds, did the cascade of conversions" +
				" until location and kept track of the lowest one.<br>" +
				"When the iteration is done the answer is the lowest one we found.",
			part2: "Now, instead of a few seeds we need to plant several ranges of seeds.<br>" +
				"To have it solved in a timely manner I decided to execute each range in a separated" +
				" thread and each thread leaves the lowest location to a predefined position of" +
				" a slice for this purpose.<br>" +
				"At the end, I just have to search in this slice of lowest locations to find the lowest."},
		"6" : {
			part1: "",
			part2: "" },
		"7" : {
			part1: "",
			part2: "" },
		"8" : {
			part1: "",
			part2: "" },
		"9" : {
			part1: "",
			part2: "" },
		"10": {
			part1: "",
			part2: "" },
		"11": {
			part1: "",
			part2: "" },
		"12": {
			part1: "",
			part2: "" },
		"13": {
			part1: "",
			part2: "" },
		"14": {
			part1: "",
			part2: "" },
		"15": {
			part1: "",
			part2: "" },
		"16": {
			part1: "",
			part2: "" },
		"17": {
			part1: "",
			part2: "" },
		"18": {
			part1: "",
			part2: "" },
		"19": {
			part1: "",
			part2: "" },
		"20": {
			part1: "",
			part2: "" },
		"21": {
			part1: "",
			part2: "" },
		"22": {
			part1: "",
			part2: "" },
		"23": {
			part1: "",
			part2: "" },
		"24": {
			part1: "",
			part2: "" },
		"25": {
			part1: "",
			part2: "" }
	},
	"2024": {
		"1" : {
			part1: "",
			part2: "" },
		"2" : {
			part1: "",
			part2: "" },
		"3" : {
			part1: "",
			part2: "" },
		"4" : {
			part1: "",
			part2: "" },
		"5" : {
			part1: "",
			part2: "" },
		"6" : {
			part1: "",
			part2: "" },
		"7" : {
			part1: "",
			part2: "" },
		"8" : {
			part1: "",
			part2: "" },
		"9" : {
			part1: "",
			part2: "" },
		"10": {
			part1: "",
			part2: "" },
		"11": {
			part1: "",
			part2: "" },
		"12": {
			part1: "",
			part2: "" },
		"13": {
			part1: "",
			part2: "" },
		"14": {
			part1: "",
			part2: "" },
		"15": {
			part1: "",
			part2: "" },
		"16": {
			part1: "",
			part2: "" },
		"17": {
			part1: "",
			part2: "" },
		"18": {
			part1: "",
			part2: "" },
		"19": {
			part1: "",
			part2: "" },
		"20": {
			part1: "",
			part2: "" },
		"21": {
			part1: "",
			part2: "" },
		"22": {
			part1: "",
			part2: "" },
		"23": {
			part1: "",
			part2: "" },
		"24": {
			part1: "",
			part2: "" },
		"25": {
			part1: "",
			part2: "" }
	}
};

