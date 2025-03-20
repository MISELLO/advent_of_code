document.addEventListener("DOMContentLoaded", function () {
    const years = Object.keys(texts); // We get the years from the object texts on data.js
    const yearsContainer = document.getElementById("years-container");

    years.forEach(year => {
        let yearDiv = document.createElement("div");
        yearDiv.classList.add("year");
        yearDiv.textContent = year;
        yearsContainer.appendChild(yearDiv);

        let daysDiv = document.createElement("div");
        daysDiv.classList.add("days");
        daysDiv.style.display = "none";

        let chooseADay = document.createElement("h2")
        chooseADay.textContent = "Choose a day"
        daysDiv.appendChild(chooseADay)

        Object.keys(texts[year]).forEach(day => {
            let dayDiv = document.createElement("div");
            dayDiv.classList.add("day");
            dayDiv.style.display = "flex";
            dayDiv.style.flexWrap = "wrap";
            dayDiv.style.alignItems = "center";

            // Day title
            let dayTitle = document.createElement("span");
            dayTitle.textContent = `Day ${day}`;
            dayTitle.style.cursor = "pointer";
            dayTitle.style.marginRight = "auto";

            let codeLink = document.createElement("a");
            day < 10 ? n = "0" + day : n = day;
            codeLink.href = `https://github.com/MISELLO/advent_of_code/tree/main/${year}/Day_${n}`;
            codeLink.target = "_blank"; 
            codeLink.textContent = "Code";
            codeLink.style.marginLeft = "10px";

            let detailsLink = document.createElement("a");
            detailsLink.href = `https://adventofcode.com/${year}/day/${day}`;
            detailsLink.target = "_blank"; 
            detailsLink.textContent = "Details";
            detailsLink.style.marginLeft = "10px";

            // dayContent -> codeLink, detailsLink and parts 1 & 2
            let dayContent = document.createElement("div");
            dayContent.style.display = "none";
            dayContent.style.width = "100%";

            // Part 1
            const part1 = document.createElement("div");
            part1.classList.add("part");

            const part1Title = document.createElement("span");
            part1Title.classList.add("part-title");
            part1Title.innerText = "Part 1";

            const part1Content = document.createElement("p");
            part1Content.classList.add("part-content");
            part1Content.innerHTML = texts[year][day].part1 || "Pending ...";

            part1.appendChild(part1Title);
            part1.appendChild(part1Content);

            // Part 2
            const part2 = document.createElement("div");
            part2.classList.add("part");

            const part2Title = document.createElement("span");
            part2Title.classList.add("part-title");
            part2Title.innerText = "Part 2";

            const part2Content = document.createElement("p");
            part2Content.classList.add("part-content");
            part2Content.innerHTML = texts[year][day].part2 || "Pending ...";

            part2.appendChild(part2Title);
            part2.appendChild(part2Content);

            // Add parts to dayContent
            dayContent.appendChild(part1);
            dayContent.appendChild(part2);

            // Event to show/hide the day content
            dayDiv.addEventListener("click", function () {
                dayContent.style.display = dayContent.style.display === "none" ? "block" : "none";
            });

            // Add the content to the day
            dayDiv.appendChild(dayTitle);
            dayDiv.appendChild(codeLink);
            dayDiv.appendChild(detailsLink);
            dayDiv.appendChild(dayContent); // Hidden by default

            daysDiv.appendChild(dayDiv);

        });

        // Event to show/hide the year content
        yearDiv.addEventListener("click", function () {
            daysDiv.style.display = daysDiv.style.display === "none" ? "block" : "none";
        });

        yearsContainer.appendChild(daysDiv);
    });
});



