document.addEventListener("DOMContentLoaded", function () {
    const years = [2023, 2024]; // Se han añadido los años 2023 y 2024
    const daysCount = 25;

    const yearsContainer = document.getElementById("years-container");

    years.forEach(year => {
        let yearDiv = document.createElement("div");
        yearDiv.classList.add("year");
        yearDiv.textContent = year;
        yearsContainer.appendChild(yearDiv);

        let daysDiv = document.createElement("div");
        daysDiv.classList.add("days");
        daysDiv.style.display = "none";

        for (let i = 1; i <= daysCount; i++) {
            let dayDiv = document.createElement("div");
            dayDiv.classList.add("day");
            dayDiv.textContent = `Day ${i}`;

            let codeLink = document.createElement("a");
            i < 10  ? n = "0" + i : n = i
            codeLink.href = `https://github.com/MISELLO/advent_of_code/tree/main/${year}/Day_${n}`;
            codeLink.target = "_blank"; 
            codeLink.textContent = "Code";

            let detailsLink = document.createElement("a");
            detailsLink.href = `https://adventofcode.com/${year}/day/${i}`;
            detailsLink.target = "_blank"; 
            detailsLink.textContent = "Details";

            let part1 = document.createElement("div");
            part1.classList.add("part");
            part1.textContent = "Part 1";

            let part2 = document.createElement("div");
            part2.classList.add("part");
            part2.textContent = "Part 2";

            dayDiv.appendChild(codeLink);
            dayDiv.appendChild(detailsLink);
            dayDiv.appendChild(part1);
            dayDiv.appendChild(part2);

            daysDiv.appendChild(dayDiv);
        }

        yearDiv.addEventListener("click", function () {
            daysDiv.style.display = daysDiv.style.display === "none" ? "block" : "none";
        });

        yearsContainer.appendChild(daysDiv);
    });
});

